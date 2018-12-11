package blc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Wallet struct {
	privateKey ecdsa.PrivateKey
	publicKey  []byte
}

//生成一个密钥对
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	publicKey := append(privateKey.X.Bytes(), privateKey.Y.Bytes()...)
	return *privateKey, publicKey
}

//新钱包对象
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	return new(Wallet{privateKey, publicKey})
}
