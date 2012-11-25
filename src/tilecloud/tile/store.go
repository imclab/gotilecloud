package tile

type Store interface {
	Contains(t Tile) bool
	Count() int
	Delete(t Tile) Tile
	Get(t Tile) Tile
	GetAll() <-chan Tile
	GetOne(t Tile) Tile
	List() <-chan Tile
	PutOne(t Tile) Tile
	Put(ch <-chan Tile) <-chan Tile
}
