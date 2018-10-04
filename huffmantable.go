package compression

import (
	"math"
	"strings"
)

type HuffmanFrequenceTable struct {
	Table map[string]*Node
}

func NewHuffmanTable(text string) *HuffmanFrequenceTable {

	table := make(map[string]*Node)
	for i := 0; i < len(text); i++ {
		c := strings.ToLower(string(text[i]))
		if table[c] == nil {
			table[c] = &Node{value: c, weight: 1}
		} else {
			table[c].weight++
		}
	}
	return &HuffmanFrequenceTable{Table: table}

}

func (hft *HuffmanFrequenceTable) GetSmallestNode() *Node {
	smallest := &Node{weight: math.MaxInt32}
	for _, v := range hft.Table {
		if v.weight < smallest.weight {
			smallest = v
		}
	}
	return smallest
}

func JoinNodes(firstNode, secondNode *Node) *Node {

	return &Node{
		LeftNode:  firstNode,
		RightNode: secondNode,
		weight:    firstNode.weight + secondNode.weight,
		value:     firstNode.value + secondNode.value,
	}

}
