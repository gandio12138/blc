package main

import "blc/BLC"

func main() {
	bc := BLC.NewBlockChain()
	defer bc.Close()
	cli := BLC.CLI{BC: bc}
	cli.Run()
}
