package life

type FieldType int

const (
	// Toroidal is a field type where field is like a tor.
	Toroidal FieldType = iota + 1
	// Flat is when cells that are out of the corners of the field are treated as dead.
	Flat
)

// StateField represents a square 2D array of cells state and it's size.
type StateField struct {
	// Height and width of field.
	size  FieldInt
	cells Cells
}

// NewField initializes a StateField structure with an empty matrix of cells and given size
func NewStateField(size FieldInt) *StateField {
	return &StateField{size: size, cells: *NewCells(size)}
}

func (f *StateField) Cells() *Cells {
	return &f.cells
}

func (f *StateField) String() string {
	return f.cells.String()
}

func (f *StateField) Set(c Coordinate, value State) {
	f.cells[c.Y][c.X] = value
}

func (f *StateField) Size() FieldInt {
	return f.size
}

func (f *StateField) cell(c Coordinate) State {
	return f.cells[c.Y][c.X]
}
