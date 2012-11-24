package main

import (
	"testing"
	"tilecloud"
)

func TestTileSizeHeight(t *testing.T) {
	tileSize := tilecloud.TileSize{1, 2}
	const out = 2
	if x := tileSize.Height; x != out {
		t.Errorf("%v.Height = %v, want %v", tileSize, x, out)
	}
}

func TestTileSizeWidth(t *testing.T) {
	tileSize := tilecloud.TileSize{1, 2}
	const out = 1
	if x := tileSize.Width; x != out {
		t.Errorf("%v.Width = %v, want %v", tileSize, x, out)
	}
}
