package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nathan.simpledb/src/internal/field"
	"nathan.simpledb/src/internal/tuple"
	"nathan.simpledb/src/test/util"
)

func Test_Tuple_ModifyFields(t *testing.T) {
	td := util.GetTupleDesc(2, "")
	tup := tuple.NewTuple(td)

	tup.SetField(0, field.NewIntField(-1))
	tup.SetField(1, field.NewIntField(0))

	assert.Equal(t, field.NewIntField(-1), tup.GetField(0))
	assert.Equal(t, field.NewIntField(0), tup.GetField(1))

	tup.SetField(0, field.NewIntField(1))
	tup.SetField(1, field.NewIntField(37))

	assert.Equal(t, field.NewIntField(1), tup.GetField(0))
	assert.Equal(t, field.NewIntField(37), tup.GetField(1))
}

func Test_Tuple_GetTupleDesc(t *testing.T) {
	td := util.GetTupleDesc(5, "")
	tup := tuple.NewTuple(td)
	assert.Equal(t, td, tup.GetTupleDesc())
}

func Test_Tuple_ModifyRecordId(t *testing.T) {
	//tup := tuple.NewTuple(util.GetTupleDesc(1, ""))
	//pid1 := page.NewHeapPageId(0, 0)
	//id := page.NewRecordId(pid1, 0)
	//tup.SetRecordId(rid1)
	// cant impl yet

}