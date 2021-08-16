package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const (
	dbFilePath   = "blockChain.db"
	blocksBucket = "blocks"
)

type (
	BlockChain struct {
		tip []byte
		db  *bolt.DB
	}
	BlockChainIterator struct {
		currentHash []byte
		db          *bolt.DB
	}
)

// NewBlockChain 创建一个有创世块的区块链
func NewBlockChain() *BlockChain {
	var tip []byte
	// 打开一个 BoltDB 文件
	db, err := bolt.Open(dbFilePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		if bucket == nil { // 以创世块初始化db
			fmt.Println("No existing blockChain found. Create a new one .....")
			genesis := NewGenesisBlock()
			bucket, err = tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				panic(err)
			}
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				panic(err)
			}
			err = bucket.Put([]byte("1"), genesis.Hash)
			if err != nil {
				panic(err)
			}
			tip = genesis.Hash
		} else { // 有直接拿创世块
			tip = bucket.Get([]byte("1"))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// 初始化区块链
	bc := BlockChain{
		tip: tip,
		db:  db,
	}
	return &bc
}

// AddBlock 加入区块时需要同时持久化到数据库
func (bc *BlockChain) AddBlock(data string) {
	var (
		lastHash []byte
	)
	// 获取最后一个块的hash用于生成新块的hash
	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("1"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	newBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err = bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			panic(err)
		}
		err = bucket.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{
		currentHash: bc.tip,
		db:          bc.db,
	}
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		block = DeserializeBlock(b.Get(i.currentHash))
		return nil
	})
	if err != nil {
		panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}

func (bc *BlockChain) Close() {
	if err := bc.db.Close(); err != nil {
		panic(err)
	}
}
