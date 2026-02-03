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

// New returns a new GoTree.Tree.
// The text parameter accepts any string, including empty strings.
// Multi-line text is supported using newline characters.
func New(text string) Tree {
	return &tree{
		text:  text,
		items: make([]Tree, 0),
	}
}

// Add adds a node to the tree with the given text.
// Returns the newly created child node to allow method chaining.
// Empty strings are valid and will create a node with empty text.
func (t *tree) Add(text string) Tree {
	n := New(text)
	t.items = append(t.items, n)
	return n
}

// AddTree adds a tree as an item.
// If the provided tree is nil, this method does nothing (safe no-op).
func (t *tree) AddTree(tree Tree) {
	if tree == nil {
		return
	}
	t.items = append(t.items, tree)
}

// Text returns the node's value
func (t *tree) Text() string {
	return t.text
}

// Items returns all items in the tree
func (t *tree) Items() []Tree {
	return t.items
}

// Print returns an visual representation of the tree
func (t *tree) Print() string {
	return newPrinter().Print(t)
}

func newPrinter() Printer {
	return &printer{}
}

// Print prints a tree to a string
func (p *printer) Print(t Tree) string {
	var builder strings.Builder
	builder.Grow(len(t.Text()) + 100) // Root text + reasonable default

	builder.WriteString(t.Text())
	builder.WriteString(newLine)
	builder.WriteString(p.printItems(t.Items(), []bool{}))

	return builder.String()
}

func (p *printer) printText(text string, spaces []bool, last bool) string {
	var builder strings.Builder

	// Pre-allocate capacity for better performance
	// Estimate: spaces * 4 chars + indicator + text + newline
	builder.Grow(len(spaces)*4 + 4 + len(text) + 1)

	// Build the prefix from spaces
	for _, space := range spaces {
		if space {
			builder.WriteString(emptySpace)
		} else {
			builder.WriteString(continueItem)
		}
	}
	prefix := builder.String()

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	// Reset builder for output
	builder.Reset()
	builder.Grow(len(prefix)*2 + len(text) + 10)

	lines := strings.Split(text, "\n")
	for i := range lines {
		lineText := lines[i] // Fix variable shadowing
		if i == 0 {
			builder.WriteString(prefix)
			builder.WriteString(indicator)
			builder.WriteString(lineText)
			builder.WriteString(newLine)
			continue
		}
		if last {
			indicator = emptySpace
		} else {
			indicator = continueItem
		}
		builder.WriteString(prefix)
		builder.WriteString(indicator)
		builder.WriteString(lineText)
		builder.WriteString(newLine)
	}

	return builder.String()
}

func (p *printer) printItems(t []Tree, spaces []bool) string {
	if len(t) == 0 {
		return ""
	}

	var builder strings.Builder

	// Estimate capacity: rough approximation based on tree size
	// Each item typically needs ~50 chars (conservative estimate)
	builder.Grow(len(t) * 50)

	for i, f := range t {
		last := i == len(t)-1
		builder.WriteString(p.printText(f.Text(), spaces, last))
		if len(f.Items()) > 0 {
			spacesChild := append(spaces, last)
			builder.WriteString(p.printItems(f.Items(), spacesChild))
		}
	}
	return builder.String()
}
