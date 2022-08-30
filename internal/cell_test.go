package internal

import "testing"

func TestCell_String(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want string
	}{
		{
			name: "Closed",
			c:    ClosedCell,
			want: "*",
		},
		{
			name: "Bomb",
			c:    ClosedCell | BombCell,
			want: "*",
		},
		{
			name: "Open",
			c:    OpenCell,
			want: "#",
		},
		{
			name: "OpenBomb",
			c:    OpenCell | BombCell,
			want: "B",
		},
		{
			name: "OneAdjacentBombs",
			c:    1,
			want: "1",
		},
		{
			name: "EightAdjacentBombs",
			c:    8,
			want: "8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
