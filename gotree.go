package gotree

import (
	"io/ioutil"
	"fmt"
	"path/filepath"
)

type GTStructure struct {
	Name  string
	Items [] GTStructure
}

func PrintTree(object GTStructure) {

	fmt.Println(object.Name)

	var spaces []bool

	ReadObjItems(object.Items, spaces)
}

func ReadFolder(directory string) GTStructure {

	var parent GTStructure

	parent.Name = directory
	parent.Items = CreateGTReadFolder(directory)

	return parent
}

func CreateGTReadFolder(directory string) []GTStructure {

	var items []GTStructure
	files, _ := ioutil.ReadDir(directory)

	for _, f := range files {

		var child GTStructure
		child.Name = f.Name()

		if f.IsDir() {
			newDirectory := filepath.Join(directory, f.Name())
			child.Items = CreateGTReadFolder(newDirectory)
		}

		items = append(items, child)
	}
	return items
}

func PrintLine(name string, spaces []bool, last bool) {

	for _, space := range spaces {
		if space {
			fmt.Print("    ")
		} else {
			fmt.Print("|   ")
		}
	}

	indicator := "├── "

	if last {
		indicator = "└── "
	}

	fmt.Println(indicator+name)

}

func ReadObjItems(items []GTStructure, spaces []bool) {

	for i, f := range items {

		last := (i>=len(items) -1 )

		PrintLine(f.Name, spaces, last)
		if len(f.Items) > 0 {

			spacesChild := append(spaces, last)

			ReadObjItems(f.Items, spacesChild)
		}
	}
}
