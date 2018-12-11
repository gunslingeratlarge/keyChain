package blc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 每次挖矿奖励10个币
const reward = 10

// UTXO
type Transaction struct {
	TxHash  []byte
	Inputs  []*TxInput
	Outputs []*TxOutput
}

// 挖矿时无中生有获得的币
func CoinBaseTransaction(data, to string) *Transaction {
	if data == "" {
		data = "mining reward to " + to
	}

	input := TxInput{[]byte{}, -1, data}
	output := TxOutput{reward, to}
	tx := Transaction{nil, []*TxInput{&input}, []*TxOutput{&output}}
	return &tx
}

// 设置当前交易的hash值
func (tx *Transaction) setHash() {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	// 返回32位数组
	hash := sha256.Sum256(result.Bytes())
	//返回切片
	tx.TxHash = hash[:]

}

// 判断该交易是否是一个coinBase交易
func (tx *Transaction) isCoinBase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].TxHash) == 1 && tx.Inputs[0].OutputIndex == -1
}
