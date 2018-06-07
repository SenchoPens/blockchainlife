package life_test

import (
	"github.com/SenchoPens/blockchainlife/life"
	"reflect"
	"testing"
)

func testField(t *testing.T, fieldKind life.FieldKind) {
	size := life.FieldInt(8)
	c := life.Coordinate{X: 7, Y: 7}
	celledEmpty := *life.NewCells(size)
	celledSet := *life.NewCells(size)
	celledSet[c.Y][c.X] = life.Alive

	var f life.Field

	switch fieldKind {
	case life.StateFieldKind:
		t.Log("Testing field life.StateField")
		f = life.NewStateField(size)
	case life.BlockFieldKind:
		t.Log("Testing field life.BlockField")
		f = life.NewBlockField(size)
	}

	if !reflect.DeepEqual(*f.Cells(), celledEmpty) {
		t.Error("Field method Cells() fails on empty field")
	}
	f.Set(c, life.Alive)
	if !reflect.DeepEqual(*f.Cells(), celledSet) {
		t.Error("Field method Cells() fails on non-empty field after using method Set with state life.Alive")
	}
	f.Set(c, life.Dead)
	if !reflect.DeepEqual(*f.Cells(), celledEmpty) {
		t.Error("Field method Cells() fails on empty field after using method Set with state life.Dead")
	}
}

func TestStateField(t *testing.T) { testField(t, life.StateFieldKind) }
func TestBlockField(t *testing.T) { testField(t, life.BlockFieldKind) }
