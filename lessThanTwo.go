package gol

func LessThanTwo(c Cell, b Board) (dead bool) {
  return b.LivingNeighborCount(c) < 2
}
