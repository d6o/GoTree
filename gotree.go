// Package gotree create and print tree.
package gotree

import (
	"strings"
)

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

type (
	tree struct {
		text  string
		items []Tree
	}

	// Tree is tree interface
	Tree interface {
		Add(text string) Tree
		AddTree(tree Tree)
		Items() []Tree
		Text() string
		Print() string
	}

	printer struct {
	}

	// Printer is printer interface
	Printer interface {
		Print(Tree) string
	}
)

// New returns a new GoTree.Tree
func New(text string) Tree {
	return &tree{
		text:  text,
		items: []Tree{},
	}
}

// Add node in tree
func (t *tree) Add(text string) Tree {
	n := New(text)
	t.items = append(t.items, n)
	return n
}

// AddTree is add tree in present tree
func (t *tree) AddTree(tree Tree) {
	t.items = append(t.items, tree)
}

// Text return text of root name
func (t *tree) Text() string {
	return t.text
}

// Items return slice of tree nodes
func (t *tree) Items() []Tree {
	return t.items
}

// Print return string of tree
func (t *tree) Print() string {
	return newPrinter().Print(t)
}

func newPrinter() Printer {
	return &printer{}
}

func (p *printer) Print(t Tree) string {
	return t.Text() + newLine + p.printItems(t.Items(), []bool{})
}

func (p *printer) printText(text string, spaces []bool, last bool) string {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continueItem
		}
	}

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	var out string
	lines := strings.Split(text, "\n")
	for i := range lines {
		text := lines[i]
		if i == 0 {
			out += result + indicator + text + newLine
			continue
		}
		if last {
			indicator = emptySpace
		} else {
			indicator = continueItem
		}
		out += result + indicator + text + newLine
	}

	return out
}

func (p *printer) printItems(t []Tree, spaces []bool) string {
	var result string
	for i, f := range t {
		last := i == len(t)-1
		result += p.printText(f.Text(), spaces, last)
		if len(f.Items()) > 0 {
			spacesChild := append(spaces, last)
			result += p.printItems(f.Items(), spacesChild)
		}
	}
	return result
}
