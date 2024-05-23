package main

type PieceType int

type Color int

const (
	White Color = iota
	Black
)
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
	LegalMoves(board *Board, row, column int) [][2]int
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

func (p Pawn) LegalMoves(board *Board, row, column int) [][2]int {
	//TODO
	return make([][2]int, 0)
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
func (k Knight) LegalMoves(board *Board, row, column int) [][2]int {
	//TODO
	return make([][2]int, 0)
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

func (b Bishop) LegalMoves(board *Board, row, column int) [][2]int {
	//TODO
	return make([][2]int, 0)
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
func (r Rook) LegalMoves(board *Board, row, column int) [][2]int {
	//TODO
	return make([][2]int, 0)
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

func (q Queen) LegalMoves(board *Board, row, column int) [][2]int {
	//TODO
	return make([][2]int, 0)
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
func (k King) LegalMoves(board *Board, row, column int) [][2]int {
	//TODO
	return make([][2]int, 0)
}
