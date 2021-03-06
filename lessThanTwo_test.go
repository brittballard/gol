package gol

import "testing"

func TestLessThanTwoNoNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1, alive: true}
  b := Board{cells: make(map[string]Cell)}
  if !LessThanTwo(c, b) {
    t.Fail()
  }
}

func TestLessThanTwoOneNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1, alive: true}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = c
  if !LessThanTwo(c, b) {
    t.Fail()
  }
}

func TestLessThanTwoTwoNeighbors(t *testing.T) {
  c := Cell{x: 1, y: 1}
  b := Board{cells: make(map[string]Cell)}
  b.cells["0:0"] = Cell{x: 0, y: 0, alive: true}
  b.cells["0:1"] = Cell{x: 0, y: 1, alive: true}
  if LessThanTwo(c, b) {
    t.Fail()
  }
}
