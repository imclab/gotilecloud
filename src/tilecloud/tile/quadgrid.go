package tile

import (
	"math"
	"tilecloud/geom"
)

type QuadGrid struct {
	extent     geom.Rectangle
	flipY      bool
	tileSize   Size
	MinZ, MaxZ int
}

func (qg *QuadGrid) Children(c Coord) <-chan Coord {
	ch := make(chan Coord)
	if qg.MinZ <= c.Z && c.Z < qg.MaxZ {
		go func(ch chan Coord, z, x, y int) {
			ch <- Coord{z + 1, 2 * x, 2 * y, 1}
			ch <- Coord{z + 1, 2 * x, 2*y + 1, 1}
			ch <- Coord{z + 1, 2*x + 1, 2 * y, 1}
			ch <- Coord{z + 1, 2*x + 1, 2*y + 1, 1}
			close(ch)
		}(ch, c.Z, c.X, c.Y)
	} else {
		close(ch)
	}
	return ch
}

func (qg *QuadGrid) Coord(z int, p geom.Point) (Coord, bool) {
	if z < qg.MinZ || qg.MaxZ < z {
		return Coord{}, false
	}
	size := qg.extent.Size()
	min := qg.extent.Min
	n := math.Pow(2, float64(z))
	x := int((p.X - min.X) * n / size.Width)
	y := int((p.Y - min.Y) * n / size.Height)
	if !qg.flipY {
		y = 1<<uint(z) - y - 1
	}
	return Coord{z, x, y, 1}, true
}

func (qg *QuadGrid) Extent(c Coord) (geom.Rectangle, bool) {
	size := qg.extent.Size()
	min := qg.extent.Min
	n := math.Pow(2, float64(c.Z))
	y := c.Y
	if !qg.flipY {
		y = 1<<uint(c.Z) - y - c.N
	}
	minX := min.X + size.Width*float64(c.X)/n
	minY := min.Y + size.Height*float64(y)/n
	maxX := min.X + size.Width*float64(c.X+c.N)/n
	maxY := min.Y + size.Height*float64(y+c.N)/n
	return geom.Rect(minX, minY, maxX, maxY), true
}

func (qg *QuadGrid) Parent(c Coord) (Coord, bool) {
	if qg.MinZ < c.Z && c.Z <= qg.MaxZ {
		return Coord{c.Z - 1, c.X / 2, c.Y / 2, 1}, true
	}
	return Coord{}, false
}

func (qg *QuadGrid) TileSize() Size {
	return qg.tileSize
}
