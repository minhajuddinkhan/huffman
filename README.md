# Huffman Algorithm

Huffman coding is a lossless data compression algorithm. The idea is to assign variable-length codes to input characters, lengths of the assigned codes are based on the frequencies of corresponding characters.

### Install

```$ go get github.com/minhajuddinkhan/huffman ```

### Usage

```go

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

	fmt.Println(encoded) //00001000111111010110010011110111110010110000100110111000101

	decoded, err := tree.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded) //Are you a gopher?
	

}


```
