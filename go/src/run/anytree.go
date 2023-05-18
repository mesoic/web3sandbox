package run

import(
	"fmt"
    "strconv"
    "strings"

   	"github.com/mesoic/web3sandbox/go/src/utils"	
)

type AnyTree struct {
	value int
	child []*AnyTree
}

func (n *AnyTree) AddChild(value int) {
	n.child = append(n.child, NewAnyTree(value))
}

func (n *AnyTree) ShowTree(depth ...int) {
	if (len(depth) == 0) {
		depth = []int{0}
	}
	fmt.Printf("%s %s\n", strings.Repeat(" ", depth[0]), strconv.Itoa(n.value))	
	if len(n.child) != 0 {
		for _, n := range n.child {
			n.ShowTree(depth[0] + 1)
		}
	}
}

func NewAnyTree(value int) *AnyTree {
	return &AnyTree{
		value,
		[]*AnyTree{},
	}
}

func BuildTree(data []utils.NullableInt) *AnyTree {

	var tree *AnyTree

	var ptr []*AnyTree
	var cix int = 0

	/* Initialize and populate tree */
	for i, value := range data {
		if i == 0 {
			tree = NewAnyTree(value.Value().(int))
			ptr = []*AnyTree{tree}
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

	return tree
}

