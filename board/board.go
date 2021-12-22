package board

import (
	"github.com/maaslalani/gambit/piece"
)

type Board struct {
	// The board is represented as a 2D array of cells.
	// The first index is the row, the second is the column.
	grid     [8][8]piece.Piece
	reversed bool
}

func New() Board {
	grid := [8][8]piece.Piece{
		{piece.Piece{Type: piece.King, Color: piece.White}, piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
		{piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty(), piece.Empty()},
	}
	return Board{grid: grid}
}
