package main

import (
	"fmt"
)

type Block struct {
	hash [32]byte
	data []byte
	prev *Block
}

func main() {
	blockChain := mkGenesisBlock()
	blockChain = addToBlock(blockChain, []byte("Send Charlie 11 Dollars"))
	blockChain = addToBlock(blockChain, []byte("Send Willis 3.5 Dollars"))
	blockChain = addToBlock(blockChain, []byte("Bill Sarah for 6 Dollars"))
	blockChain = addToBlock(blockChain, []byte("Bill Philip for 79 Dollars"))
	blockChain = addToBlock(blockChain, []byte("Send Jeffrey 18 Dollars"))
	blockChain = addToBlock(blockChain, []byte("Carl is now banned for life"))
	blockChain = addToBlock(blockChain, []byte("Bill Ill Bill 56 Dollars to Jill"))
	blockChain = addToBlock(blockChain, []byte("Websocket is a pain"))
	blockChain = addToBlock(blockChain, []byte("Send Betsy 17 Dollars"))
	blockChain = addToBlock(blockChain, []byte("32 byte hash"))
	blockChain = addToBlock(blockChain, []byte("hash of prev block"))

	printBlockChain(blockChain)
}

func verifyBlockChain(blockChain *Block) bool {
	if blockChain.prev == nil {
		return true
	}
	if blockChain.prev.prev == nil {
		return blockChain.hash == mkHash(*blockChain.prev, blockChain.data)
	}
	return blockChain.hash == mkHash(*blockChain.prev, blockChain.data) && verifyBlockChain(blockChain.prev)
}

func printBlockChain(blockChain *Block) {
	if blockChain != nil {
		fmt.Printf("-----BLOCK-----\n")
		fmt.Printf("Hash: %x\n", blockChain.hash)
		fmt.Printf("Data: %s\n", blockChain.data)
		printBlockChain(blockChain.prev)
	} else {
		fmt.Printf("-----END-----\n")
	}
}
