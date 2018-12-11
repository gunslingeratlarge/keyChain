package blc

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Block 区块数据结构
type Block struct {
	// 当前区块高度
	Height int64
	// 数据
	Data []byte
	// 前hash
	PrevHash []byte
	// 后区块hash
	Hash []byte
	// 时间戳：区块生成的时间
	TimeStamp int64
	// 随机数
	Nonce int
}

// CreateNewBlock 创建新的区块： 将这个方法绑定到Block类型上
func CreateNewBlock(height int64, data string, prevHash []byte) *Block {

	block := &Block{height, []byte(data), prevHash, nil, time.Now().UnixNano(), 0}
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	fmt.Printf("%x\n", hash)
	return block
}

// 将区块序列化成字节数组
func (block *Block) Serialize() []byte {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化
func DeserializeBlock(blockBytes []byte) *Block {

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
