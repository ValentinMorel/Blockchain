package main

import (
	"blockchain/block"
)

func main() {

	blockchain := block.InitBlockChain()
	blockchain.AddBlock("First Block after Genesis")
	blockchain.AddBlock("Second Block after")
	blockchain.ShowInfo()

}
