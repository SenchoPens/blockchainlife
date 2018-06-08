package life

// BlockRegion returns coordinates of all blocks that are bounded with the given block
func (f *BlockField) BlockRegion(c Coordinate) []Coordinate {
	visited := make(map[Coordinate]bool)
	region := make([]Coordinate, 0)
	f.blockDFS(c, &visited, &region)
	return region
}

func (f *BlockField) blockDFS(c Coordinate, visited *map[Coordinate]bool, region *[]Coordinate) {
	(*visited)[c] = true
	*region = append(*region, c)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				x := c.X + FieldInt(i)
				y := c.Y + FieldInt(j)
				neighborC := Coordinate{X: x, Y: y}
				if !(*visited)[neighborC] && neighborC.inSquare(f.BlockSize()) {
					if !f.Block(neighborC).Empty() {
						f.blockDFS(neighborC, visited, region)
					}
				}
			}
		}
	}
}
