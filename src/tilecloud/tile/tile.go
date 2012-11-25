package tile

type Tile interface {
	Coord() Coord
	ContentEncoding() string
	ContentType() string
	Data() []byte
	Error() string
}
