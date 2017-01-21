package gotree_test

import (
	"github.com/disiqueira/gotree"
)

func ExamplePrintTree() {

	obj := gotree.ReadFolder("/Users/disiqueira/Documents/placa_display_led")
	gotree.PrintTree(obj)
}

func ExamplePrintCustom() {

	var artist gotree.GTStructure
	artist.Name = "Pantera"

	var album gotree.GTStructure
	album.Name = "Far Beyond Driven"

	var music gotree.GTStructure
	music.Name = "5 Minutes Alone"

	//Add Music to the Album
	album.Items = append(album.Items, music)

	//Add Album to the Artist
	artist.Items = append(artist.Items, album)

	gotree.PrintTree(artist)
}
