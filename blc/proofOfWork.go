package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const (
	targetBits = 24
	maxNonce   = math.MaxInt64
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBits)
	return &ProofOfWork{
		block:  b,
		target: target,
	}
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	return bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntTOHex(pow.block.TimeStamp),
		IntTOHex(int64(targetBits)),
		IntTOHex(nonce),
	}, []byte{})
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var (
		hashInt big.Int
		hash    [32]byte
		nonce   = int64(0)
	)
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x\n", hash)
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var (
		hashInt big.Int
	)
	hash := sha256.Sum256(pow.prepareData(pow.block.Nonce))
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
