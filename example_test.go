package gotree_test

import (
	"github.com/DiSiqueira/gotree"
)

func ExamplePrintTree() {

	obj := gotree.ReadFolder("/Users/disiqueira/Documents/Arduino/")
	gotree.PrintTree(obj)

}