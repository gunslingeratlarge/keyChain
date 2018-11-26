package blc

import "time"

import "bytes"
import "crypto/sha256"

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
}

func (b *Block) setHash() {
	heightBytes := IntToByteSlice(b.Height)
	timeBytes := IntToByteSlice(b.timeStamp)
	blockBytes := bytes.Join([][]byte{b.data, heightBytes, timeBytes, b.PrevHash}, []byte{'-'})
	// 返回的是固定长度32的数组
	hash := sha256.Sum256(blockBytes)
	//取切片并赋给hash
	b.Hash = hash[:]
}

// CreateNewBlock 创建新的区块： 将这个方法绑定到Block类型上
func CreateNewBlock(height int64, data string, prevHash []byte) *Block {

	block := &Block{height, []byte(data), prevHash, nil, time.Now().UnixNano()}
	block.setHash()
	return block
}
