package main

import (
	"fmt"

	"github.com/minhajuddinkhan/huffman"
)

func main() {

	tree := huffman.NewHuffmanTree("whats your nAme?")

	var result string
	err := tree.Encode(&result)
	if err != nil {
		panic(err)
	}

	decoded, _ := tree.Decode(result)
	fmt.Println(decoded)
}
