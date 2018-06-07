package life

import "bytes"

// Cells represents a square 2D array of State of each cell.
type Cells [][]State

func NewCells(size FieldInt) *Cells {
	cells := make(Cells, size)
	for i := range cells {
		cells[i] = make([]State, size)
	}
	return &cells
}

// String converts cells to string as a table, in which "-" symbol means that cell is dead, and "*" means it is alive.
func (cells *Cells) String() string {
	var buf bytes.Buffer
	for _, row := range *cells {
		for _, c := range row {
			b := byte('-')
			if c {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
