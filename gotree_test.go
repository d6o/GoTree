package gotree

import (
	"fmt"
	"reflect"
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
			if got := New(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tree.Add() = %v, want %v", got, tt.want)
			}
			if tt.parentCount != len(tree.Items()) {
				t.Errorf("tree total items = %v, want %v", got, tt.want)
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
			if got := tree.Items(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tree.Items() = %v, want %v", got, tt.want)
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
