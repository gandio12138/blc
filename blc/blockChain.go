package blc

type BlockChain struct {
	Blocks []*Block
}

// CreateBlockChainWithGenesisBlock 创建一个区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	return &BlockChain{Blocks: []*Block{CreateGenesisBlock("Genesis Block ......")}}
}

// AddBlockToBlockChain 添加区块到区块链
func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	blc.Blocks = append(blc.Blocks, NewBlock(data, height, prevHash))
}
