package tile

import (
	"testing"
)

func TestOSMLayoutCoord(t *testing.T) {
	ol := OSMLayout{}
	in := "1/2/3z"
	out := Coord{1, 2, 3, 1}
	if x, ok := ol.Coord(in); !ok || x != out {
		t.Errorf("%v.Coord(%v) = %v, %v, want %v, true", ol, in, x, ok, out)
	}
}

func TestOSMLayoutInterface(t *testing.T) {
	var _ Layout = &OSMLayout{}
}

func TestOSMLayoutFilename(t *testing.T) {
	ol := OSMLayout{}
	in := Coord{1, 2, 3, 1}
	const out = "1/2/3"
	if x := ol.Filename(in); x != out {
		t.Errorf("%v.Filename(%v) = %v, want %v", ol, in, x, out)
	}
}
