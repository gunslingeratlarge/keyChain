package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"publicChain/003_bolt_db/blc"
)

func main() {
	block := blc.CreateNewBlock(1, "genesis", make([]byte, 32))
	fmt.Println(block)

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic(err)
			}
		}
		err = b.Put(block.Hash, block.Serialize())
		if err != nil {
			log.Panic(err)
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte("blocks"))
		b := table.Get(block.Hash)
		block := blc.DeserializeBlock(b)
		fmt.Println(block)
		return nil
	})
}
