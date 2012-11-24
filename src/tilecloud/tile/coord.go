package tile

import (
	"fmt"
)

type Coord struct {
	Z, X, Y, N int
}

func (c *Coord) Each() <-chan Coord {
	ch := make(chan Coord)
	go func(z, x, y, n int) {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ch <- Coord{z, x + i, y + j, 1}
			}
		}
		close(ch)
	}(c.Z, c.X, c.Y, c.N)
	return ch
}

func (c *Coord) Hash() int {
	return ((c.X / c.N) << uint(c.Z)) ^ (c.Y / c.N)
}

func (c *Coord) MetaCoord(n int) Coord {
	return Coord{
		c.Z,
		c.N * (c.X / c.N),
		c.N * (c.Y / c.N),
		c.N,
	}
}

func (c *Coord) String() string {
	if c.N == 1 {
		return fmt.Sprintf("%d/%d/%d", c.Z, c.X, c.Y)
	}
	return fmt.Sprintf("%d/%d/%d:+%d/+%d", c.Z, c.X, c.Y, c.N, c.N)
}

/*
type TileGrid interface {
	maxExtent GeoExtent
	tileSize Size
	flipY bool
	Children(tileCoord Coord) chan Coord
	Extent(tileCoord Coord, border int) GeoExtent
	Parent(tileCoord Coord) Coord
	Roots() chan Coord
	Coord(z int, geoPoint GeoPoint) Coord
	Zs() chan int
}

type TileStore interface {
	Contains(tile Tile) bool
	Count() int
	Delete(ch chan Tile) chan Tile
	DeleteOne(tile Tile) Tile
	Get() chan Tile
	GetAll() chan Tile
	GetOne(tile Tile) Tile
	List() chan Tile
	PutOne(tile Tile) Tile
	Put(ch chan Tile) chan Tile
}
*/
