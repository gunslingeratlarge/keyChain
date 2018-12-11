package blc

// 每一个transactionInput必须引用一个前面交易的output
type TxInput struct {
	//前面交易的txHash
	TxHash      []byte
	OutputIndex int
	// 可以先理解为用户名。这个input想要有权去动用上面的output，该签名就必须是上面的output的公钥所对应的
	ScriptSig string
}

func (input *TxInput) canUnlockOutput(unlockingData string) bool {
	return unlockingData == input.ScriptSig
}
