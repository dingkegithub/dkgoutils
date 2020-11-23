package types

import (
	"bytes"
	"encoding/binary"
)

func ByteToInt(b []byte) int {
	buff := bytes.NewBuffer(b)

	var data int64
	binary.Read(buff, binary.BigEndian, &data)
	return int(data)
}

func IntToByte(num int) []byte {
	data := int64(num)
	byteBuf := bytes.NewBuffer([]byte{})
	binary.Write(byteBuf, binary.BigEndian, data)
	return byteBuf.Bytes()
}
