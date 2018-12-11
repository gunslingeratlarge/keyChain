package blc

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

const DB_NAME = "blockchain.db"
const TABLE_NAME = "blocks"

type Blockchain struct {
	Last []byte
	DB   *bolt.DB
}

// CreateBlockChain 创建一个包含创世区块的区块链
func CreateBlockChain(address string) *Blockchain {
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
		blockchainTx := CoinBaseTransaction("", address)
		block := CreateNewBlock(1, []*Transaction{blockchainTx}, make([]byte, 32))
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
func (bc *Blockchain) AppendBlock(txs []*Transaction) {
	db := bc.DB
	err := db.Update(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(TABLE_NAME))
		if table == nil {
			return errors.New("table does not exist")
		}
		prevBytes := table.Get(table.Get([]byte("Last")))
		prevBlock := DeserializeBlock(prevBytes)
		newBlock := CreateNewBlock(prevBlock.Height+1, txs, prevBlock.Hash)
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

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.Last, bc.DB}
}

func (bc *Blockchain) PrintChain() {

	iterator := bc.Iterator()
	for iterator.HasNext() {
		block := iterator.Next()
		fmt.Printf("height:%d\n\nprevHash:%x\nhash:%x\ntimestamp:%s\nnonce:%d\n", block.Height, block.PrevHash,
			block.Hash, time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"), block.Nonce)
		fmt.Print("Transactions:\n")
		for _, tx := range block.Txs {
			fmt.Printf("\tTxHash:%x\n", tx.TxHash)
			fmt.Printf("\tTxInputs:\n")
			for _, input := range tx.Inputs {
				fmt.Printf("\t\tTxHash:%x\n", input.TxHash)
				fmt.Println("\t\toutputIndex:", input.OutputIndex)
				fmt.Println("\t\taddress:", input.ScriptSig)
			}
			fmt.Printf("\tTxOutputs:\n")
			for _, output := range tx.Outputs {
				fmt.Println("\t\tValue:", output.Value)
				fmt.Printf("\t\taddress:%s\n", output.ScriptPubKey)
			}
		}

		fmt.Println("========================")
	}
}

// 判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(DB_NAME); os.IsNotExist(err) {
		return false
	}

	return true
}

// 打开数据库，返回最后一个区块哈希，构造blockchain对象
func BlockchainObject() *Blockchain {

	db, err := bolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var last []byte

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(TABLE_NAME))

		if b != nil {
			// 读取最新区块的Hash
			last = b.Get([]byte("Last"))

		}

		return nil
	})

	return &Blockchain{last, db}
}

/*======================下方是绑定到区块链上的交易相关的方法===================*/

// 找到所有没有完全花费完的交易 （由于区块的遍历是从后到前的，因此最先找到的未花费output是真的没有花费，而已花费的input的加入正好在遍历下一个区块获得未花费output之前）
func (bc *Blockchain) FindUnspentTransactions(address string) []*Transaction {
	// 未（完全）花费的交易
	var unSpentTxs []*Transaction
	// 已经花费了的输出
	spentTxOutputs := make(map[string][]int)
	iterator := bc.Iterator()
	for iterator.HasNext() {
		block := iterator.Next()
		for _, tx := range block.Txs {
			// 对于每个交易，先处理其所有的output，如果有关于当前address的未花费交易，则将其加入unSpentTxs

			// 跳过所有花费的output
		Outputs:
			for outIndex, out := range tx.Outputs {
				// 转成string方便作为map的key值
				txId := hex.EncodeToString(tx.TxHash)
				if spentTxOutputs[txId] != nil {
					for _, spentIndex := range spentTxOutputs[txId] {
						if outIndex == spentIndex {
							continue Outputs
						}
					}
				}

				if out.canBeUnlocked(address) {
					unSpentTxs = append(unSpentTxs, tx)
				}
			}

			// 再处理所有的input，如果该input是当前address的，则该input对应的上一个output就已经被花费，所以加入spentTxOutputs中
			if !tx.isCoinBase() {
				for _, in := range tx.Inputs {
					// 这个input是被当前用户所使用了，那么所对应的一定有一个属于该用户的output被使用了
					if in.canUnlockOutput(address) {
						inTxId := hex.EncodeToString(in.TxHash)
						spentTxOutputs[inTxId] = append(spentTxOutputs[inTxId], in.OutputIndex)
					}
				}
			}
		}
	}
	return unSpentTxs
}

// 寻找所有未花费交易，该方法是用来计算余额的（这样才需要遍历所有节点）
func (bc *Blockchain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput
	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Outputs {
			if out.canBeUnlocked(address) {
				UTXOs = append(UTXOs, *out)
			}
		}
	}
	return UTXOs
}

//找到属于这个address的至少大于amount数量的未花费输出：txHash的字符串表示->属于该交易的output的index
func (bc *Blockchain) FindEnoughUTXOFor(amount int64, address string) (accumulated int64, outputs map[string][]int) {
	accumulated = 0
	UTXs := bc.FindUnspentTransactions(address)
	UTXOs := make(map[string][]int)
Work:
	for _, tx := range UTXs {
		txID := hex.EncodeToString(tx.TxHash)

		for outIdx, out := range tx.Outputs {
			if out.canBeUnlocked(address) && accumulated < amount {
				accumulated += out.Value
				UTXOs[txID] = append(UTXOs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}
	return accumulated, UTXOs
}

func NewUTXOTransaction(from, to string, amount int64, bc *Blockchain) *Transaction {
	var inputs []*TxInput
	var outputs []*TxOutput

	acc, validOutputs := bc.FindEnoughUTXOFor(int64(amount), from)

	if acc < amount {
		log.Panic("ERROR: 账户余额不足")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic("解码交易哈希失败")
		}
		for _, out := range outs {
			input := &TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// 新建交易输出
	outputs = append(outputs, &TxOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, &TxOutput{acc - amount, from}) // 如果有找零，创建新的output
	}

	tx := Transaction{nil, inputs, outputs}
	tx.setHash()

	return &tx
}
