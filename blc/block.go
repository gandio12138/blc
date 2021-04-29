package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp     int64
	Nonce         int64
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// NewGenesisBlock 生成创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// SetHash 设置当前块hash
func (blc *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(blc.TimeStamp, 10))
	headers := bytes.Join(
		[][]byte{
			blc.PrevBlockHash,
			blc.Data,
			timestamp,
		}, []byte{},
	)
	hash := sha256.Sum256(headers)
	blc.Hash = hash[:]
}
