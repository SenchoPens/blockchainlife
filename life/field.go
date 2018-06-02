// field.go provides Field struct and NewField function to initialize an empty field
package life

// Field represents a matrix of cells state and it's size
type Field struct {
	size  uint64 // height and width of field
	cells [][]bool
}

// NewField initializes a Field structure with an empty matrix of cells and given size
func NewField(size uint64) *Field {
	cells := make([][]bool, size)
	for i := range cells {
		cells[i] = make([]bool, size)
	}
	return &Field{size: size, cells: cells}
}
