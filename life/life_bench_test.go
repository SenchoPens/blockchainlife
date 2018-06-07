package life_test

import (
	"github.com/SenchoPens/blockchainlife/life"
	"math/rand"
	"testing"
)

func fillRandomly(l life.Life, density int) {
	users := [3]life.User{"0", "1", "2"}
	for x := life.FieldInt(0); x < l.Field().Size(); x++ {
		for y := life.FieldInt(0); y < l.Field().Size(); y++ {
			if rand.Intn(density) == 0 {
				c := life.Coordinate{x, y}
				l.Field().Set(c, life.Alive)
				(*l.Users())[c] = users[rand.Intn(3)]
			}
		}
	}
}

func benchmarkLifeNaive_Tick(b *testing.B, fieldKind life.FieldKind) {
	size := life.FieldInt(1 << 10)

	var f life.Field
	switch fieldKind {
	case life.StateFieldKind:
		f = life.NewStateField(size)
	case life.BlockFieldKind:
		f = life.NewBlockField(size)
	}

	l := life.NewLifeNaive(f, life.Flat)
	fillRandomly(l, 3)

	b.ResetTimer()
	for x := 0; x < 10; x++ {
		l.Tick()
	}
}

func BenchmarkLifeNaiveStateField_Tick(b *testing.B) { benchmarkLifeNaive_Tick(b, life.StateFieldKind) }
func BenchmarkLifeNaiveBlockField_Tick(b *testing.B) { benchmarkLifeNaive_Tick(b, life.BlockFieldKind) }
