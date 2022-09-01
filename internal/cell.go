package internal

import "fmt"

type Cell byte

const (
	OpenCell   Cell = 8<<iota + 1 // 9
	ClosedCell                    // 17
	BombCell                      // 33

	OpenBombCell = OpenCell | BombCell // 41
)

func (c Cell) String() string {
	if c&ClosedCell == ClosedCell {
		return "*"
	}

	if c&OpenBombCell == OpenBombCell {
		return "B"
	}

	if c&OpenCell == OpenCell {
		return "#"
	}

	if c < OpenCell {
		res := fmt.Sprintf("%d", c)
		return res
	}

	return ""
}
