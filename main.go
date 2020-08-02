package main

import (
	"blockchain/block"
	"os"
)

func main() {
	defer os.Exit(0)
	chain := block.InitBlockChain()
	defer chain.Database.Close()

	cli := block.CommandLine{chain}
	cli.Run()
}
