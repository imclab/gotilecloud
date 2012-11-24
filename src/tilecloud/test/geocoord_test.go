package main

import (
	"testing"
	"tilecloud"
)

func TestGeoCoordString(t *testing.T) {
	in := tilecloud.GeoCoord{1, 2}
	out := "(1, 2)"
	if x := in.String(); x != out {
		t.Errorf("%v.String() = %v, want %v", in, x, out)
	}
}
