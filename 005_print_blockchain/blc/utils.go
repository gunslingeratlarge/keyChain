package blc

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToByteSlice 将int64转换为字节数组(64位，每8位取一个数作为一个byte放进去）
func IntToByteSlice(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
