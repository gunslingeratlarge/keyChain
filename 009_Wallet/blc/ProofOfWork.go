package blc

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const difficulty = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork 创建一个pow对象，该对象可以实现挖矿。初始化对象的目标hash值（target）
func NewProofOfWork(b *Block) *ProofOfWork {
	pow := ProofOfWork{b, nil}
	// 0000 0001  d = 2  shift 8 - 2 = 6位
	target := big.NewInt(1)
	target = target.Lsh(target, 256-difficulty)
	pow.target = target
	return &pow
}

//Run 执行挖矿：产出正确的哈希值和nonce
func (pow *ProofOfWork) Run() ([]byte, int) {
	block := pow.block

	var nonce int
	var hash [32]byte
	var hashInt big.Int
	// 循环直到找到正确的nonce为止
	for {
		hash = sha256.Sum256(prepareData(block, nonce))
		hashInt.SetBytes(hash[:])
		// 若target > hashInt 说明0比target多，满足难度要求
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}

	return hash[:], nonce

}

func prepareData(b *Block, nonce int) []byte {
	dataJoined := bytes.Join([][]byte{
		IntToByteSlice(b.Height),
		b.HashTransactions(),
		b.PrevHash,
		IntToByteSlice(b.TimeStamp),
		IntToByteSlice(int64(nonce))}, []byte{'-'})

	return dataJoined
}

// BlockValidate 判断区块的hash是否合法
func (pow *ProofOfWork) BlockValidate() bool {
	b := pow.block
	hash := sha256.Sum256(prepareData(b, b.Nonce))
	hashInt := new(big.Int).SetBytes(hash[:])
	if pow.target.Cmp(hashInt) == 1 {
		return true
	}
	return false

}
