package huffman

import "errors"

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

func (h *HuffmanTree) encodeCharacter(ch string, s *ItemStack, rootNode *Node, result *string) {
	encoded := ""

	if rootNode.LeftNode == nil && rootNode.RightNode == nil {
		if rootNode.value == ch {
			for _, i := range s.items {
				encoded += string(i)
			}
			*result += encoded

		}

		return
	}
	s.Push("0")
	h.encodeCharacter(ch, s, rootNode.LeftNode, result)
	s.Pop()

	s.Push("1")
	h.encodeCharacter(ch, s, rootNode.RightNode, result)
	s.Pop()
}

//Encode Encode
func (h *HuffmanTree) Encode(result *string) error {

	if h.Root == nil {
		return errors.New("root cannot be null")
	}
	for _, c := range h.Text {
		stack := ItemStack{}
		h.encodeCharacter(string(c), stack.New(), h.Root, result)
	}
	return nil

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
