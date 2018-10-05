package huffman

import (
	"math"
)

//HuffmanFrequenceTable HuffmanFrequenceTable
type HuffmanFrequenceTable struct {
	Table map[string]*Node
}

//NewHuffmanTable NewHuffmanTable
func NewHuffmanTable(text string) *HuffmanFrequenceTable {

	table := make(map[string]*Node)
	for i := 0; i < len(text); i++ {
		c := (string(text[i]))
		if table[c] == nil {
			table[c] = &Node{value: c, weight: 1}
		} else {
			table[c].weight++
		}
	}
	return &HuffmanFrequenceTable{Table: table}

}

//GetSmallestNode GetSmallestNode
func (hft *HuffmanFrequenceTable) GetSmallestNode() *Node {
	smallest := &Node{weight: math.MaxInt32}
	for _, v := range hft.Table {
		if v.weight < smallest.weight {
			smallest = v
		}
	}
	return smallest
}

//JoinNodes JoinNodes
func JoinNodes(firstNode, secondNode *Node, label string) *Node {

	return &Node{
		LeftNode:  firstNode,
		RightNode: secondNode,
		weight:    firstNode.weight + secondNode.weight,
		value:     label,
	}

}
