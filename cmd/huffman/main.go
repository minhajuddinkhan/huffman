package main

import (
	"fmt"

	"github.com/minhajuddinkhan/compression"
)

func main() {

	tree := compression.NewHuffmanTree("abbccc")
	//	stack := compression.ItemStack{}

	var result string
	tree.Encode(&result)

	fmt.Println(result)
}
