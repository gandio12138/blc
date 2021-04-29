package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Hash          []byte // 本区块hash
	PrevBlockHash []byte // 前一个区块hash
	Data          []byte // 交易数据
	TimeStamp     int64  // 时间戳
	Height        int64  // 区块链高度
	Nonce         int64  // 工作量证明
}

// SetHash 设置当前区块hash
func (block *Block) SetHash() {
	// 当前区块高度、时间戳转换成字节数组
	heightBytes := IntToHex(block.Height)
	timeStamp := []byte(strconv.FormatInt(block.TimeStamp, 2))
	// 拼接属性
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeStamp,
		block.Hash,
	}, []byte{})
	currHash := sha256.Sum256(blockBytes)
	fmt.Printf("curr block hash value : %v\n", currHash)
	block.Hash = blockBytes
}

// NewBlock 创建新区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		TimeStamp:     time.Now().Unix(),
		Height:        height,
	}
	pow := NewProofOfWork(block)
	block.Hash, block.Nonce = pow.Run()
	fmt.Printf("\r%d-%x\n", block.Nonce, block.Hash)
	return block
}

// CreateGenesisBlock 创建创世块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
