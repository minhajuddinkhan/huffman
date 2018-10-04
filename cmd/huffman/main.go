package main

import (
	"fmt"

	"github.com/minhajuddinkhan/compression"
)

func main() {

	tree := compression.NewHuffmanTree("abbccc")
	encoded, err := tree.Encode()
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded) //000101111
}
