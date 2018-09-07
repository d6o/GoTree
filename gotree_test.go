package gotree_test

import (
	"fmt"

	gotree "github.com/DiSiqueira/GoTree"
)

func ExampleTree() {
	artist := gotree.New("Pantera")
	album := artist.Add("Far Beyond Driven")
	album.Add("5 minutes Alone")
	artist.Add("Power Metal")
	fmt.Println(artist.Print())

	// Output:
	// Pantera
	// ├── Far Beyond Driven
	// │   └── 5 minutes Alone
	// └── Power Metal
}
