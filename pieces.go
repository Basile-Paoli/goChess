package main

type PieceType int

const (
	PAWN PieceType = iota
	KNIGHT
	BISHOP
	ROOK
	QUEEN
	KING
)

type Piece interface {
	Color() Color
	Type() PieceType
	Symbol() string
	LegalMoves(game *Game, from *Square) []Move
}
type Pawn struct {
	color Color
}

func (p Pawn) Color() Color {
	return p.color
}
func (p Pawn) Type() PieceType {
	return PAWN
}
func (p Pawn) Symbol() string {
	if p.color == White {
		return "♟"
	}
	return "♙"
}

func (p Pawn) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	row, column := from[0], from[1]
	if p.color == White {
		if row == row8 {
			return moves
		}
		if game.board[row+1][column] == nil {
			moves = append(moves, Move{from, &Square{row + 1, column}})

			if row == row2 && game.board[row+2][column] == nil {
				moves = append(moves, Move{from, &Square{row + 2, column}})
			}
		}
		if column > columnA && game.board[row+1][column-1] != nil && game.board[row+1][column-1].Color() == Black {
			moves = append(moves, Move{from, &Square{row + 1, column - 1}})
		}
		if column < columnH && game.board[row+1][column+1] != nil && game.board[row+1][column+1].Color() == Black {
			moves = append(moves, Move{from, &Square{row + 1, column + 1}})
		}
	}
	if p.color == Black {
		if row == row1 {
			return moves
		}
		if game.board[row-1][column] == nil {
			moves = append(moves, Move{from, &Square{row - 1, column}})
			if row == row7 && game.board[row-2][column] == nil {
				moves = append(moves, Move{from, &Square{row - 2, column}})
			}
		}
		if column > columnA && game.board[row-1][column-1] != nil && game.board[row-1][column-1].Color() == White {
			moves = append(moves, Move{from, &Square{row - 1, column - 1}})
		}
		if column < columnH && game.board[row-1][column+1] != nil && game.board[row-1][column+1].Color() == White {
			moves = append(moves, Move{from, &Square{row - 1, column + 1}})
		}

	}
	return moves
}

type Knight struct {
	color Color
}

func (k Knight) Color() Color {
	return k.color
}
func (k Knight) Type() PieceType {
	return KNIGHT
}
func (k Knight) Symbol() string {
	if k.color == White {
		return "♞"
	}
	return "♘"
}
func (k Knight) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	row, column := from[0], from[1]
	possibleSquares := []Square{
		{row + 2, column + 1},
		{row + 2, column - 1},
		{row + 1, column + 2},
		{row + 1, column - 2},
		{row - 2, column + 1},
		{row - 2, column - 1},
		{row - 1, column + 2},
		{row - 1, column - 2},
	}
	for _, square := range possibleSquares {
		row, column := square[0], square[1]
		if row1 <= row && row <= row8 && columnA <= column && column <= columnH {
			if game.board[row][column] == nil || game.board[row][column].Color() != k.color {
				moves = append(moves, Move{from, &square})
			}
		}
	}
	return moves
}

type Bishop struct {
	color Color
}

func (b Bishop) Color() Color {
	return b.color
}
func (b Bishop) Type() PieceType {
	return BISHOP
}
func (b Bishop) Symbol() string {
	if b.color == White {
		return "♝"
	}
	return "♗"
}

func (b Bishop) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		for row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			if game.board[destRow][destColumn] == nil {
				moves = append(moves, Move{from, &Square{destRow, destColumn}})
			} else {
				if game.board[destRow][destColumn].Color() != b.color {
					moves = append(moves, Move{from, &Square{destRow, destColumn}})
				}
				break
			}
			destRow += direction[0]
			destColumn += direction[1]
		}

	}

	return moves
}

type Rook struct {
	color Color
}

func (r Rook) Color() Color {
	return r.color
}
func (r Rook) Type() PieceType {
	return ROOK
}
func (r Rook) Symbol() string {
	if r.color == White {
		return "♜"
	}
	return "♖"
}
func (r Rook) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		for row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			if game.board[destRow][destColumn] == nil {
				moves = append(moves, Move{from, &Square{destRow, destColumn}})
			} else {
				if game.board[destRow][destColumn].Color() != r.color {
					moves = append(moves, Move{from, &Square{destRow, destColumn}})
				}
				break
			}
			destRow += direction[0]
			destColumn += direction[1]
		}
	}
	return moves
}

type Queen struct {
	color Color
}

func (q Queen) Color() Color {
	return q.color
}
func (q Queen) Type() PieceType {
	return QUEEN
}
func (q Queen) Symbol() string {
	if q.color == White {
		return "♛"
	}
	return "♕"
}

func (q Queen) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		for row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			if game.board[destRow][destColumn] == nil {
				moves = append(moves, Move{from, &Square{destRow, destColumn}})
			} else {
				if game.board[destRow][destColumn].Color() != q.color {
					moves = append(moves, Move{from, &Square{destRow, destColumn}})
				}
				break
			}
			destRow += direction[0]
			destColumn += direction[1]
		}
	}
	return moves
}

type King struct {
	color Color
}

func (k King) Color() Color {
	return k.color
}
func (k King) Type() PieceType {
	return KING
}
func (k King) Symbol() string {
	if k.color == White {
		return "♚"
	}
	return "♔"
}
func (k King) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		if row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			if game.board[destRow][destColumn] == nil || game.board[destRow][destColumn].Color() != k.color {
				moves = append(moves, Move{from, &Square{destRow, destColumn}})
			}
		}
	}
	return moves
}
