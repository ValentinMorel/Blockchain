package main

import (
	"blockchain/block"
	"fmt"
)

func main() {
	blockchain := block.InitBlockChain()
	fmt.Printf("Blockchain : %x", blockchain)

}
