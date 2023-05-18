package run

import (
	"github.com/mesoic/web3sandbox/go/src/utils"
)

func Run() {

	data := []utils.NullableInt{
		utils.NewInt(1),
		utils.NewNull(),
		utils.NewInt(2),
		utils.NewInt(3),
		utils.NewInt(4),
		utils.NewInt(5),
		utils.NewNull(),
		utils.NewInt(6),
		utils.NewInt(7),
		utils.NewInt(8),
		utils.NewNull(),
		utils.NewNull(),
		utils.NewInt(9),
		utils.NewInt(0),
	}

	tree := BuildTree(data)
	tree.ShowTree()
}
