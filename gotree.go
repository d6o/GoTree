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
	Tree struct {
		text  string
		items []Node
	}

	// Node is tree inteface
	Node interface {
		Items() []Node
		Text() string
	}
)

// New returns a new GoTree.Tree
func New(text string) *Tree {
	return &Tree{
		text:  text,
		items: []Node{},
	}
}

// Add node in tree
func (t *Tree) Add(text string) *Tree {
	n := New(text)
	t.items = append(t.items, n)
	return n
}

// AddNode is add tree in present tree
func (t *Tree) AddNode(n Node) {
	t.items = append(t.items, n)
}

// Text return text of root name
func (t *Tree) Text() string {
	return t.text
}

// Items return slice of tree nodes
func (t *Tree) Items() []Node {
	return t.items
}

// Print return string of tree
func (t *Tree) Print() string {
	return Print(t)
}

func Print(n Node) string {
	var out strings.Builder
	printNode(&out, n)
	return out.String()
}

func printNode(out *strings.Builder, n Node) {
	out.WriteString(n.Text())
	out.WriteString(newLine)
	printItems(out, n.Items(), []bool{})
}

func printText(out *strings.Builder, text string, spaces []bool, last bool) {
	var prefix string
	for _, space := range spaces {
		if space {
			prefix += emptySpace
		} else {
			prefix += continueItem
		}
	}

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	lines := strings.Split(text, "\n")
	for i := range lines {
		text := lines[i]
		if i == 0 {
			out.WriteString(prefix + indicator + text + newLine)
			continue
		}
		if last {
			indicator = emptySpace
		} else {
			indicator = continueItem
		}
		out.WriteString(prefix + indicator + text + newLine)
	}
}

func printItems(out *strings.Builder, t []Node, spaces []bool) string {
	var result string
	for i, f := range t {
		last := i == len(t)-1
		printText(out, f.Text(), spaces, last)
		if items := f.Items(); len(items) > 0 {
			spacesChild := append(spaces, last)
			printItems(out, f.Items(), spacesChild)
		}
	}
	return result
}
