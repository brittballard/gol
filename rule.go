package gol

type Rule interface {
  Applicable(c Cell, b Board)
}
