package main

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"
)

type block struct {
	Index     int
	Timestamp string
	Hash      string
	Parent    string
}

var Minichain []Block

var mutex = &sync.Mutex{}

// var input string = flag.StringVar(&input, "i", "fubu for us by us", "a string")

// var height int
// flag.IntVar(&height, "n", 12, "a number")

func calcHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Parent
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func addBlock(previousBlock Block) Block {

	var newBlock block

	t := time.Now()

	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.ParentHash = previousBlock.hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
