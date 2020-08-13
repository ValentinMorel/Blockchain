package main

import (
	"blockchain/blockchain"
	"blockchain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := cli.CommandLine{chain}
	cli.Run()
}
