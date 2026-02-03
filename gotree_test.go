package gotree

import (
	"fmt"
	"testing"
)

func ExampleTree() {
	artist := New("Pantera")
	album := artist.Add("Far Beyond Driven\nsee https://en.wikipedia.org/wiki/Pantera\n(1994)")
	five := album.Add("5 minutes Alone")
	five.Add("song by American\ngroove metal")
	album.Add("I’m Broken")
	album.Add("Good Friends and a Bottle of Pills")

	artist.Add("Power Metal\n(1988)")
	artist.Add("Cowboys from Hell\n(1990)")
	fmt.Println(artist.Print())

	// Output:
	// Pantera
	// ├── Far Beyond Driven
	// │   see https://en.wikipedia.org/wiki/Pantera
	// │   (1994)
	// │   ├── 5 minutes Alone
	// │   │   └── song by American
	// │   │       groove metal
	// │   ├── I’m Broken
	// │   └── Good Friends and a Bottle of Pills
	// ├── Power Metal
	// │   (1988)
	// └── Cowboys from Hell
	//     (1990)
}

func TestNew(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			name: "Create new Tree",
			args: args{
				text: "new tree",
			},
			want: &tree{
				text:  "new tree",
				items: []Tree{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.text)
			if got.Text() != tt.want.Text() {
				t.Errorf("New() text = %v, want %v", got.Text(), tt.want.Text())
			}
			if len(got.Items()) != len(tt.want.Items()) {
				t.Errorf("New() items length = %v, want %v", len(got.Items()), len(tt.want.Items()))
			}
		})
	}
}

func Test_tree_Add(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	type args struct {
		text string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        Tree
		parentCount int
	}{
		{
			name: "Adding a new item into an empty tree",
			args: args{
				text: "child item",
			},
			fields: fields{
				items: []Tree{},
			},
			want: &tree{
				text:  "child item",
				items: []Tree{},
			},
			parentCount: 1,
		},
		{
			name: "Adding a new item into a full tree",
			args: args{
				text: "fourth item",
			},
			fields: fields{
				items: []Tree{
					New("test"),
					New("test2"),
					New("test3"),
				},
			},
			want: &tree{
				text:  "fourth item",
				items: []Tree{},
			},
			parentCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				text:  tt.fields.text,
				items: tt.fields.items,
			}
			got := tree.Add(tt.args.text)
			if got.Text() != tt.want.Text() {
				t.Errorf("tree.Add() text = %v, want %v", got.Text(), tt.want.Text())
			}
			if len(got.Items()) != len(tt.want.Items()) {
				t.Errorf("tree.Add() items length = %v, want %v", len(got.Items()), len(tt.want.Items()))
			}
			if tt.parentCount != len(tree.Items()) {
				t.Errorf("tree total items = %v, want %v", len(tree.Items()), tt.parentCount)
			}
		})
	}
}

func Test_tree_AddTree(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	type args struct {
		tree Tree
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		itemCount int
	}{
		{
			name: "Adding a new item into an empty tree",
			args: args{
				tree: New("child item"),
			},
			fields: fields{
				items: []Tree{},
			},
			itemCount: 1,
		},
		{
			name: "Adding a new item into a full tree",
			args: args{
				tree: New("fourth item"),
			},
			fields: fields{
				items: []Tree{
					New("test"),
					New("test2"),
					New("test3"),
				},
			},
			itemCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				text:  tt.fields.text,
				items: tt.fields.items,
			}
			tree.AddTree(tt.args.tree)
		})
	}
}

