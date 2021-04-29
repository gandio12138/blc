package BLC

type BlockChain struct {
	Blocks []*Block
}

// NewBlockChain 创建一个有创世块的区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{NewGenesisBlock()}}
}

// AddBlock 向链中添加块
func (bc *BlockChain) AddBlock(data string) {
	bc.Blocks = append(bc.Blocks, NewBlock(data, bc.Blocks[len(bc.Blocks)-1].Hash))
}
