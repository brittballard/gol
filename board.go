package gol

type Board struct {
  cells map[string]Cell
}

func (b Board) LivingNeighborCount(c Cell) (count int) {
  neighborCount := 0
  neighbors := c.GetNeighbors()
  for _, v := range neighbors {
    if b.cells[v].alive {
      neighborCount++
    }
  }
  return neighborCount
}
