// Naive implementation of life using 2D array of bool state
package life

import (
	"math/rand"
)

type FieldKind int

const (
	StateFieldKind FieldKind = iota + 1
	BlockFieldKind
)

type Naive struct {
	field     Field
	fieldType FieldType
	users     *UserMap
}

func (l *Naive) Field() Field {
	return l.field
}

func (l *Naive) Users() *UserMap {
	return l.users
}

func NewLifeNaive(f Field, fieldType FieldType) *Naive {
	u1 := make(UserMap)
	return &Naive{
		field:     f,
		fieldType: fieldType,
		users:     &u1,
	}
}

// pickDominant chooses the one who is dominant between three not necessarily unique users.
func pickDominant(a, b, c User) User {
	// Literally the best solution I came with.
	switch {
	case a == b:
		return a
	case b == c:
		return b
	case a == c:
		return a
	default:
		// I hope no one will speculate on math/rand, but it works faster than crypto/rand
		return [3]User{a, b, c}[rand.Intn(3)]
	}
}

// Alive returns state of the cell with given coordinate c, handling various types of fields.
func (l *Naive) alive(c Coordinate) State {
	switch l.fieldType {
	case Toroidal:
		return l.field.cell(c.projectPlainOnTor(l.field.Size()))
	case Flat:
		if c.inSquare(l.field.Size()) {
			return Dead
		}
		return l.field.cell(c)
	default: // should never be triggered
		return Dead
	}
}

// getCellNeighbors returns a slice of cell's alive neighbors.
func (l *Naive) getCellNeighbors(c Coordinate) *[]User {
	var aliveNeighbors []User
	for dx := FieldInt(-1); dx <= 1; dx++ {
		for dy := FieldInt(-1); dy <= 1; dy++ {
			neighborC := Coordinate{c.X + dx, c.Y + dy}
			if (dx != 0 || dy != 0) && l.alive(neighborC) {
				aliveNeighbors = append(aliveNeighbors, (*l.users)[neighborC])
			}
		}
	}
	return &aliveNeighbors
}

func (l *Naive) Tick() {
	// Following 2 slices make it possible to avoid having 2 maps of users and 2 fields.
	var cellsToKill []Coordinate
	var cellsToVivify []OwnedCell

	for i := FieldInt(0); i < l.Field().Size(); i++ {
		for j := FieldInt(0); j < l.Field().Size(); j++ {
			c := Coordinate{X: i, Y: j}
			aliveNeighbors := *l.getCellNeighbors(c)

			if l.alive(c) && (len(aliveNeighbors) < 2 || len(aliveNeighbors) > 3) {
				cellsToKill = append(cellsToKill, c)
			}

			if !l.alive(c) && len(aliveNeighbors) == 3 {
				cellOwner := pickDominant(aliveNeighbors[0], aliveNeighbors[1], aliveNeighbors[2])
				cellsToVivify = append(cellsToVivify, OwnedCell{c, cellOwner})
			}
		}
	}
	for _, c := range cellsToKill {
		l.field.Set(c, Dead)
		delete(*l.users, c)
	}

	for _, oc := range cellsToVivify {
		l.field.Set(oc.C, Alive)
		(*l.users)[oc.C] = oc.Owner
	}
}
