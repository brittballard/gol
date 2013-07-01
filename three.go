package gol

func Three(c Cell, b Board) (dead bool) {
  return b.LivingNeighborCount(c) == 3
}
