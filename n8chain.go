package main

import (
	"crypto/sha256"
	"fmt"
) 

type Block struct{
	hash [32]byte
	data []byte
	prev *Block
}

