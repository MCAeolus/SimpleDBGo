package field

import (
	"io"
	"strings"

	"nathan.simpledb/src/internal/enums"
	"nathan.simpledb/src/internal/errors"
	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/util"
)

type StringField struct {
	value string
	maxSize int
}

func NewStringField(value string, maxSize int) StringField {
	if len(value) > maxSize {
		value = value[:maxSize]
	}
	return StringField{value: value, maxSize: maxSize}
}

func (s StringField) GetValue() string {
	return s.value
}

func (s StringField) String() string {
	return s.value
}

// string field writes the length (4 bytes) + the string
// remained is zero padded
func (s StringField) Serialize(writer io.Writer) error {
	sval := s.value[:s.maxSize]
	overflow := s.maxSize - len(sval)
	//write in length
	_, err := writer.Write(util.IntToBytes(int32(len(sval))))
	if err != nil {
		return err
	}
	//write in string
	_, err = writer.Write([]byte(sval))
	if err != nil {

		return err //TODO: this is problematic
	}
	for overflow > 0 {
		_, err := writer.Write([]byte{0})
		if err != nil {
			return err // also problematic
		}
		overflow--
	}
	return nil
}

func (s StringField) GetType() interfaces.Type {
	return STRING_TYPE
}

func (s StringField) Compare(op enums.PredicateOp, value interfaces.Field) (bool, error) {
	t, ok := value.(StringField) 
	if !ok {
		return false, errors.ErrIllegalCast
	}
	strCmp := strings.Compare(s.value, t.value)
	switch (op) {
	case enums.EQUALS:
		return strCmp == 0, nil
	case enums.NOT_EQUALS:
		return strCmp != 0, nil
	case enums.GREATER_THAN:
		return strCmp > 0, nil
	case enums.GREATER_THAN_OR_EQ:
		return strCmp >= 0, nil
	case enums.LESS_THAN:
		return strCmp < 0, nil
	case enums.LESS_THAN_OR_EQ:
		return strCmp <= 0, nil
	case enums.LIKE:
		return strings.Contains(s.value, t.value), nil
	default:
		return false, errors.ErrOpComparisonNotImplemented
	}
}

