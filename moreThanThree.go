package gol

func MoreThanThree(c Cell, b Board) (dead bool) {
  return b.LivingNeighborCount(c) > 3
}
