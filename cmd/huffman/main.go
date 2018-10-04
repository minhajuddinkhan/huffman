package main

import (
	"fmt"

	"github.com/minhajuddinkhan/compression"
)

func main() {

	tree := compression.NewHuffmanTree("abbccc")
	stack := compression.ItemStack{}
	encoded := tree.Encode("a", stack.New(), tree.Root)
	fmt.Printf(encoded)
}
