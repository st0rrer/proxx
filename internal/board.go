package internal

import (
	"math"
	"math/rand"
)

var directions = [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

type Board struct {
	board      [][]Cell
	size       int
	emptyCells int
}

func NewBoard(size int, bombs int) *Board {
	board := make([][]Cell, size)
	for i := 0; i < size; i++ {
		row := make([]Cell, size)
		for j := 0; j < size; j++ {
			row[j] = ClosedCell
		}
		board[i] = row
	}

	cellCount := int(math.Pow(float64(size), 2))
	bombSet := make(map[int]struct{})
	for i := 0; i < bombs; i++ {
		cell := rand.Intn(cellCount)
		bombSet[cell] = struct{}{}
	}

	for b := range bombSet {
		x := b / size
		y := b % size

		board[x][y] |= BombCell
	}

	return &Board{
		board:      board,
		size:       len(board),
		emptyCells: cellCount - len(bombSet),
	}
}

func (b *Board) IsBomb(x, y int) bool {
	cell := b.board[x][y]
	return cell&BombCell == BombCell
}

func (b *Board) Open(x, y int) int {
	size := len(b.board)

	q := NewQueue()
	q.Push(size*x + y)

	for !q.IsEmpty() {
		point1D := q.Pop()
		cellX := point1D / size
		cellY := point1D % size

		if b.isOpen(cellX, cellY) {
			continue
		}

		adjacentBombs, adjacentBFS := b.openAdjacentCell(cellX, cellY)

		b.emptyCells--
		b.board[cellX][cellY] = OpenCell
		if adjacentBombs != 0 {
			b.board[cellX][cellY] = Cell(adjacentBombs)
			continue
		}

		for !adjacentBFS.IsEmpty() {
			adjacentPoint1D := adjacentBFS.Pop()
			q.Push(adjacentPoint1D)
		}
	}

	return b.emptyCells
}

func (b *Board) openAdjacentCell(x, y int) (adjacent byte, bfs *Queue) {
	bfs = NewQueue()
	for _, d := range directions {
		adjacentX := x + d[0]
		adjacentY := y + d[1]

		if adjacentX > -1 && adjacentX < b.size && adjacentY > -1 && adjacentY < b.size {
			if b.IsBomb(adjacentX, adjacentY) {
				adjacent++
			}
			bfs.Push(b.size*adjacentX + adjacentY)
		}
	}
	return
}

func (b *Board) isOpen(x, y int) bool {
	cell := b.board[x][y]
	return cell <= OpenCell
}
