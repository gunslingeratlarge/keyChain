package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"publicChain/004_persistence_create_genesis_block/blc"
)

func main() {
	blockchain := blc.CreateBlockChain()
	defer blockchain.DB.Close()

	blockchain.AppendBlock("second block")
	blockchain.AppendBlock("third block")
	err := blockchain.DB.View(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(blc.TABLE_NAME))
		hash := table.Get([]byte("Last"))
		block := table.Get(hash)
		fmt.Printf("%s", blc.DeserializeBlock(block).Data)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
