package fn

import (
	"fmt"
	"strconv"
	"strings"

	"../msg"
	"../utils"
)

type AnyTreeNode struct {
	value int
	child []*AnyTreeNode
}

func (node *AnyTreeNode) AddChild(value int) {
	node.child = append(node.child, InitializeNode(value))
}

func (n *AnyTreeNode) ShowTree(depth int) {
	fmt.Printf("%s %s\n", strings.Repeat(" ", depth), strconv.Itoa(n.value))
	if len(n.child) != 0 {
		for _, node := range n.child {
			node.ShowTree(depth + 1)
		}
	}
}

func InitializeNode(value int) *AnyTreeNode {
	return &AnyTreeNode{
		value,
		[]*AnyTreeNode{},
	}
}

type AnyTreeRunner struct {
	data []utils.NilableInt
	tree *AnyTreeNode
}

func (n *AnyTreeRunner) ShowData() {

	m := msg.Msg{}
	m.Message("[")
	for _, value := range n.data {
		if value.Value() == nil {
			m.Message(fmt.Sprintf("  %s,", "nil"))
		} else {
			m.Message(fmt.Sprintf("  %s,", strconv.Itoa(value.Value().(int))))
		}
	}
	m.Message("]")
}

func (n *AnyTreeRunner) BuildTree() {

	var ptr []*AnyTreeNode
	var cix int = 0

	/* Initialize and populate tree */
	for i, value := range n.data {
		if i == 0 {
			n.tree = InitializeNode(value.Value().(int))
			ptr = []*AnyTreeNode{n.tree}
		} else if i == 1 {
			continue
		} else if value.Value() != nil {
			ptr[cix].AddChild(value.Value().(int))
		} else if value.Value() == nil && cix < len(ptr)-1 {
			cix = cix + 1
		} else {
			ptr = ptr[0].child
			cix = 0
		}
	}
}

func Run() {
	var tree = AnyTreeRunner{
		[]utils.NilableInt{
			utils.NewInt(1),
			utils.NewNil(),
			utils.NewInt(2),
			utils.NewInt(3),
			utils.NewInt(4),
			utils.NewInt(5),
			utils.NewNil(),
			utils.NewInt(6),
			utils.NewInt(7),
			utils.NewInt(8),
			utils.NewNil(),
			utils.NewNil(),
			utils.NewInt(9),
			utils.NewInt(0),
		},
		nil,
	}
	tree.ShowData()
	tree.BuildTree()
	tree.tree.ShowTree(0)
}
