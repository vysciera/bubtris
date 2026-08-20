package game

import "testing"

func TestNewBoardIsEmpty(t *testing.T) {
	board := Board()

	if len(board) != 20 {
		t.Fatalf("expected 20 rows, got %d", len(board))
	}

	for y, row := range board {
		if len(row) != 10 {
			t.Fatalf("row %d: expected 10 columns, got %d", y, len(row))
		}

		for x, cell := range row {
			if cell != 0 {
				t.Errorf("cell (%d, %d): expected empty, got %v", x, y, cell)
			}
		}
	}
}
