package life

import "fmt"

// Field represents the life's game field - it stores the state of each cell.
type Field interface {
	fmt.Stringer
	// Cells returns a pointer to a square 2D array of State of each cell.
	Cells() *Cells
	// Set sets the state of a coordinate of the field.
	// It does not check whether the coordinate is out of bounds, the code that calls the function must do that itself.
	Set(Coordinate, State)
	// Size returns size of the field in cells, such that size * size == number of cells.
	Size() FieldInt
	// cell returns the state of a coordinate.
	// It does not check whether the coordinate is out of bounds, the code that calls the function must do that itself.
	cell(Coordinate) State
}

type Life interface {
	// Field returns Life's Dead / Alive field of state of cells, as a pointer.
	Field() Field
	// Tick computes each cell
	Tick()
	// Users returns a pointer to a map from a coordinate of living cell to it's Owner's id.
	Users() *UserMap
}
