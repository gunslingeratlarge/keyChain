package blc

import (
	"fmt"
	"time"
)

// Blockchain 区块链结构体
type Blockchain struct {
	Blocks []*Block
}

// CreateBlockChain 创建一个包含创世区块的区块链
func CreateBlockChain() *Blockchain {
	blockchain := new(Blockchain)
	blockchain.Blocks = append(blockchain.Blocks, CreateNewBlock(1, "genesis Block", make([]byte, 32)))
	return blockchain
}

// AppendBlock 向该区块链中添加新的区块，只需传入该区块中需存的数据，其他都会自动handle好
func (bc *Blockchain) AppendBlock(data string) {
	fmt.Println("time", time.Now().Unix())
	var len int = len(bc.Blocks)
	bc.Blocks = append(bc.Blocks, CreateNewBlock(bc.Blocks[len-1].Height+1, data, bc.Blocks[len-1].Hash))
}
