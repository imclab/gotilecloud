package tile

import (
	"testing"
)

func TestBoundsEach(t *testing.T) {
	b := Bounds{1, 2, 3, 4}
	expected := map[Coord]bool{
		Coord{1, 0, 1, 1}: false,
		Coord{1, 0, 2, 1}: false,
		Coord{1, 0, 3, 1}: false,
		Coord{1, 0, 4, 1}: false,
		Coord{1, 0, 5, 1}: false,
		Coord{1, 1, 1, 1}: false,
		Coord{1, 1, 2, 1}: true,
		Coord{1, 1, 3, 1}: true,
		Coord{1, 1, 4, 1}: false,
		Coord{1, 1, 5, 1}: false,
		Coord{1, 2, 1, 1}: false,
		Coord{1, 2, 2, 1}: true,
		Coord{1, 2, 3, 1}: true,
		Coord{1, 2, 4, 1}: false,
		Coord{1, 2, 5, 1}: false,
		Coord{1, 3, 1, 1}: false,
		Coord{1, 3, 2, 1}: false,
		Coord{1, 3, 3, 1}: false,
		Coord{1, 3, 4, 1}: false,
		Coord{1, 3, 5, 1}: false,
	}
	for in, out := range expected {
		if x := b.Contains(&in); x != out {
			t.Errorf("%v.Contains(%v) = %v, want %v", b, in, x, out)
		}
	}

}

func TestBoundsHeight(t *testing.T) {
	b := Bounds{1, 2, 3, 4}
	const out = 2
	if x := b.Height(); x != out {
		t.Errorf("%v.Height() = %v, want %v", b, x, out)
	}
}

func TestBoundsWidth(t *testing.T) {
	b := Bounds{1, 2, 3, 4}
	const out = 2
	if x := b.Width(); x != out {
		t.Errorf("%v.Width() = %v, want %v", b, x, out)
	}
}
