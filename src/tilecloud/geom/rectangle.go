package geom

type Rectangle struct {
	Min, Max Point
}

func (r *Rectangle) Contains(p *Point) bool {
	return r.Min.X <= p.X && p.X <= r.Max.X && r.Min.Y <= p.Y && p.Y <= r.Max.Y
}

func (r *Rectangle) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r *Rectangle) Width() float64 {
	return r.Max.X - r.Min.X
}