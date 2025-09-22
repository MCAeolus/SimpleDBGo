package field

import (
	"fmt"
	"io"

	"nathan.simpledb/src/internal/enums"
	"nathan.simpledb/src/internal/errors"
	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/util"
)

type IntField struct {
	value int32
}

func NewIntField(value int32) IntField {
	return IntField{value: value}
}

func (i IntField) GetValue() int32 {
	return i.value
}

// TODO: serializing- we need to confirm we de/serialize back to signed
func (i IntField) Serialize(writer io.Writer) error {
	_, err := writer.Write(util.IntToBytes(i.value))
	return err
}

func (i IntField) String() string {
	return fmt.Sprintf("%d", i.value)
}

func (i IntField) GetType() interfaces.Type {
	return INT_TYPE
}

func (i IntField) Compare(op enums.PredicateOp, value interfaces.Field) (bool, error) {
	// value MUST be IntField or will error
	t, ok := value.(IntField) 
	if !ok {
		return false, errors.ErrIllegalCast
	}
	switch (op) {
	case enums.EQUALS:
		return i.value == t.value, nil
	case enums.NOT_EQUALS:
		return i.value != t.value, nil
	case enums.GREATER_THAN:
		return i.value > t.value, nil
	case enums.GREATER_THAN_OR_EQ:
		return i.value >= t.value, nil
	case enums.LESS_THAN:
		return i.value < t.value, nil
	case enums.LESS_THAN_OR_EQ:
		return i.value <= t.value, nil
	case enums.LIKE:
		return i.value == t.value, nil
	default:
		return false, errors.ErrOpComparisonNotImplemented
	}
}


//Compare(op predicate.PredicateOp, value Field) bool
	//GetType() Type
