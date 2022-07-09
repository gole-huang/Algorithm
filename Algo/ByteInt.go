package Algo

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuf, binary.BigEndian, x)
	if err != nil {
		return nil
	}
	return bytesBuf.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b)
	var x int32
	err := binary.Read(bytesBuf, binary.BigEndian, &x)
	if err != nil {
		return 0
	}
	return int(x)
}
