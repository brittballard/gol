package gol

import "testing"

func ThreeThreeNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = Cell{x: 0, y: 0, alive: true}
  b.cells["0:1"] = Cell{x: 0, y: 1, alive: true}
  b.cells["0:2"] = Cell{x: 0, y: 2, alive: true}
  if !Three(c, b) {
    t.Fail()
  }
}

func FourThreeNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = Cell{x: 0, y: 0, alive: true}
  b.cells["0:1"] = Cell{x: 0, y: 1, alive: true}
  b.cells["0:2"] = Cell{x: 0, y: 2, alive: true}
  b.cells["1:2"] = Cell{x: 1, y: 2, alive: true}
  if Three(c, b) {
    t.Fail()
  }
}
