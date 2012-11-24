package tile

type Bounds struct {
	MinX, MinY, MaxX, MaxY int
}

func (b *Bounds) Contains(c *Coord) bool {
	return b.MinX <= c.X && c.X < b.MaxX && b.MinY <= c.Y && c.Y < b.MaxY
}

func (b *Bounds) Each(z int) <-chan Coord {
	ch := make(chan Coord)
	go func(z, minX, minY, maxX, maxY int) {
		for x := minX; x < maxX; x++ {
			for y := minY; y < maxY; y++ {
				ch <- Coord{z, x, y, 1}
			}
		}
		close(ch)
	}(z, b.MinX, b.MinY, b.MaxX, b.MaxY)
	return ch
}

func (b *Bounds) Height() int {
	return b.MaxY - b.MinY
}

func (b *Bounds) Width() int {
	return b.MaxX - b.MinX
}
