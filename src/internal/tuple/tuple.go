package tuple

import (
	"fmt"
	"strings"

	"nathan.simpledb/src/internal/interfaces"
)

/** taken from SimpelDB comments
 * Tuple maintains information about the contents of a tuple.
 * Tuples have a specified schema specified by a TupleDesc object and contain
 * Field objects with the data for each field.
 */
type Tuple struct {
	fields []interfaces.Field
	description TupleDesc
	recordId RecordId
}

func NewTuple(td TupleDesc) Tuple {
	return Tuple{
		description: td,
		fields: make([]interfaces.Field, td.numFields),
	}
}

// representing the schema of this tuple
func (t *Tuple) GetTupleDesc() TupleDesc {
	return t.description;
}

// location of tuple on disk or nil
func (t *Tuple) GetRecordId() RecordId {
	return t.recordId
}

func (t *Tuple) SetRecordId(rid RecordId) {
	t.recordId = rid
}

func (t *Tuple) SetField(index int, field interfaces.Field) {
	if t.description.numFields <= index {
		//do nothing? panic?
		panic("too large index passed to tuple")
	}
	if t.description.types[index] != field.GetType() {
		// do nothing? error? panic???
		panic("field type does not match")
	}
	//todo is this okay? do we need to clone?
	t.fields[index] = field
}

func (t *Tuple) GetField(index int) interfaces.Field {
	if t.description.numFields <= index {
		panic("too large index passed to tuple")
	}
	return t.fields[index]
}

func (t *Tuple) String() string {
	var sb strings.Builder
	for ix, f := range t.fields {
		delim := ""
		if ix > 0 {
			delim = "\t"
		}
		sb.WriteString(fmt.Sprintf("%s%s", delim, f)) // note: f implements String()
	}
	return sb.String()
}