package BLC

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
Usage:
  addBlock -data BLOCK_DATA    add a block to the blockchain
  printChain                   print all the blocks of the blockchain
`

type CLI struct {
	BC *BlockChain
}

// 打印命令提示
func (cli *CLI) printUsage() {
	fmt.Print(usage)
}

// 检测命令个数
func (cli *CLI) valDataArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run 运行命令工具
func (cli *CLI) Run() {
	cli.valDataArgs()
	// 获取cmd 命令
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")
	switch os.Args[1] { // 分支处理
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.BC.AddBlock(*addBlockData)
	}
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
