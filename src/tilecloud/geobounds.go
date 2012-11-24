package tilecloud

import (
	"fmt"
)

type GeoBounds struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func (p *GeoBounds) Contains(geoCoord *GeoCoord) bool {
	return p.MinX <= geoCoord.X && geoCoord.X <= p.MaxX && p.MinY <= geoCoord.Y && geoCoord.Y <= p.MaxY
}

func (p *GeoBounds) Height() float64 {
	return p.MaxY - p.MinY
}

func (p *GeoBounds) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", p.MinX, p.MinY, p.MaxX, p.MaxY)
}

func (p *GeoBounds) Width() float64 {
	return p.MaxX - p.MinX
}
