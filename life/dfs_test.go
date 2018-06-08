package life_test

import (
	"github.com/SenchoPens/blockchainlife/life"
	"reflect"
	"testing"
)

func TestBlockField_BlockRegion(t *testing.T) {
	size := life.FieldInt(64)
	f := life.NewBlockField(size)

	nonEmptyBlocks := []life.Coordinate{
		life.Coordinate{X: 0, Y: 0},
		life.Coordinate{X: 1, Y: 0},
		life.Coordinate{X: 2, Y: 1},
	}
	for _, c := range nonEmptyBlocks {
		block := f.Block(c)
		block.SetData(life.BlockData(1))
	}

	region := f.BlockRegion(life.Coordinate{X: 0, Y: 0})
	if !reflect.DeepEqual(region, nonEmptyBlocks) {
		t.Error("BlockField's BlockRegion method returns wrong region")
	}
}
