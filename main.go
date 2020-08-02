package main

import (
	"blockchain/block"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type CommandLine struct {
	blockchain *block.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		blck := iter.Next()

		fmt.Printf("Prev. hash: %x\n", blck.PrevHash)
		fmt.Printf("Data: %s\n", blck.Data)
		fmt.Printf("Hash: %x\n", blck.Hash)
		pow := block.NewProof(blck)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(blck.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		block.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		block.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := block.InitBlockChain()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.run()
}
