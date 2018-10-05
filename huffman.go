package huffman

import "errors"

//Node Node
type Node struct {
	value     string
	weight    int
	LeftNode  *Node
	RightNode *Node
}

//Tree Huffman Tree Structure
type Tree struct {
	Text string
	Root *Node
}

func (t *Tree) encodeCharacter(ch string, s *ItemStack, rootNode *Node, result *string) {

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
	t.encodeCharacter(ch, s, rootNode.LeftNode, result)
	s.Pop()

	s.Push("1")
	t.encodeCharacter(ch, s, rootNode.RightNode, result)
	s.Pop()
}

//Encode Encode
func (t *Tree) Encode(result *string) error {

	if t.Root == nil {
		return errors.New("root cannot be null")
	}
	for _, c := range t.Text {
		stack := ItemStack{}
		t.encodeCharacter(string(c), stack.New(), t.Root, result)
	}
	return nil

}

//Decode Decode
func (t *Tree) Decode(encoded string) (string, error) {
	if t.Root == nil {
		return "", errors.New("root cannot be null")
	}

	n := t.Root
	decoded := ""
	for i := 0; i < len(encoded); i++ {
		ch := string(encoded[i])
		if ch == "0" {
			n = n.LeftNode
		}
		if ch == "1" {
			n = n.RightNode
		}
		if n.LeftNode == nil && n.RightNode == nil {
			decoded += n.value
			n = t.Root

		}
	}

	return decoded, nil

}

//NewHuffmanTree NewHuffmanTree
func NewHuffmanTree(text string) *Tree {

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
	return &Tree{Root: j, Text: text}
}
