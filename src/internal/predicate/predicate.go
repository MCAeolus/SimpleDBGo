package predicate

import (
	"errors"
	"fmt"
	"strconv"

	"nathan.simpledb/src/internal/enums"
	"nathan.simpledb/src/internal/interfaces"

	"nathan.simpledb/src/internal/tuple"
)

// accepts int or string
func GetOp(inp any) (enums.PredicateOp, error) {
	enumInt := -1
	var err error
	switch t := inp.(type) {
	case string:
		enumInt, err = strconv.Atoi(t)
		if err != nil {
			return -1, err
		}
	case int:
		enumInt = t	
	default:
		return -1, errors.New("unsupported type for GetOp")
	}
	return enums.PredicateOp(enumInt), nil
}

type Predicate struct {
	FieldNumber int
	Op enums.PredicateOp
	FieldValue interfaces.Field
}

func (p *Predicate) Filter(t tuple.Tuple) bool {
	// TODO
	return false;
}

func (p Predicate) String() string {
	// TODO
	return fmt.Sprintf("TODO")
} 