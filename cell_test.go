package gol

import "testing"

func TestIsAlive(t *testing.T) {
  c := Cell{x: 1, y: 2}
  c.SwapXAndY()
  if c.x == 1 {
    t.Error("No deal!")
  }
}

func TestDoTheSwap(t *testing.T) {
  c := &Cell{x: 5, y: 4}
  DoTheSwap(c)
  if c.x == 5 {
    t.Error("No deal!")
  }
}
