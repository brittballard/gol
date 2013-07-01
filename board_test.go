package gol

import "testing"

func TestLivingNeighborCount(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = Cell{x: 0, y: 0, alive: true}
  b.cells["0:1"] = Cell{x: 0, y: 1, alive: true}
  if b.LivingNeighborCount(c) != 2 {
    t.Fail()
  }
}
