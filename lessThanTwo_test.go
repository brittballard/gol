package gol

import "testing"

func TestNoNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  if !LessThanTwo(c, b) {
    t.Fail()
  }
}

func TestOneNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = Cell{x: 0, y: 0}
  if !LessThanTwo(c, b) {
    t.Fail()
  }
}

func TestTwoNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = Cell{x: 0, y: 0}
  b.cells["0:1"] = Cell{x: 0, y: 1}
  if LessThanTwo(c, b) {
    t.Fail()
  }
}
