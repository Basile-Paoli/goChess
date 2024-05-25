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
	Attacks(game *Game, from *Square) []Square
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
	if from == nil {
		return moves
	}
	row, column := from[0], from[1]
	if p.color == White {
		if row == row8 {
			return moves
		}
		if row == row7 && game.board[row+1][column] == nil {
			for i := KNIGHT; i <= QUEEN; i++ {
				moves = append(moves, Move{From: from, To: &Square{row + 1, column}, PromotionType: &i})
			}
		} else if game.board[row+1][column] == nil {
			moves = append(moves, Move{From: from, To: &Square{row + 1, column}})

			if row == row2 && game.board[row+2][column] == nil {
				moves = append(moves, Move{From: from, To: &Square{row + 2, column}})
			}
		}
		//en passant
		if row == row5 {
			if column > columnA && game.board[row][column-1] != nil && game.board[row][column-1].Type() == PAWN && game.lastMove != nil && game.lastMove.From[0] == row+2 && game.lastMove.To[0] == row && game.lastMove.To[1] == column-1 {
				moves = append(moves, Move{From: from, To: &Square{row + 1, column - 1}})
			}
			if column < columnH && game.board[row][column+1] != nil && game.board[row][column+1].Type() == PAWN && game.lastMove != nil && game.lastMove.From[0] == row+2 && game.lastMove.To[0] == row && game.lastMove.To[1] == column+1 {
				moves = append(moves, Move{From: from, To: &Square{row + 1, column + 1}})
			}
		}
		if column > columnA && game.board[row+1][column-1] != nil && game.board[row+1][column-1].Color() == Black {
			if row == row7 {
				for i := KNIGHT; i <= QUEEN; i++ {
					moves = append(moves, Move{From: from, To: &Square{row + 1, column - 1}, PromotionType: &i})
				}
			} else {
				moves = append(moves, Move{From: from, To: &Square{row + 1, column - 1}})
			}
		}
		if column < columnH && game.board[row+1][column+1] != nil && game.board[row+1][column+1].Color() == Black {
			if row == row7 {
				for i := KNIGHT; i <= QUEEN; i++ {
					moves = append(moves, Move{From: from, To: &Square{row + 1, column + 1}, PromotionType: &i})
				}
			} else {
				moves = append(moves, Move{From: from, To: &Square{row + 1, column + 1}})
			}
		}
	}
	if p.color == Black {
		if row == row1 {
			return moves
		}
		if row == row2 && game.board[row-1][column] == nil {
			for i := KNIGHT; i <= QUEEN; i++ {
				moves = append(moves, Move{From: from, To: &Square{row - 1, column}, PromotionType: &i})
			}
		} else if game.board[row-1][column] == nil {
			moves = append(moves, Move{From: from, To: &Square{row - 1, column}})
			if row == row7 && game.board[row-2][column] == nil {
				moves = append(moves, Move{From: from, To: &Square{row - 2, column}})
			}
		}
		//en passant
		if row == row4 {
			if column > columnA && game.board[row][column-1] != nil && game.board[row][column-1].Type() == PAWN && game.lastMove != nil && game.lastMove.From[0] == row-2 && game.lastMove.To[0] == row && game.lastMove.To[1] == column-1 {
				moves = append(moves, Move{From: from, To: &Square{row - 1, column - 1}})
			}
			if column < columnH && game.board[row][column+1] != nil && game.board[row][column+1].Type() == PAWN && game.lastMove != nil && game.lastMove.From[0] == row-2 && game.lastMove.To[0] == row && game.lastMove.To[1] == column+1 {
				moves = append(moves, Move{From: from, To: &Square{row - 1, column + 1}})
			}
		}
		if column > columnA && game.board[row-1][column-1] != nil && game.board[row-1][column-1].Color() == White {
			if row == row2 {
				for i := KNIGHT; i <= QUEEN; i++ {
					moves = append(moves, Move{From: from, To: &Square{row - 1, column - 1}, PromotionType: &i})
				}
			} else {
				moves = append(moves, Move{From: from, To: &Square{row - 1, column - 1}})
			}
		}
		if column < columnH && game.board[row-1][column+1] != nil && game.board[row-1][column+1].Color() == White {
			if row == row2 {
				for i := KNIGHT; i <= QUEEN; i++ {
					moves = append(moves, Move{From: from, To: &Square{row - 1, column + 1}, PromotionType: &i})
				}
			} else {
				moves = append(moves, Move{From: from, To: &Square{row - 1, column + 1}})
			}
		}

	}
	moves = filterMovesThatLoseKing(game, moves, p.color)
	return moves
}
func (p Pawn) Attacks(game *Game, from *Square) []Square {
	attacks := make([]Square, 0)
	if from == nil {
		return attacks
	}
	row, column := from[0], from[1]
	if p.color == White {
		if column > columnA {
			attacks = append(attacks, Square{row + 1, column - 1})
		}
		if column < columnH {
			attacks = append(attacks, Square{row + 1, column + 1})
		}
	}
	if p.color == Black {
		if column > columnA {
			attacks = append(attacks, Square{row - 1, column - 1})
		}
		if column < columnH {
			attacks = append(attacks, Square{row - 1, column + 1})
		}
	}
	return attacks
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
func (k Knight) Attacks(game *Game, from *Square) []Square {
	attacks := make([]Square, 0)
	if from == nil {
		return attacks
	}
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
			attacks = append(attacks, square)
		}
	}
	return attacks

}
func (k Knight) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	attacks := k.Attacks(game, from)
	for _, square := range attacks {
		if game.board[square[0]][square[1]] == nil || game.board[square[0]][square[1]].Color() != k.color {
			moves = append(moves, Move{From: from, To: &square})
		}
	}
	moves = filterMovesThatLoseKing(game, moves, k.color)
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

