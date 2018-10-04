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
	Text string
	Root *Node
}

//Encode Encode
func (h *HuffmanTree) Encode(ch string, s *ItemStack, rootNode *Node) (string, error) {

	encoded := ""

	if rootNode.LeftNode == nil && rootNode.RightNode == nil {
		if rootNode.value == ch {
			for _, i := range s.items {
				encoded += string(i)
			}
			fmt.Println(encoded)
			return fmt.Sprintf("%s", encoded), nil
		}
		return "", nil
	}
	s.Push("0")
	h.Encode(ch, s, rootNode.LeftNode)
	s.Pop()

	s.Push("1")
	h.Encode(ch, s, rootNode.RightNode)
	s.Pop()

	return encoded, nil

}

//Decode Decode
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
	return &HuffmanTree{Root: j, Text: text}
}
