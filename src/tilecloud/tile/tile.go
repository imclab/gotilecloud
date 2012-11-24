package tile

type Tile struct {
	Coord
	contentEncoding string
	contentType     string
	data            string
	err             error
}

func (p *Tile) Error() string {
	return p.err.Error()
}
