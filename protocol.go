package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader       = "KeyWord"
	ConstHeaderLength = len(ConstHeader)
	ConstDataLength   = 4
)

// 封包
func Pack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

// 解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			msgLen := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstDataLength])
			if length < i+ConstHeaderLength+ConstDataLength+msgLen {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstDataLength : i+ConstHeaderLength+ConstDataLength+msgLen]
			readerChannel <- data

			i += ConstHeaderLength + ConstDataLength + msgLen - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	var x int32
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &x)
	return int(x)
}
