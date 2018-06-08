package life

// BlockData is a 8*8 square of bits, 0 - cell is dead, 1 - it is alive.
type BlockData uint64

// Timestamp is the number of block's passed generations
type Timestamp uint16

// Block is BlockData with a Timestamp on it
type Block struct {
	data BlockData
	t    Timestamp
}

func (b *Block) Data() BlockData {
	return b.data
}

func (b *Block) SetData(newData BlockData) {
	b.data = newData
}

func (b *Block) IncreaseTimestamp(value Timestamp) {
	b.t += value
}

func (b *Block) Timestamp() Timestamp {
	return b.t
}

func (b *Block) Empty() bool {
	return b.data == BlockData(0)
}

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

func (f *BlockField) BlockSize() FieldInt {
	return f.size
}

// coordinateBlock returns field's block in which the coordinate c locates.
func (f *BlockField) coordinateBlock(c Coordinate) *Block {
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
	return ((*f.coordinateBlock(c)).data>>f.calcCoordinateBlockBit(c))%2 == 1
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
	block := f.coordinateBlock(c)
	bit := BlockData(1 << f.calcCoordinateBlockBit(c))
	if state {
		(*block).data |= bit
	} else {
		(*block).data &= ^bit
	}
}

func (f *BlockField) String() string {
	return f.Cells().String()
}

// Block returns a pointer to a block with given Coordinate.
func (f *BlockField) Block(c Coordinate) *Block {
	return &f.blocks[c.Y][c.X]
}
