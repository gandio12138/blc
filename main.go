package main

import (
	"blc/blc"
	"fmt"
)

func main() {
	blockChain := blc.CreateBlockChainWithGenesisBlock()
	prevBlock1 := blockChain.Blocks[len(blockChain.Blocks)-1]
	blockChain.AddBlockToBlockChain("first block", prevBlock1.Height, prevBlock1.Hash)
	prevBlock2 := blockChain.Blocks[len(blockChain.Blocks)-1]
	blockChain.AddBlockToBlockChain("first block", prevBlock2.Height, prevBlock2.Hash)
	prevBlock3 := blockChain.Blocks[len(blockChain.Blocks)-1]
	blockChain.AddBlockToBlockChain("first block", prevBlock3.Height, prevBlock3.Hash)
	fmt.Println(blockChain)
}
