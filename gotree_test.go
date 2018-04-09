package gotree_test

import (
	"testing"

	gotree "github.com/DiSiqueira/GoTree"
)

func TestLast(t *testing.T) {
	want := `Pantera
├── Far Beyond Driven
│   └── 5 minutes Alone
└── Power Metal
`

	artist := gotree.New("Pantera")
	album := artist.Add("Far Beyond Driven")
	album.Add("5 minutes Alone")
	artist.Add("Power Metal")

	got := artist.Print()
	if got != want {
		t.Errorf("Expected: \n%s\n", want)
		t.Errorf("Got: \n%s\n", got)
	}
}
