package gotree

import (
	"testing"
)

// TestUpdatingItemsStructure should test whenever item updates in the tree structure are
// reflected correctly in the rendered structure
func TestUpdatingItemsStructure(t *testing.T) {
	expected := "Pantera\n" +
		"└── Far Beyond Driven\n" +
		"    └── 5 minutes Alone\n"

	var artist GTStructure
	artist.Name = "Pantera Typo0"

	var album GTStructure
	album.Name = "Far Beyond Driven Typo1"

	var music GTStructure
	music.Name = "5 Minutes Alone Typo2"

	// Add Music to the Album
	album.Items = append(album.Items, &music)

	// Add Album to the Artist
	artist.Items = append(artist.Items, &album)

	// apply updates to the items that are already in the tree structure
	music.Name = "5 minutes Alone"
	album.Name = "Far Beyond Driven"
	artist.Name = "Pantera"

	actual := StringTree(&artist)

	if actual != expected {
		t.Fatalf("Actual tree::\n[%s]\nis not the same as expected:\n[%s]", actual, expected)
	}
}
