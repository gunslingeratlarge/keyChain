package blc

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"log"
)

const DB_NAME = "blockchain.db"
const TABLE_NAME = "blocks"

type Blockchain struct {
	Last []byte
	DB   *bolt.DB
}

// CreateBlockChain 创建一个包含创世区块的区块链
func CreateBlockChain() *Blockchain {
	blockchain := new(Blockchain)
	db, err := bolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	blockchain.DB = db
	err = db.Update(func(tx *bolt.Tx) error {
		table, err := tx.CreateBucket([]byte(TABLE_NAME))
		if err != nil {
			log.Panic(err)
		}

		block := CreateNewBlock(1, "genesis Block", make([]byte, 32))
		err = table.Put(block.Hash, block.Serialize())
		if err != nil {
			panic(err)
		}
		err = table.Put([]byte("Last"), block.Hash)
		blockchain.Last = block.Hash
		return nil
	})

	return blockchain
}

// AppendBlock 向该区块链中添加新的区块，只需传入该区块中需存的数据，其他都会自动handle好
func (bc *Blockchain) AppendBlock(data string) {
	db := bc.DB
	err := db.Update(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TABLE_NAME))
		if table == nil {
			return errors.New("table does not exist")
		}
		prevBytes := table.Get(table.Get([]byte("Last")))
		prevBlock := DeserializeBlock(prevBytes)
		newBlock := CreateNewBlock(prevBlock.Height+1, data, prevBlock.Hash)
		err := table.Put([]byte(newBlock.Hash), newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = table.Put([]byte("Last"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.Last = newBlock.Hash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
