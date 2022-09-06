package helper

import (
	"bytes"
	"encoding/binary"
)

func Uint16ToBytes(n uint16) []byte {
	x := n

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}

	return bytesBuffer.Bytes()
}

func BytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)

	var x uint16
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		panic(err)
	}

	return x
}

func Uint32ToBytes(n uint32) []byte {
	x := n

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
	return bytesBuffer.Bytes()
}

func BytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x uint32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		panic(err)
	}

	return x
}
