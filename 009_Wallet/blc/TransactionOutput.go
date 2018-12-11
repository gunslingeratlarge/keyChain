package blc

// 交易的输出，包含公钥，如果有人想用这里面的钱，必须使用它对应的私钥发送对应的签名
type TxOutput struct {
	Value        int64
	ScriptPubKey string
}

func (output *TxOutput) canBeUnlocked(unlockingData string) bool {
	return unlockingData == output.ScriptPubKey
}
