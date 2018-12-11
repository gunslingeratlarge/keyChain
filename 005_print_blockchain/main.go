package main

import (
	"publicChain/005_print_blockchain/blc"
)

func main() {
	blockchain := blc.CreateBlockChain()
	defer blockchain.DB.Close()

	blockchain.AppendBlock("second block")
	blockchain.AppendBlock("third block")
	blockchain.PrintChain()
}
