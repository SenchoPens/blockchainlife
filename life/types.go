package life

// Max map size - ~100 megabytes
type FieldInt uint16

type Coordinate struct {
	x, y FieldInt
}

// Unique user identifier
type User string

// Map from cells that are alive to users that own them
type UserMap map[Coordinate]User
