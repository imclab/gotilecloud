package main

import (
	"testing"
	"tilecloud"
)

func TestGeoBoundsContains(t *testing.T) {
	geoBounds := tilecloud.GeoBounds{1, 2, 3, 4}
	expected := map[tilecloud.GeoCoord]bool{
		tilecloud.GeoCoord{0, 1}: false,
		tilecloud.GeoCoord{0, 2}: false,
		tilecloud.GeoCoord{0, 3}: false,
		tilecloud.GeoCoord{0, 4}: false,
		tilecloud.GeoCoord{0, 5}: false,
		tilecloud.GeoCoord{1, 1}: false,
		tilecloud.GeoCoord{1, 2}: true,
		tilecloud.GeoCoord{1, 3}: true,
		tilecloud.GeoCoord{1, 4}: true,
		tilecloud.GeoCoord{1, 5}: false,
		tilecloud.GeoCoord{2, 1}: false,
		tilecloud.GeoCoord{2, 2}: true,
		tilecloud.GeoCoord{2, 3}: true,
		tilecloud.GeoCoord{2, 4}: true,
		tilecloud.GeoCoord{2, 5}: false,
		tilecloud.GeoCoord{3, 1}: false,
		tilecloud.GeoCoord{3, 2}: true,
		tilecloud.GeoCoord{3, 3}: true,
		tilecloud.GeoCoord{3, 4}: true,
		tilecloud.GeoCoord{3, 5}: false,
		tilecloud.GeoCoord{4, 1}: false,
		tilecloud.GeoCoord{4, 2}: false,
		tilecloud.GeoCoord{4, 3}: false,
		tilecloud.GeoCoord{4, 4}: false,
		tilecloud.GeoCoord{4, 5}: false,
	}
	for in, out := range expected {
		if x := geoBounds.Contains(&in); x != out {
			t.Errorf("%v.Contains(%v) := %v, want %v", geoBounds, in, x, out)
		}
	}
}

func TestGeoBoundsHeight(t *testing.T) {
	in := tilecloud.GeoBounds{1, 2, 3, 4}
	out := float64(2)
	if x := in.Height(); x != out {
		t.Errorf("%v.Height() = %v, want %v", in, x, out)
	}
}

func TestGeoBoundsString(t *testing.T) {
	in := tilecloud.GeoBounds{1, 2, 3, 4}
	out := "(1, 2, 3, 4)"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}

func TestGeoBoundsWidth(t *testing.T) {
	in := tilecloud.GeoBounds{1, 2, 3, 4}
	out := float64(2)
	if x := in.Width(); x != out {
		t.Errorf("%v.Width() = %v, want %v", in, x, out)
	}
}
