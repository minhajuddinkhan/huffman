package main

import (
	"fmt"

	"github.com/minhajuddinkhan/huffman"
)

func main() {

	tree := huffman.NewHuffmanTree("Are you a gopher?")

	var encoded string
	err := tree.Encode(&encoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(encoded) //11111111100011110010011101000011111011011110010100100011010110010010110110000

	decoded, err := tree.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded) //Whats your name bro?

}
