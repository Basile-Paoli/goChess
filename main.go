package main

const (
	row1 = iota
	row2
	row3
	row4
	row5
	row6
	row7
	row8
)

const (
	columnA = iota
	columnB
	columnC
	columnD
	columnE
	columnF
	columnG
	columnH
)

type Tile struct {
	Piece Piece
}

type Board struct {
	tiles [8][8]Tile
}

func (b Board) print() {
	for row := row8; row >= row1; row-- {
		for column := columnA; column <= columnH; column++ {
			if piece := b.tiles[row][column].Piece; piece == nil {
				print(" ")
			} else {
				print(piece.Symbol())
			}
			print(" ")
		}
		println()
	}
}

func newBoard() Board {
	var board Board
	board.tiles[row1][columnA].Piece = &Rook{White}
	board.tiles[row1][columnB].Piece = &Knight{White}
	board.tiles[row1][columnC].Piece = &Bishop{White}
	board.tiles[row1][columnD].Piece = &Queen{White}
	board.tiles[row1][columnE].Piece = &King{White}
	board.tiles[row1][columnF].Piece = &Bishop{White}
	board.tiles[row1][columnG].Piece = &Knight{White}
	board.tiles[row1][columnH].Piece = &Rook{White}
	for column := columnA; column <= columnH; column++ {
		board.tiles[row2][column].Piece = &Pawn{White}
	}
	for column := columnA; column < columnH; column++ {
		board.tiles[row7][column].Piece = &Pawn{Black}
	}
	board.tiles[row8][columnA].Piece = &Rook{Black}
	board.tiles[row8][columnB].Piece = &Knight{Black}
	board.tiles[row8][columnC].Piece = &Bishop{Black}
	board.tiles[row8][columnD].Piece = &Queen{Black}
	board.tiles[row8][columnE].Piece = &King{Black}
	board.tiles[row8][columnF].Piece = &Bishop{Black}
	board.tiles[row8][columnG].Piece = &Knight{Black}
	board.tiles[row8][columnH].Piece = &Rook{Black}
	return board
}

func main() {
	board := newBoard()
	board.print()
}