func (b Bishop) Attacks(game *Game, from *Square) []Square {
	attacks := make([]Square, 0)
	if from == nil {
		return attacks
	}
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		for row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			attacks = append(attacks, Square{destRow, destColumn})
			if game.board[destRow][destColumn] != nil {
				break
			}
			destRow += direction[0]
			destColumn += direction[1]
		}
	}
	return attacks
}
func (b Bishop) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	attacks := b.Attacks(game, from)
	for _, square := range attacks {
		if game.board[square[0]][square[1]] == nil || game.board[square[0]][square[1]].Color() != b.color {
			moves = append(moves, Move{From: from, To: &square})
		}
	}
	moves = filterMovesThatLoseKing(game, moves, b.color)
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
func (r Rook) Attacks(game *Game, from *Square) []Square {
	attacks := make([]Square, 0)
	if from == nil {
		return attacks
	}
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		for row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			attacks = append(attacks, Square{destRow, destColumn})
			if game.board[destRow][destColumn] != nil {
				break
			}
			destRow += direction[0]
			destColumn += direction[1]
		}
	}
	return attacks
}
func (r Rook) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	attacks := r.Attacks(game, from)
	for _, square := range attacks {
		if game.board[square[0]][square[1]] == nil || game.board[square[0]][square[1]].Color() != r.color {
			moves = append(moves, Move{From: from, To: &square})
		}
	}
	filterMovesThatLoseKing(game, moves, r.color)
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

func (q Queen) Attacks(game *Game, from *Square) []Square {
	attacks := make([]Square, 0)
	if from == nil {
		return attacks
	}
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		for row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			attacks = append(attacks, Square{destRow, destColumn})
			if game.board[destRow][destColumn] != nil {
				break
			}
			destRow += direction[0]
			destColumn += direction[1]
		}
	}
	return attacks
}
func (q Queen) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	attacks := q.Attacks(game, from)
	for _, square := range attacks {
		if game.board[square[0]][square[1]] == nil || game.board[square[0]][square[1]].Color() != q.color {
			moves = append(moves, Move{From: from, To: &square})
		}
	}
	moves = filterMovesThatLoseKing(game, moves, q.color)
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
func (k King) Attacks(game *Game, from *Square) []Square {
	attacks := make([]Square, 0)
	if from == nil {
		return attacks
	}
	row, column := from[0], from[1]
	for _, direction := range [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		destRow, destColumn := row+direction[0], column+direction[1]
		if row1 <= destRow && destRow <= row8 && columnA <= destColumn && destColumn <= columnH {
			attacks = append(attacks, Square{destRow, destColumn})
		}
	}
	return attacks
}
func (k King) LegalMoves(game *Game, from *Square) []Move {
	moves := make([]Move, 0)
	if from == nil {
		return moves
	}
	attacks := k.Attacks(game, from)
	attacked := game.SquaresAttacked(1 - k.color)
	for _, square := range attacks {
		if game.board[square[0]][square[1]] == nil || game.board[square[0]][square[1]].Color() != k.color && !attacked[Square{square[0], square[1]}.ToString()] {
			moves = append(moves, Move{From: from, To: &square})
		}
	}

	row := from[0]

	if game.CanShortCastle(k.color) {
		if game.board[row][columnB] == nil && game.board[row][columnC] == nil && game.board[row][columnD] == nil && !attacked[Square{row, columnC}.ToString()] && !attacked[Square{row, columnD}.ToString()] && !attacked[Square{row, columnE}.ToString()] {
			moves = append(moves, Move{From: from, To: &Square{row, columnC}})
		}
	}
	if game.CanLongCastle(k.color) {
		if game.board[row][columnG] == nil && game.board[row][columnF] == nil && !attacked[Square{row, columnF}.ToString()] && !attacked[Square{row, columnG}.ToString()] && !attacked[Square{row, columnE}.ToString()] {
			moves = append(moves, Move{From: from, To: &Square{row, columnG}})
		}
	}
	moves = filterMovesThatLoseKing(game, moves, k.color)
	return moves
}

func filterMovesThatLoseKing(game *Game, moves []Move, color Color) []Move {
	moves = filter(moves, func(move Move) bool {
		g := game.Copy()
		g.playWithoutChecking(&move)
		return !g.IsCheck(color)
	})
	return moves
}
