package main

import (
	"fmt"

	"github.com/minhajuddinkhan/compression"
)

func main() {

	tree := compression.NewHuffmanTree("abbccc")

	var result string
	err := tree.Encode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) //000101111
}
