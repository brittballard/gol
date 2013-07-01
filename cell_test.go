package gol

import (
  "testing"
  "fmt"
)

func TestGetNeighbors(t *testing.T) {
  c := Cell{x: 3, y: 2}
  got := fmt.Sprintf("%v", c.GetNeighbors())
  expected := fmt.Sprintf("%v", []string{"2:2", "2:1", "3:1", "4:1", "4:2", "4:3", "3:3", "2:3"})
  if got != expected {
    t.Fail()
  }
}
