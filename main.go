package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"golang.org/x/sync/semaphore"
)

// BIIITCONNNNNNEEEECCCCT!

type Block struct {
	Index     int
	Timestamp string
	Hash      string
	Parent    *Block
	Nonce     int
}

var sem = semaphore.NewWeighted(21)

func calcHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Parent.(string)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func addBlock(previousBlock Block, Nonce int) (Block, error) {

	var newBlock block

	t := time.Now()

	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.ParentHash = previousBlock.hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func validateBlock(newBlock Block, parentBlock Block) bool {
	if parentBlock.Index+1 != newBlock.Index {
		return false
	}

	if parentBlock.Hash != newBlock.ParentHash {
		return false
	}

	if calcHash(newBlock) != newBlock.hash {
		return false
	}

	return true
}
