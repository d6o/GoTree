package gotree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*GTStructure Structure to output print */
type GTStructure struct {
	Name  string
	Items []GTStructure
}

/*PrintTree - Print the tree in console */
func PrintTree(object GTStructure) {

	fmt.Println(object.Name)

	var spaces []bool

	readObjItems(object.Items, spaces)
}

/*ReadFolder - Read a folder and return the generated object */
func ReadFolder(directory string) GTStructure {

	var parent GTStructure

	parent.Name = directory
	parent.Items = createGTReadFolder(directory)

	return parent
}

func createGTReadFolder(directory string) []GTStructure {

	var items []GTStructure
	files, _ := ioutil.ReadDir(directory)

	for _, f := range files {

		var child GTStructure
		child.Name = f.Name()

		if f.IsDir() {
			newDirectory := filepath.Join(directory, f.Name())
			child.Items = createGTReadFolder(newDirectory)
		}

		items = append(items, child)
	}
	return items
}

func printLine(name string, spaces []bool, last bool) {

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

	fmt.Println(indicator + name)

}

func readObjItems(items []GTStructure, spaces []bool) {

	for i, f := range items {

		last := (i >= len(items)-1)

		printLine(f.Name, spaces, last)
		if len(f.Items) > 0 {

			spacesChild := append(spaces, last)

			readObjItems(f.Items, spacesChild)
		}
	}
}
