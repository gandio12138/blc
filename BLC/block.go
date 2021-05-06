package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	TimeStamp     int64          // 时间戳
	Nonce         int64          // 难度值
	Transactions  []*Transaction // 当前区块交易
	PrevBlockHash []byte         // 前一个区块hash
	Hash          []byte         // 当前区块hash
}

func (b *Block) Serialize() []byte {
	var (
		result  bytes.Buffer
		encoder = gob.NewEncoder(&result)
	)
	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(data []byte) *Block {
	var (
		block   Block
		decoder = gob.NewDecoder(bytes.NewReader(data))
	)
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return &block
}

// HashTransactions 计算区块中所有交易的hash
func (b *Block) HashTransactions() []byte {
	var (
		txHashes [][]byte
		txHash   [32]byte
	)
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Id)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Nonce:         0,
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

