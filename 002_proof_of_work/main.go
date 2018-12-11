package main

import (
	"fmt"
	"publicChain/002_proof_of_work/blc"
)

func main() {
	block := blc.CreateNewBlock(1, "genesis", make([]byte, 32))
	res := block.Serialize()

	block = blc.DeserializeBlock(res)
	fmt.Println(block)

}
