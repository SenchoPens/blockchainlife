package life

// State of a cell.
type State bool

const (
	Dead  State = false
	Alive State = true
)

type FieldInt int32

type Coordinate struct {
	X, Y FieldInt
}

// Unique Owner identifier.
type User string

// Map from cells that are Alive to users that own them.
type UserMap map[Coordinate]User

type OwnedCell struct {
	C     Coordinate
	Owner User
}
