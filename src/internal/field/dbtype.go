package field

import (
	"fmt"
	"io"

	"nathan.simpledb/src/internal/interfaces"
	"nathan.simpledb/src/internal/util"
)

var (
	INT_TYPE = IntType{}
	STRING_TYPE = StrType{}
)

const STRING_LEN = 128

type IntType struct {}

func (i IntType) GetLen() int {
	return 4
}

func (i IntType) Parse(reader io.Reader) interfaces.Field {
	b := ReadN(4, reader)
	v := util.BytesToInt(b)
	return NewIntField(v)
}

func (i IntType) String() string {
	return "INTEGER"
}


type StrType struct  {}

func (s StrType) GetLen() int {
	return STRING_LEN + 4;
}

func (s StrType) Parse(reader io.Reader) interfaces.Field {
	// read length in
	blen := ReadN(4, reader)
	strLen := util.BytesToInt(blen)
	bstr := ReadN(int(strLen), reader)
	_ = ReadN(STRING_LEN-int(strLen), reader)
	return NewStringField(string(bstr), STRING_LEN)

}

func (s StrType) String() string {
	return "STRING"
}

func ReadN(n int, reader io.Reader) []byte {
	b := make([]byte, n)
	o, err := io.ReadFull(reader, b)
	if err != nil {
		panic(fmt.Sprintf("could not read in %d bytes! (read in %d)", n, o))
	}
	return b
}