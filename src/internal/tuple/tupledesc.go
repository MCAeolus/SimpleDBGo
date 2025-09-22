package tuple

import (
	"fmt"
	"slices"
	"strings"

	"nathan.simpledb/src/internal/errors"
	"nathan.simpledb/src/internal/interfaces"
)

// types are the value types, fields represents the names of the types
// name could be null
// tuple descs are _ordered_
func NewTupleDesc(types []interfaces.Type, fields []string) TupleDesc {
	
	byteSize := 0
	for _, t := range types {
		byteSize += t.GetLen()
	} 
	
	return TupleDesc{
		types: types,
		fields: fields,
		byteSize: byteSize,
		numFields: len(types),
	}
}

func CombineTupleDesc(td1, td2 TupleDesc) TupleDesc {
	// these list combines could be improved
	combinedTypes := slices.Clone(td1.types)
	for _, t2 := range td2.types {
		combinedTypes = append(combinedTypes, t2)
	}
	combinedFields := slices.Clone(td1.fields)
	for _, f2 := range td2.fields {
		combinedFields = append(combinedFields, f2)
	}
	return TupleDesc{
		types: combinedTypes,
		fields: combinedFields,
		byteSize: td1.byteSize + td2.byteSize,
		numFields: td1.numFields + td2.numFields,
	}
}

type TupleDesc struct {
	types []interfaces.Type
	fields []string
	byteSize int
	numFields int
}

// # of fields
func (t *TupleDesc) GetLength() int {
	return len(t.types)
}

func (t *TupleDesc) GetFieldName(index int) (string, error) {
	if len(t.fields) <= index{
		return "", errors.ErrDoesNotExist
	}
	return t.fields[index], nil
}

func (t *TupleDesc) NameToIndex(name string) (int, error) {
	// we could store a hash of our fields, for now just use an inefficient lookup
	loc := slices.Index(t.fields, name)
	if loc == -1 {
		return loc, errors.ErrDoesNotExist
	}
	return loc, nil
}

func (t *TupleDesc) GetType(index int) (interfaces.Type, error) {
	if len(t.types) <= index {
		return nil, errors.ErrDoesNotExist
	}
	return t.types[index], nil
}

// size in bytes of the tuples contained in TupleDesc
func (t *TupleDesc) GetSize() int {
	return t.byteSize
}

// ignores field name
// also todo: verify uniqueness of types
func (t *TupleDesc) Equals(o *TupleDesc) bool {
	if o == nil {
		return false
	}
	if len(t.types) != len(o.types) {
		return false
	}
	for ix, typ := range t.types {
		if typ != o.types[ix] {
			return false
		}
	}
	return true
}

func (t *TupleDesc) String() string {
	var sb strings.Builder
	for ix, typ := range t.types {
		delim := ""
		if ix > 0 {
			delim = ", "
		}
		fname := t.fields[ix]
		if fname == "" {
			fname = "`?`"
		}
		sb.WriteString(fmt.Sprintf("%s%s(%s)", delim, typ, fname))

	}
	return sb.String()
}