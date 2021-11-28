package main

import (
	"fmt"
)

func main() {
	chain := blockchain.initBlockchain()

	chain.AddBlock("The first block")
	chain.AddBlock("The second block")
	chain.AddBlock("The third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
	}
}
