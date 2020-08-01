package main

import (
	"blockchain/block"
)

func main() {
	blockchain := block.InitBlockChain()
	blockchain.AddBlock("First Block")
	blockchain.AddBlock("Second Block")
	blockchain.ShowInfo()

}
