package main

import (
	"testing"
	"tilecloud"
)

func TestTileCoordEach1(t *testing.T) {
	in := tilecloud.TileCoord{1, 2, 3, 1}
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

func TestTileCoordEach2(t *testing.T) {
	in := tilecloud.TileCoord{1, 2, 3, 2}
	expected := []tilecloud.TileCoord{
		tilecloud.TileCoord{1, 2, 3, 1},
		tilecloud.TileCoord{1, 2, 4, 1},
		tilecloud.TileCoord{1, 3, 3, 1},
		tilecloud.TileCoord{1, 3, 4, 1},
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

func TestTileCoordString1(t *testing.T) {
	in := tilecloud.TileCoord{1, 2, 3, 1}
	out := "1/2/3"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}

func TestTileCoordString2(t *testing.T) {
	in := tilecloud.TileCoord{1, 2, 3, 4}
	out := "1/2/3:+4/+4"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}
