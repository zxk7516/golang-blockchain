package main

import (
	"fmt"

	"github.com/zxk7516/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Thrid Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("previous Hash: %X\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %X\n", block.Hash)
	}

}
