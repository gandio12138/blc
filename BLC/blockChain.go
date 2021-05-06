package BLC

import "github.com/boltdb/bolt"

const (
	dbFile              = "blockchain.db"
	blocksBucket        = "blocks"
	genesisCoinbaseData = "The Times 30/April/2021 Chancellor on brink of second bailout for banks"
)

type (
	BlockchainIterator struct {
		currentHash []byte
		db          *bolt.DB
	}
	Blockchain struct {
		tip []byte
		db  *bolt.DB
	}
)

func (bc *Blockchain) MineBlock(transactions []*Transaction) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	newBlock := NewBlock(transactions, lastHash)
	err1 := bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err = b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			panic(err)
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
	if err1 != nil {
		panic(err)
	}
}
