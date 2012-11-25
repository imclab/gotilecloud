package geom

import (
	"testing"
)

func TestRectangleContains(t *testing.T) {
	r := Rectangle{Point{1, 2}, Point{3, 4}}
	expected := map[Point]bool{
		Point{0, 1}: false,
		Point{0, 2}: false,
		Point{0, 3}: false,
		Point{0, 4}: false,
		Point{0, 5}: false,
		Point{1, 1}: false,
		Point{1, 2}: true,
		Point{1, 3}: true,
		Point{1, 4}: true,
		Point{1, 5}: false,
		Point{2, 1}: false,
		Point{2, 2}: true,
		Point{2, 3}: true,
		Point{2, 4}: true,
		Point{2, 5}: false,
		Point{3, 1}: false,
		Point{3, 2}: true,
		Point{3, 3}: true,
		Point{3, 4}: true,
		Point{3, 5}: false,
		Point{4, 1}: false,
		Point{4, 2}: false,
		Point{4, 3}: false,
		Point{4, 4}: false,
		Point{4, 5}: false,
	}
	for in, out := range expected {
		if x := r.Contains(in); x != out {
			t.Errorf("%v.Contains(%v) = %v, want %v", r, in, x, out)
		}
	}
}

func TestRectangleHeight(t *testing.T) {
	r := Rectangle{Point{1, 2}, Point{3, 4}}
	out := float64(2)
	if x := r.Height(); x != out {
		t.Errorf("%v.Height() = %v, want %v", r, x, out)
	}
}

func TestRectangleWidth(t *testing.T) {
	r := Rectangle{Point{1, 2}, Point{3, 4}}
	out := float64(2)
	if x := r.Width(); x != out {
		t.Errorf("%v.Width() = %v, want %v", r, x, out)
	}
}
