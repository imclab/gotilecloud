package tile

import (
	"testing"
)

func TestCoordEach1(t *testing.T) {
	in := Coord{1, 2, 3, 1}
	out := in
	ch := in.Each()
	x, ok := <-ch
	if !ok {
		t.Errorf("%v.Each() returned closed channel", in)
	}
	if x != out {
		t.Errorf("%v.Each() = %v, want %v", in, x, out)
	}
	_, ok = <-ch
	if ok {
		t.Errorf("%v.Each() returned extra results", in)
	}
}

func TestCoordEach2(t *testing.T) {
	in := Coord{1, 2, 3, 2}
	expected := []Coord{
		Coord{1, 2, 3, 1},
		Coord{1, 2, 4, 1},
		Coord{1, 3, 3, 1},
		Coord{1, 3, 4, 1},
	}
	ch := in.Each()
	for _, out := range expected {
		x := <-ch
		if x != out {
			t.Errorf("%v.Each() = %v, want %v", in, x, out)
		}
	}
	_, ok := <-ch
	if ok {
		t.Errorf("%v.Each() returned extra results", in)
	}
}

func TestCoordString1(t *testing.T) {
	in := Coord{1, 2, 3, 1}
	out := "1/2/3"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}

func TestCoordString2(t *testing.T) {
	in := Coord{1, 2, 3, 4}
	out := "1/2/3:+4/+4"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}
