package main

import (
	"fmt"

	"github.com/minhajuddinkhan/compression"
)

func main() {

	tree := compression.NewHuffmanTree("abbccc")
	stack := compression.ItemStack{}
	encoded, err := tree.Encode("a", stack.New(), tree.Root)
	if err != nil {
		panic(err)
	}
	fmt.Println("here", encoded)
}
