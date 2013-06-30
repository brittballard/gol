package gol

type Cell struct {
  x int
  y int
}

type Swaper interface {
  SwapXAndY()
}

func (c *Cell) SwapXAndY() {
  c.x = c.y
  c.y = c.x
}

func DoTheSwap(s Swaper) {
  s.SwapXAndY()
}
