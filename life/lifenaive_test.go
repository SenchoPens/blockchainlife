package life_test

import (
	"github.com/SenchoPens/blockchainlife/life"
	"reflect"
	"testing"
)

func vivifyOwnedCells(l life.Life, cells []life.OwnedCell) {
	for _, cell := range cells {
		l.Field().Set(cell.C, life.Alive)
		(*l.Users())[cell.C] = cell.Owner
	}
}

func testLifeNaive(t *testing.T, fieldKind life.FieldKind) {
	size := life.FieldInt(8)

	livingCells := []life.OwnedCell{
		{life.Coordinate{X: 0, Y: 0}, "0"},

		{life.Coordinate{X: 6, Y: 6}, "0"},
		{life.Coordinate{X: 6, Y: 5}, "0"},
		{life.Coordinate{X: 7, Y: 6}, "0"},
		{life.Coordinate{X: 6, Y: 7}, "0"},
		{life.Coordinate{X: 5, Y: 6}, "1"},
	}
	nextCells := []life.OwnedCell{
		{life.Coordinate{X: 5, Y: 5}, "0"},
		{life.Coordinate{X: 6, Y: 5}, "0"},
		{life.Coordinate{X: 7, Y: 5}, "0"},
		{life.Coordinate{X: 7, Y: 6}, "0"},
		{life.Coordinate{X: 7, Y: 7}, "0"},
		{life.Coordinate{X: 6, Y: 7}, "0"},
		{life.Coordinate{X: 5, Y: 7}, "0"},
		{life.Coordinate{X: 5, Y: 6}, "1"},
	}

	var f1, f2 life.Field
	switch fieldKind {
	case life.StateFieldKind:
		t.Log("Testing life.Naive with life.StateField as field")
		f1 = life.NewStateField(size)
		f2 = life.NewStateField(size)
	case life.BlockFieldKind:
		t.Log("Testing life.Naive with life.BlockField as field")
		f1 = life.NewStateField(size)
		f2 = life.NewStateField(size)
	}

	// Tests only life.Flat
	l1 := life.NewLifeNaive(f1, life.Flat)
	l2 := life.NewLifeNaive(f2, life.Flat)

	// Following test does not cover only random ownership
	vivifyOwnedCells(l1, livingCells)

	t.Log("Field 1 generation 1:")
	t.Log(l1.Field())

	l1.Tick()

	vivifyOwnedCells(l2, nextCells)

	t.Log("Field 1 generation 2:")
	t.Log(l1.Field())
	t.Log("Field 2 generation 1:")
	t.Log(l2.Field())

	if !reflect.DeepEqual(*l1.Field().Cells(), *l2.Field().Cells()) {
		t.Error("Fields do not match")
	}
	if !reflect.DeepEqual(*l1.Users(), *l2.Users()) {
		t.Error("Users do not match")
	}
}

func TestLifeNaive_StateField(t *testing.T) { testLifeNaive(t, life.StateFieldKind) }
func TestLifeNaive_BlockField(t *testing.T) { testLifeNaive(t, life.BlockFieldKind) }
