package compression

import (
	"fmt"
)

var (
	text = "abbccc"
)

//Node Node
type Node struct {
	value     string
	weight    int
	LeftNode  *Node
	RightNode *Node
}

//HuffmanTree HuffmanTree
type HuffmanTree struct {
	Root *Node
}

//Encode Encode
func (h *HuffmanTree) Encode(character string, s *ItemStack, rootNode *Node) string {
	if rootNode.LeftNode == nil && rootNode.RightNode == nil {

		if rootNode.value == character {
			encoded := ""
			for _, i := range s.items {
				encoded += string(i)
			}
			fmt.Println(encoded)
			return encoded
		}
		return ""

	}

	s.Push("0")
	h.Encode(character, s, rootNode.LeftNode)
	s.Pop()

	s.Push("1")
	h.Encode(character, s, rootNode.RightNode)
	s.Pop()

	return ""

}

func (h *HuffmanTree) Decode() {

}

//NewHuffmanTree NewHuffmanTree
func NewHuffmanTree(text string) *HuffmanTree {

	huffmanTable := NewHuffmanTable(text)
	var j *Node
	for len(huffmanTable.Table) > 1 {

		firstSmallestNode := huffmanTable.GetSmallestNode()
		delete(huffmanTable.Table, firstSmallestNode.value)

		secondSmallestNode := huffmanTable.GetSmallestNode()
		delete(huffmanTable.Table, secondSmallestNode.value)

		j = JoinNodes(firstSmallestNode, secondSmallestNode)
		huffmanTable.Table[j.value] = j
	}
	return &HuffmanTree{Root: j}
}
