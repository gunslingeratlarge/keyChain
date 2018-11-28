package blc

import (
	"fmt"
	"time"
)

// Block 区块数据结构
type Block struct {
	// 当前区块高度
	Height int64
	// 数据
	data []byte
	// 前hash
	PrevHash []byte
	// 后区块hash
	Hash []byte
	// 时间戳：区块生成的时间
	timeStamp int64
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
