package catalog

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strings"

	"nathan.simpledb/src/internal/errors"
	"nathan.simpledb/src/internal/field"
	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/tuple"
)

// singleton
var inited bool = false
var catalog Catalog

func GetCatalog() *Catalog {
	if !inited {
		return nil
	}
	return &catalog
}

func InitCatalog() {
	if inited {
		return
	}
	catalog = Catalog{
		tableMapping: map[string]int{},
		tables: []tableInfo{},
	}
	inited = true
}

type tableInfo struct {
	name string
	pkey string
	tableId int // comes from file.GetId()
	description tuple.TupleDesc
	file interfaces.DbFile
}

type Catalog struct {
	tableMapping map[string]int
	tables map[int]tableInfo
}

// stored in dbfile
//pkey can be empty ""
func (c *Catalog) AddTable(file interfaces.DbFile, name string, pkeyField string) {
	// we override existing 'name' (so not checking for existence)
	ti := tableInfo{
		name: name,
		pkey: pkeyField,
		file: file,
		tableId: file.GetId(),
		description: file.GetTupleDesc(),
	}
	c.tableMapping[name] = ti.tableId
	c.tables[ti.tableId] = ti // todo: this might be problematic
}

func (c *Catalog) GetTableId(name string) (int, error) {
	if ti, ok := c.tableMapping[name]; ok {
		return ti, nil
	}
	return -1, errors.ErrDoesNotExist
}

func (c *Catalog) GetTupleDesc(tableId int) (tuple.TupleDesc, error) {
	if len(c.tables) < tableId {
		return tuple.TupleDesc{}, errors.ErrDoesNotExist
	}
	return c.tables[tableId].description, nil
}

func (c *Catalog) GetDbFile(tableId int) (interfaces.DbFile, error) {
	if len(c.tables) < tableId {
		return nil, errors.ErrDoesNotExist
	}
	return c.tables[tableId].file, nil
}

func (c *Catalog) GetPrimaryKey(tableId int) (string, error) {
	if len(c.tables) < tableId {
		return "", errors.ErrDoesNotExist
	}
	return c.tables[tableId].pkey, nil
}

func (c *Catalog) GetTableName(tableId int) (string, error) {
	if len(c.tables) < tableId {
		return "", errors.ErrDoesNotExist
	}
	return c.tables[tableId].name, nil
}

func (c *Catalog) Clear() {
	c.tableMapping = map[string]int{}
	c.tables = map[int]tableInfo{}
}

func (c *Catalog) TableIdIterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		for k := range c.tables {
			if !yield(k) {
				return
			}
		}
	}	
}

func (c *Catalog) LoadSchema(catalogFile string) {
	file, err := os.Open(catalogFile)
	if err != nil {
		panic(fmt.Sprintf("file %s does not exist", file))
	}
	defer file.Close()
	br := bufio.NewScanner(file)
	lineNo := 0
	for br.Scan() {
		line := br.Text()
		// format:name (fieldname type, fieldname type, ...)
		name := strings.SplitN(line, " ", 1)
		if len(name) < 1 {
			panic(fmt.Sprintf("line %d: could not parse `name`", lineNo))
		}
		fieldStart := strings.Index(line, "(")
		fieldEnd := strings.Index(line, ")")
		if fieldStart < 0 || fieldEnd < 0 {
			panic(fmt.Sprintf("line %d: could not parse `fields` (start=%d, end=%d)", lineNo, fieldStart, fieldEnd))
		}
		fields := line[fieldStart:fieldEnd]
		fieldTypes := strings.Split(fields, ",")
		names := []string{}
		types := []interfaces.Type{}
		for fieldNo, typ := range fieldTypes {
			els := strings.SplitN(typ, " ", 3)
			if len(els) < 2 {
				panic(fmt.Sprintf("line %d: could not parse `fields`[%d]", lineNo, fieldNo))
			}
			colName, colType := els[0], els[1]
			names = append(names, strings.Trim(colName, " "))
			switch strings.Trim(colType, " ") {
			case "int":
				types = append(types, field.INT_TYPE)
			case "string":
				types = append(types, field.STRING_TYPE)
			default:
				panic(fmt.Sprintf("line %d: could not parse `fields`[%d]: unknown type %s", lineNo, fieldNo, colType))
			}
			if len(els) < 3 {
				continue
			}
			annotation := els[2]
			switch strings.Trim(annotation, " ") {
			case "pk":
			default:
				panic(fmt.Sprintf("line %d: could not parse `fields`[%d]: unknown annotation %s", lineNo, fieldNo, annotation))
			}

		}



		lineNo++
	}

}

