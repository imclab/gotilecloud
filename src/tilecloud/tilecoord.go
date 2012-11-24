package tilecloud

import (
	"fmt"
)

type TileCoord struct {
	Z int
	X int
	Y int
	N int
}

func (p *TileCoord) Each() chan TileCoord {
	ch := make(chan TileCoord)
	go func() {
		for i := 0; i < p.N; i++ {
			for j := 0; j < p.N; j++ {
				ch <- TileCoord{p.Z, p.X + i, p.Y + j, 1}
			}
		}
		close(ch)
	}()
	return ch
}

func (p *TileCoord) Hash() int {
	return ((p.X / p.N) << uint(p.Z)) ^ (p.Y / p.N)
}

func (p *TileCoord) MetaTileCoord(n int) TileCoord {
	return TileCoord{
		p.Z,
		p.N * (p.X / p.N),
		p.N * (p.Y / p.N),
		p.N,
	}
}

func (p *TileCoord) String() string {
	if p.N == 1 {
		return fmt.Sprintf("%d/%d/%d", p.Z, p.X, p.Y)
	}
	return fmt.Sprintf("%d/%d/%d:+%d/+%d", p.Z, p.X, p.Y, p.N, p.N)
}

/*
type Tile struct {
	TileCoord
	contentEncoding string
	contentType string
	data string
	err error
}

type TileGrid interface {
	maxExtent GeoExtent
	tileSize Size
	flipY bool
	Children(tileCoord TileCoord) chan TileCoord
	Extent(tileCoord TileCoord, border int) GeoExtent
	Parent(tileCoord TileCoord) TileCoord
	Roots() chan TileCoord
	TileCoord(z int, geoPoint GeoPoint) TileCoord
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
