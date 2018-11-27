package main

import (
	"fmt"
	"publicchain/001_basic_prototype/blc"
)

func main() {
	// block := blc.CreateNewBlock(257, "genesis block", make([]byte, 32))
	// fmt.Println(block)
	bchain := blc.CreateBlockChain()
	bchain.AppendBlock("send zs 100 yuan")
	bchain.AppendBlock("send ls  200 yuan")
	bchain.AppendBlock("send ls  200 yuan")

	fmt.Println(bchain)
	fmt.Println(bchain.Blocks)

	for _, v := range bchain.Blocks {
		fmt.Println(v)
	}

}
