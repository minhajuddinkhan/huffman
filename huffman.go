package compression

import (
	"errors"
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

func (h *HuffmanTree) getEncodedStringForCharacter(ch string, s *ItemStack, rootNode *Node) string {

	if rootNode.LeftNode == nil && rootNode.RightNode == nil {
		if rootNode.value == ch {
			encoded := ""
			for _, i := range s.items {
				encoded += string(i)
			}
			fmt.Print(encoded)
			return encoded
		}
		return ""
	}
	s.Push("0")
	h.getEncodedStringForCharacter(ch, s, rootNode.LeftNode)
	s.Pop()

	s.Push("1")
	h.getEncodedStringForCharacter(ch, s, rootNode.RightNode)
	s.Pop()

	return ""

}

//Encode Encode
func (h *HuffmanTree) Encode() (string, error) {
	if h.Root == nil {
		return "", errors.New("Root Node is null")
	}
	var str string
	for _, c := range h.Text {
		stack := ItemStack{}
		binary := h.getEncodedStringForCharacter(string(c), stack.New(), h.Root)
		str += binary
	}
	return str, nil

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
