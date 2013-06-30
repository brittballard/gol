package gol

import "testing"

type testSwaper struct {
  callCount int
}

func (t *testSwaper) SwapXAndY() {
  t.callCount++
}

func TestSwapXAndY(t *testing.T) {
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

func TestDoTheSwapWithMock(t *testing.T) {
  ts := &testSwaper{}
  DoTheSwap(ts)
  if ts.callCount != 1 {
    t.Error("No deal!")
  }
}
