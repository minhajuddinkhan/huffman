package main

import (
	"fmt"

	"github.com/minhajuddinkhan/huffman"
)

func main() {

	tree := huffman.NewHuffmanTree("abbccc")

	var result string
	err := tree.Encode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) //000101111
}
