package gotree_test

import (
	"testing"

	gotree "github.com/DiSiqueira/GoTree"
)

func TestExampleTree(t *testing.T) {
	artist := gotree.New("Pantera")
	album := artist.Add("Far Beyond Driven\nsee https://en.wikipedia.org/wiki/Pantera\n(1994)")
	five := album.Add("5 minutes Alone")
	five.Add("song by American\ngroove metal")
	album.Add("I’m Broken")
	album.Add("Good Friends and a Bottle of Pills")

	artist.Add("Power Metal\n(1988)")
	artist.Add("Cowboys from Hell\n(1990)")

	output := `Pantera
├── Far Beyond Driven
│   see https://en.wikipedia.org/wiki/Pantera
│   (1994)
│   ├── 5 minutes Alone
│   │   └── song by American
│   │       groove metal
│   ├── I’m Broken
│   └── Good Friends and a Bottle of Pills
├── Power Metal
│   (1988)
└── Cowboys from Hell
    (1990)
`

	if tree := artist.Print(); tree != output {
		t.Errorf("artist.Print() =\n%v, want\n%v", tree, output)
	}
}
