package gol

import (
  "fmt"
)

type Cell struct {
  x int
  y int
}

func (c Cell) GetNeighbors() ([]string) {
  neighbors := []string{}
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x - 1, c.y))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x - 1, c.y - 1))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x, c.y - 1))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x + 1, c.y - 1))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x + 1, c.y))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x + 1, c.y + 1))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x, c.y + 1))
  neighbors = append(neighbors, fmt.Sprintf("%d:%d", c.x - 1, c.y + 1))
  return neighbors
}
