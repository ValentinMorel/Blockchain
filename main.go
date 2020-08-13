package main

import (
	"blockchain/blockchain"
	"blockchain/cli"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("Usage : ./main [api/cli] **kwargs")
	}

	switch os.Args[1] {
	case "cli":
		defer os.Exit(0)
		chain := blockchain.InitBlockChain()
		defer chain.Database.Close()
		cli := cli.CommandLine{chain}
		cli.Run()

	case "api":
		log.Println("Running api side.")
		//TODO : api backend

	default:
		log.Fatal("Not a valid argument.")
	}

}
