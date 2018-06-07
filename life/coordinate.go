package life

func (c Coordinate) projectPlainOnTor(size FieldInt) Coordinate {
	c.Y = (c.Y + size) % size
	c.X = (c.X + size) % size
	return c
}

func (c Coordinate) inSquare(size FieldInt) bool {
	return c.X < 0 || c.X >= size || c.Y < 0 || c.Y >= size
}
