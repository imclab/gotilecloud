package tilecloud

type Tile struct {
	TileCoord
	contentEncoding string
	contentType     string
	data            string
	err             error
}

func (p *Tile) Error() string {
	return p.err.Error()
}
