package life

// Block is a 8*8 square of bits, 0 - cell is dead, 1 - it is alive.
type Block uint64

// blockSize defines size, where size*size = block.
// It is 8 because 8 * 8 = 64 bits, which is the size of the block.
const blockSize FieldInt = 8

// BlockField is a square field of blocks.
type BlockField struct {
	blocks [][]Block
	// size is size of the field in blocks.
	size FieldInt
	// fullSize is size of the field in cells.
	fullSize FieldInt
}

// block returns field's block in which the coordinate c locates.
func (f *BlockField) block(c Coordinate) *Block {
	return &f.blocks[c.Y/blockSize][c.X/blockSize]
}

// Size returns the size of BlockField if it was represented in cells.
func (f *BlockField) Size() FieldInt {
	return f.fullSize
}

// NewField makes new BlockField with size in blocks equal to biggest a, such that a * 8 <= size.
func NewBlockField(size FieldInt) *BlockField {
	f := BlockField{}
	f.size = size / 8
	f.blocks = make([][]Block, f.size)
	for i := range f.blocks {
		f.blocks[i] = make([]Block, f.size)
	}
	f.fullSize = f.size * 8
	return &f
}

func (f *BlockField) calcCoordinateBlockBit(c Coordinate) uint {
	return uint(8*(c.Y%8) + c.X%8)
}

// cell returns state of cell at given Coordinate, not handling fieldType.
func (f *BlockField) cell(c Coordinate) State {
	return (*f.block(c)>>f.calcCoordinateBlockBit(c))%2 == 1
}

func (f *BlockField) Cells() *Cells {
	fullSize := f.size * 8
	cells := *NewCells(fullSize)

	for x := FieldInt(0); x < fullSize; x++ {
		for y := FieldInt(0); y < fullSize; y++ {
			cells[y][x] = f.cell(Coordinate{x, y})
		}
	}

	return &cells
}

func (f *BlockField) Set(c Coordinate, state State) {
	block := f.block(c)
	bit := Block(1 << f.calcCoordinateBlockBit(c))
	if state {
		*block |= bit
	} else {
		*block &= ^bit
	}
}

func (f *BlockField) String() string {
	return f.Cells().String()
}