package util

import (
	"fmt"

	"nathan.simpledb/src/internal/field"
	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/tuple"
)

// n fields, type INT_TYPE name name+n
func GetTupleDesc(n int, name string) tuple.TupleDesc {
	
	types := []interfaces.Type{}
	fields := []string{}
	for ix := range n {
		types = append(types, field.INT_TYPE)
		fields = append(fields, fmt.Sprintf("%s%d", name, ix))
	}
	return tuple.NewTupleDesc(
		types,
		fields,
	)
}

func CombineTupleStringArrays(td1, td2, combined tuple.TupleDesc) bool {
	ix := 0
	for i := range td1.GetLength() {
		f1, err := td1.GetFieldName(i)
		if err != nil {
			return false
		}
		fc, err := combined.GetFieldName(ix)
		if err != nil {
			return false
		}
		if f1 != fc {
			return false
		}
		ix++
	}	
	for i := range td2.GetLength() {
		f1, err := td2.GetFieldName(i)
		if err != nil {
			return false
		}
		fc, err := combined.GetFieldName(ix)
		if err != nil {
			return false
		}
		if f1 != fc {
			return false
		}
		ix++
	}	
	return true
}