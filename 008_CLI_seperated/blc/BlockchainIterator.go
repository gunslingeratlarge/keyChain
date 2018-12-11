package blc

import (
	"bytes"
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (iterator *BlockchainIterator) Next() *Block {
	var block *Block
	err := iterator.DB.View(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TABLE_NAME))
		blockBytes := table.Get(iterator.CurrentHash)
		block = DeserializeBlock(blockBytes)
		iterator.CurrentHash = block.PrevHash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}

func (iterator *BlockchainIterator) HasNext() (hasNext bool) {
	return !bytes.Equal(make([]byte, 32), iterator.CurrentHash)
}
