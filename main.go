package main

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
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

func addBlock(old)
