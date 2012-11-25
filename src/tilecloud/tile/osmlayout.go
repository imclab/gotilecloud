package tile

import (
	"fmt"
	"regexp"
	"strconv"
)

type OSMLayout struct {
}

func (ol *OSMLayout) Filename(c Coord) string {
	return fmt.Sprintf("%d/%d/%d", c.Z, c.X, c.Y)
}

var re = regexp.MustCompile("\\A([0-9]+)/([0-9]+)/([0-9]+)\\z")

func (ol *OSMLayout) Coord(s string) (Coord, bool) {
	m := re.FindStringSubmatch(s)
	if m == nil {
		return Coord{}, false
	}
	z, _ := strconv.Atoi(m[1])
	x, _ := strconv.Atoi(m[2])
	y, _ := strconv.Atoi(m[3])
	return Coord{z, x, y, 1}, true
}
