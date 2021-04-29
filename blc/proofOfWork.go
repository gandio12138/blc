package blc

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const targetBits = 16

type ProofOfWork struct {
	Block  *Block   // 求工作量的block
	target *big.Int // 工作量难度
}

// 拼接区块属性，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
			IntToHex(pow.Block.Height),
		},
		[]byte{})
	return data
}

// IsValid 当前区块有效性验证
func (pow *ProofOfWork) IsValid() bool {
	var (
		hashInt big.Int
	)
	hashInt.SetBytes(pow.Block.Hash)
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

// Run 运行工作量证明 找难度随机数
func (pow *ProofOfWork) Run() ([]byte, int64) {
	var (
		nonce   = 0
		hashInt big.Int
		hash    [32]byte
	)
	for {
		// 拼接区块字段
		// 生成当前区块hash
		hash = sha256.Sum256(pow.prepareData(nonce))
		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	return hash[:], int64(nonce)
}

// NewProofOfWork 创建带有工作证明的区块
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBits)
	return &ProofOfWork{
		Block:  block,
		target: target,
	}
}
