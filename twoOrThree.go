package gol

func TwoOrThree(c Cell, b Board) (dead bool) {
  livingNeighbors := b.LivingNeighborCount(c)
  return livingNeighbors == 2 || livingNeighbors == 3
}
