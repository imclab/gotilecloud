package tile

type Layout interface {
	Filename(c Coord) string
	Coord(s string) (Coord, bool)
}
