package util

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.AppendUint32(b, uint32(i))
	return b
}

// automatically truncates
func BytesToInt(b []byte) int32 {
	b = b[:4]
	var i int32
	err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &i)
	if err != nil {
		// TODO: error?
		return 0
	}
	return i
}