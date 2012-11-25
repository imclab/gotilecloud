package tile

import (
	"tilecloud/geom"
)

type Grid interface {
	Children(c Coord) <-chan Coord
	Coord(z int, p geom.Point) (Coord, bool)
	Extent(c Coord) (geom.Rectangle, bool)
	Parent(c Coord) (Coord, bool)
	TileSize() Size
}
