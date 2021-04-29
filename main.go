package main

import (
	"blc/BLC"
	"fmt"
)

func main() {
	bc := BLC.NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
