package tilecloud

import (
	"fmt"
)

type GeoCoord struct {
	X float64
	Y float64
}

func (p *GeoCoord) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}
