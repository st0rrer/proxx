package internal

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

type Encoder struct {
	openAll bool
}

func NewEncoder(openAll bool) *Encoder {
	return &Encoder{
		openAll: openAll,
	}
}

func (e *Encoder) Encode(b *Board) (string, error) {
	buf := bytes.Buffer{}
	wr := tabwriter.NewWriter(&buf, 0, 0, 1, '|', tabwriter.TabIndent)

	for _, row := range b.board {
		cells := make([]string, 0, len(row))
		for _, cell := range row {
			if cell&BombCell == BombCell && e.openAll {
				cell = OpenBomb
			}

			cells = append(cells, cell.String())
		}
		if _, err := fmt.Fprintln(wr, strings.Join(cells, "\t")); err != nil {
			return "", fmt.Errorf("failed to add row: %w", err)
		}
	}
	if err := wr.Flush(); err != nil {
		return "", fmt.Errorf("failed to flush buffer: %w", err)
	}

	if _, err := fmt.Fprintf(&buf, "\nRemaining number of empty cells: %d\n", b.emptyCells); err != nil {
		return "", fmt.Errorf("failed to add the remaining number of empty cells: %w", err)
	}
	return buf.String(), nil
}