func Test_tree_Text(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return the correct value",
			fields: fields{
				text: "item",
			},
			want: "item",
		},
		{
			name: "Return the correct value while empty",
			fields: fields{
				text: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				text:  tt.fields.text,
				items: tt.fields.items,
			}
			if got := tree.Text(); got != tt.want {
				t.Errorf("tree.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Items(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   []Tree
	}{
		{
			name: "Return empty if there is no items under the tree",
			fields: fields{
				text:  "top level item",
				items: []Tree{},
			},
			want: []Tree{},
		},
		{
			name: "Return all items under the tree",
			fields: fields{
				text: "top level item",
				items: []Tree{
					New("first child"),
					New("second child"),
				},
			},
			want: []Tree{
				New("first child"),
				New("second child"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				text:  tt.fields.text,
				items: tt.fields.items,
			}
			got := tree.Items()
			if len(got) != len(tt.want) {
				t.Fatalf("Items() length = %v, want %v", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].Text() != tt.want[i].Text() {
					t.Errorf("Items()[%d].Text() = %v, want %v", i, got[i].Text(), tt.want[i].Text())
				}
			}
		})
	}
}

func Test_tree_Print(t *testing.T) {
	threeLevelTree := New("First Level")
	threeLevelTree.Add("Second level").Add("Third Level")

	complexTree := New("Daft Punk")
	ram := complexTree.Add("Random Access Memories")
	complexTree.Add("Humam After All")
	alive := complexTree.Add("Alive 2007")

	ram.Add("Give Life Back to Music")
	ram.Add("Giorgio by Moroder")
	ram.Add("Within")

	alive.Add("Touch It/Technologic")
	alive.Add("Face to Face/Too Long")

	type fields struct {
		tree Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Print a single item tree",
			fields: fields{
				tree: New("single item"),
			},
			want: `single item
`,
		},
		{
			name: "Print a three level tree",
			fields: fields{
				tree: threeLevelTree,
			},
			want: `First Level
└── Second level
    └── Third Level
`,
		},
		{
			name: "Print a three level tree",
			fields: fields{
				tree: complexTree,
			},
			want: `Daft Punk
├── Random Access Memories
│   ├── Give Life Back to Music
│   ├── Giorgio by Moroder
│   └── Within
├── Humam After All
└── Alive 2007
    ├── Touch It/Technologic
    └── Face to Face/Too Long
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.tree.Print(); got != tt.want {
				t.Errorf("tree.Print() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

// Benchmark tests

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("test node")
	}
}

func BenchmarkAdd(b *testing.B) {
	tree := New("root")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Add("child")
	}
}

func BenchmarkPrintSmallTree(b *testing.B) {
	tree := New("Root")
	tree.Add("Child 1")
	tree.Add("Child 2")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Print()
	}
}

func BenchmarkPrintDeepTree(b *testing.B) {
	tree := New("Root")
	current := tree
	for i := 0; i < 10; i++ {
		current = current.Add("Level " + string(rune(i+'0')))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Print()
	}
}

func BenchmarkPrintWideTree(b *testing.B) {
	tree := New("Root")
	for i := 0; i < 100; i++ {
		tree.Add("Child " + string(rune(i+'0')))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Print()
	}
}

func BenchmarkPrintComplexTree(b *testing.B) {
	tree := New("Root")
	for i := 0; i < 10; i++ {
		branch := tree.Add("Branch " + string(rune(i+'0')))
		for j := 0; j < 10; j++ {
			leaf := branch.Add("Leaf " + string(rune(i+'0')) + "-" + string(rune(j+'0')))
			if j%2 == 0 {
				leaf.Add("Sub-leaf")
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Print()
	}
}

func BenchmarkPrintMultilineText(b *testing.B) {
	tree := New("Root")
	multilineText := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5"
	for i := 0; i < 5; i++ {
		tree.Add(multilineText)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Print()
	}
}

// Edge case tests

func TestAddTree_Nil(t *testing.T) {
	tree := New("root")
	tree.AddTree(nil) // Should not panic

	if len(tree.Items()) != 0 {
		t.Errorf("AddTree(nil) should not add item, got %d items", len(tree.Items()))
	}
}

func TestPrint_EmptyTree(t *testing.T) {
	tree := New("")
	got := tree.Print()
	want := "\n"

	if got != want {
		t.Errorf("Print() = %#v, want %#v", got, want)
	}
}

func TestPrint_VeryDeepTree(t *testing.T) {
	tree := New("Root")
	current := tree

	// Create a tree 100 levels deep
	for i := 0; i < 100; i++ {
		current = current.Add(fmt.Sprintf("Level %d", i))
	}

	// Should not panic or cause stack overflow
	output := tree.Print()

	// Verify it contains expected depth
	if len(output) == 0 {
		t.Error("Deep tree should produce output")
	}
}

func TestText_PreservesInput(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{"empty", ""},
		{"simple", "simple"},
		{"with newlines", "with\nnewlines"},
		{"with tabs", "with\ttabs"},
		{"with unicode", "with unicode: 🌲🌳🌴"},
		{"with special chars", "with special chars: !@#$%^&*()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New(tt.text)
			if tree.Text() != tt.text {
				t.Errorf("Text() = %v, want %v", tree.Text(), tt.text)
			}
		})
	}
}
