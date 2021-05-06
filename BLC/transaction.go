package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

const (
	subsidy = 10
)

type (
	TXInput struct {
		Taxied    []byte // 一个交易输入引用了之前一笔交易的一个输出 id 表明是之前的哪笔交易
		Volute    int64  // 一笔交易可能有多个输入 Volute 为输入索引
		ScriptSig string // 提供解锁输出 Taxied: Volute 的数据
	}
	TXOutPut struct {
		Value        int64  // 交易币币的数量
		ScriptPubKey string // 输出锁定
	}
	Transaction struct { // 一个交易结构
		Id   []byte
		Vin  []TXInput
		VOut []TXOutPut
	}
)

// IsCoinbase 判断是否是 coinbase 交易
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Taxied) == 0 && tx.Vin[0].Volute == -1
}

func (tx *Transaction) SetId() {
	var (
		encoded bytes.Buffer
		hash    [32]byte
	)
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.Id = hash[:]
}

// CanUnlockOutputWith 这里的 unlockingData 可以理解为地址
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutPut) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewUTXOTransaction(from, to string, amount int64, bc *Blockchain) *Transaction {
	var (
		input  []TXInput
		output []TXOutPut
	)
	acc, validOutputs :
}
