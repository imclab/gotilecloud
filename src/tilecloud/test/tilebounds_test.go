package main

import (
	"testing"
	"tilecloud"
)

func TestTileBoundsEach(t *testing.T) {
	tileBounds := tilecloud.TileBounds{1, 2, 3, 4}
	expected := map[tilecloud.TileCoord]bool{
		tilecloud.TileCoord{1, 0, 1, 1}: false,
		tilecloud.TileCoord{1, 0, 2, 1}: false,
		tilecloud.TileCoord{1, 0, 3, 1}: false,
		tilecloud.TileCoord{1, 0, 4, 1}: false,
		tilecloud.TileCoord{1, 0, 5, 1}: false,
		tilecloud.TileCoord{1, 1, 1, 1}: false,
		tilecloud.TileCoord{1, 1, 2, 1}: true,
		tilecloud.TileCoord{1, 1, 3, 1}: true,
		tilecloud.TileCoord{1, 1, 4, 1}: false,
		tilecloud.TileCoord{1, 1, 5, 1}: false,
		tilecloud.TileCoord{1, 2, 1, 1}: false,
		tilecloud.TileCoord{1, 2, 2, 1}: true,
		tilecloud.TileCoord{1, 2, 3, 1}: true,
		tilecloud.TileCoord{1, 2, 4, 1}: false,
		tilecloud.TileCoord{1, 2, 5, 1}: false,
		tilecloud.TileCoord{1, 3, 1, 1}: false,
		tilecloud.TileCoord{1, 3, 2, 1}: false,
		tilecloud.TileCoord{1, 3, 3, 1}: false,
		tilecloud.TileCoord{1, 3, 4, 1}: false,
		tilecloud.TileCoord{1, 3, 5, 1}: false,
	}
	for in, out := range expected {
		if x := tileBounds.Contains(&in); x != out {
			t.Errorf("%v.Contains(%v) = %v, want %v", tileBounds, in, x, out)
		}
	}

}

func TestTileBoundsHeight(t *testing.T) {
	tileBounds := tilecloud.TileBounds{1, 2, 3, 4}
	const out = 2
	if x := tileBounds.Height(); x != out {
		t.Errorf("%v.Height() = %v, want %v", tileBounds, x, out)
	}
}

func TestTileBoundsString(t *testing.T) {
	in := tilecloud.TileBounds{1, 2, 3, 4}
	out := "(1, 2, 3, 4)"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}

func TestTileBoundsWidth(t *testing.T) {
	tileBounds := tilecloud.TileBounds{1, 2, 3, 4}
	const out = 2
	if x := tileBounds.Width(); x != out {
		t.Errorf("%v.Width() = %v, want %v", tileBounds, x, out)
	}
}
