package tilecloud

import (
	"fmt"
)

type TileBounds struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func (p *TileBounds) Contains(tileCoord *TileCoord) bool {
	return p.MinX <= tileCoord.X && tileCoord.X < p.MaxX && p.MinY <= tileCoord.Y && tileCoord.Y < p.MaxY
}

func (p *TileBounds) Height() int {
	return p.MaxX - p.MinX
}

func (p *TileBounds) Each(z int) chan TileCoord {
	ch := make(chan TileCoord)
	go func() {
		for x := p.MinX; x < p.MaxX; x++ {
			for y := p.MinY; y < p.MaxY; y++ {
				ch <- TileCoord{z, x, y, 1}
			}
		}
		close(ch)
	}()
	return ch
}

func (p *TileBounds) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", p.MinX, p.MinY, p.MaxX, p.MaxY)
}

func (p *TileBounds) Width() int {
	return p.MaxY - p.MinY
}
