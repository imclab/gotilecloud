package tile

import (
	"testing"
	"tilecloud/geom"
)

func TestQuadGridChildren1(t *testing.T) {
	qg := QuadGrid{MaxZ: 8}
	expected := []Coord{
		Coord{2, 4, 6, 1},
		Coord{2, 4, 7, 1},
		Coord{2, 5, 6, 1},
		Coord{2, 5, 7, 1},
	}
	c := Coord{1, 2, 3, 1}
	ch := qg.Children(c)
	for _, out := range expected {
		if x := <-ch; x != out {
			t.Errorf("<-%v.Children(&%v) = %v, want %v", qg, c, x, out)
		}
	}

}

func TestQuadGridChildren2(t *testing.T) {
	qg := QuadGrid{}
	c := Coord{0, 0, 0, 1}
	ch := qg.Children(c)
	if x, ok := <-ch; ok {
		t.Errorf("<-%v.Children(&%v) = %v, %v, want %v, false", qg, c, x, ok, nil)
	}
}

func TestQuadGridCoord(t *testing.T) {
	qg := QuadGrid{extent: geom.Rect(0, 1, 2, 3), MaxZ: 4}
	for z := 0; z < 4; z++ {
		n := 1 << uint(z)
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				out := Coord{z, x, y, 1}
				r, ok1 := qg.Extent(out)
				if !ok1 {
					t.Errorf("%v.Extent(%v) = %v, %v, want _, true", qg, out, r, ok1)
				}
				in := r.Min
				if x, ok2 := qg.Coord(z, in); !ok2 || x != out {
					t.Errorf("%v.Coord(%v, &%v) = %v, %v, want %v, true", qg, z, in, x, ok2, out)
				}
			}
		}
	}
}

func TestQuadGridExtent(t *testing.T) {
	qg := QuadGrid{extent: geom.Rect(0, 1, 2, 3)}
	expected := map[Coord]geom.Rectangle{
		Coord{0, 0, 0, 1}: geom.Rect(0, 1, 2, 3),
		Coord{1, 0, 0, 1}: geom.Rect(0, 2, 1, 3),
		Coord{1, 0, 1, 1}: geom.Rect(0, 1, 1, 2),
		Coord{1, 1, 0, 1}: geom.Rect(1, 2, 2, 3),
		Coord{1, 1, 1, 1}: geom.Rect(1, 1, 2, 2),
		Coord{2, 0, 0, 1}: geom.Rect(0, 2.5, 0.5, 3),
		Coord{2, 1, 1, 1}: geom.Rect(0.5, 2, 1, 2.5),
		Coord{2, 2, 2, 1}: geom.Rect(1, 1.5, 1.5, 2),
		Coord{2, 3, 3, 1}: geom.Rect(1.5, 1, 2, 1.5),
	}
	for in, out := range expected {
		if x, ok := qg.Extent(in); !ok || x != out {
			t.Errorf("%v.Extent(&%v) = %v, %v, want %v, true", qg, in, x, ok, out)
		}
	}
}

func TestQuadGridInterface(t *testing.T) {
	var _ Grid = &QuadGrid{}
}

func TestQuadGridParent1(t *testing.T) {
	qg := QuadGrid{}
	c := Coord{0, 0, 0, 1}
	if x, ok := qg.Parent(c); ok {
		t.Errorf("%v.Parent(&%v) = %v, %v, want _, false", qg, c, x, ok)
	}
}

func TestQuadGridParent2(t *testing.T) {
	qg := QuadGrid{MaxZ: 8}
	expected := map[Coord]Coord{
		Coord{2, 4, 6, 1}: Coord{1, 2, 3, 1},
		Coord{2, 4, 7, 1}: Coord{1, 2, 3, 1},
		Coord{2, 5, 6, 1}: Coord{1, 2, 3, 1},
		Coord{2, 5, 7, 1}: Coord{1, 2, 3, 1},
	}
	for in, out := range expected {
		if x, ok := qg.Parent(in); !ok || x != out {
			t.Errorf("%v.Parent(&%v) = %v, %v, want %v, true", qg, in, x, ok, out)
		}
	}
}

func TestQuadGridTileSize(t *testing.T) {
	tileSize := Size{256, 256}
	qg := QuadGrid{tileSize: tileSize}
	if x := qg.TileSize(); x != tileSize {
		t.Errorf("%v.TileSize() = %v, want %v", qg, x, tileSize)
	}
}
