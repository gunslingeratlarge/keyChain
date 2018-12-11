package main

import (
	"publicChain/006_Iterator/blc"
)

func main() {
	blockchain := blc.CreateBlockChain()
	defer blockchain.DB.Close()

	blockchain.AppendBlock("second block")
	blockchain.AppendBlock("third block")
	blockchain.PrintChain()
}
