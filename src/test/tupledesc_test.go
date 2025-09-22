package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"nathan.simpledb/src/internal/errors"
	"nathan.simpledb/src/internal/field"
	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/tuple"
	"nathan.simpledb/src/test/util"
)

func Test_TupleDesc_Combine(t *testing.T) {
	var td1, td2, td3 tuple.TupleDesc
	td1 = util.GetTupleDesc(1, "td1")
	td2 = util.GetTupleDesc(2, "td2")

	// td.combine td1+td2
	td3 = tuple.CombineTupleDesc(td1, td2)
	assert.Equal(t, 3, td3.GetLength())
	assert.Equal(t, 3 * field.INT_TYPE.GetLen(), td3.GetSize())
	for i := range 3 {
		ft, err := td3.GetType(i)
		assert.Nil(t, err)
		assert.Equal(t, field.INT_TYPE, ft)
	}
	assert.True(t, util.CombineTupleStringArrays(td1, td2, td3))

	// combine td2+td1
	td3 = tuple.CombineTupleDesc(td2, td1)
	assert.Equal(t, 3, td3.GetLength())
	assert.Equal(t, 3 * field.INT_TYPE.GetLen(), td3.GetSize())
	for i := range 3 {
		ft, err := td3.GetType(i)
		assert.Nil(t, err)
		assert.Equal(t, field.INT_TYPE, ft)
	}
	assert.True(t, util.CombineTupleStringArrays(td2, td1, td3))

	// combine td2+td2
	td3 = tuple.CombineTupleDesc(td2, td2)
	assert.Equal(t, 4, td3.GetLength())
	assert.Equal(t, 4 * field.INT_TYPE.GetLen(), td3.GetSize())
	for i := range 4 {
		ft, err := td3.GetType(i)
		assert.Nil(t, err)
		assert.Equal(t, field.INT_TYPE, ft)
	}
	assert.True(t, util.CombineTupleStringArrays(td2, td2, td3))
}

func Test_TupleDesc_GetType(t *testing.T) {
	lengths := []int{1, 2, 1000}
	for _, len := range lengths {
		td := util.GetTupleDesc(len, "")
		for i := range len {
			f1, err := td.GetType(i)
			assert.Nil(t, err)
			assert.Equal(t, field.INT_TYPE, f1)
		}
	}
}

func Test_TupleDesc_NameToId(t *testing.T) {
	lengths := []int{1, 2, 1000}
	prefix := "test"

	for _, len := range lengths {
		td := util.GetTupleDesc(len, prefix)
		for i := range len {
			f1, err := td.NameToIndex(fmt.Sprintf("%s%d", prefix, i))
			assert.Nil(t, err)
			assert.Equal(t, i, f1)
		}

		f1, err := td.NameToIndex("foo")
		assert.ErrorIs(t, errors.ErrDoesNotExist, err)
		assert.Equal(t, -1, f1)

		td = util.GetTupleDesc(len, "")
		f1, err = td.NameToIndex(prefix)
		assert.ErrorIs(t, errors.ErrDoesNotExist, err)
		assert.Equal(t, -1, f1)
	}
}

func Test_TupleDesc_GetSize(t *testing.T) {
	lengths := []int{1, 2, 1000}
	for _, len := range lengths {
		td := util.GetTupleDesc(len, "")
		assert.Equal(t, len*field.INT_TYPE.GetLen(), td.GetSize())
	}
}

func Test_TupleDesc_NumFields(t *testing.T) {
	lengths := []int{1, 2, 1000}
	for _, len := range lengths {
		td := util.GetTupleDesc(len, "")
		assert.Equal(t, len, td.GetLength())
	}
}

func Test_TupleDesc_Equals(t *testing.T) {
	singleInt := tuple.NewTupleDesc([]interfaces.Type{field.INT_TYPE}, nil)
	singleInt2 := tuple.NewTupleDesc([]interfaces.Type{field.INT_TYPE}, nil)
	intString := tuple.NewTupleDesc([]interfaces.Type{field.INT_TYPE, field.STRING_TYPE}, nil)

	assert.False(t, singleInt.Equals(nil))
	
	assert.True(t, singleInt.Equals(&singleInt2))
	assert.True(t, singleInt.Equals(&singleInt))
	assert.True(t, singleInt2.Equals(&singleInt))
	assert.True(t, intString.Equals(&intString))

	assert.False(t, singleInt.Equals(&intString))
	assert.False(t, singleInt2.Equals(&intString))
	assert.False(t, intString.Equals(&singleInt))
	assert.False(t, intString.Equals(&singleInt2))
}
