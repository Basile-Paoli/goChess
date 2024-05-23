package main

import "fmt"

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

// Square represents a square on the chess board where the first element is the row and the second element is the column.
type Square [2]int

func (s Square) toString() string {
	return string('a'+s[1]) + string('1'+s[0])
}
func newSquare(name string) *Square {
	if len(name) != 2 {
		return nil
	}
	if !('a' <= name[0] && name[0] <= 'h' && '1' <= name[1] && name[1] <= '8') {
		return nil
	}
	square := new(Square)
	square[0] = int(name[1] - '1')
	square[1] = int(name[0] - 'a')
	return square

}

type Color int

const (
	White Color = iota
	Black
)

type Move struct {
	From *Square
	To   *Square
}

func (m Move) toString() string {
	return m.From.toString() + m.To.toString()
}

func newMove(start string, end string) *Move {
	move := new(Move)
	move.From = newSquare(start)
	move.To = newSquare(end)
	if move.From == nil || move.To == nil {
		return nil
	}
	return move
}

type Game struct {
	board        [8][8]Piece
	toPlay       Color
	castleRights [2][2]bool
}

func (g *Game) canShortCastle(color Color) bool {
	return g.castleRights[color][0]
}
func (g *Game) canLongCastle(color Color) bool {
	return g.castleRights[color][1]
}

func (g *Game) isLegal(move *Move) bool {
	if move == nil {
		return false
	}
	piece := g.board[move.From[0]][move.From[1]]
	if piece == nil || piece.Color() != g.toPlay {
		return false
	}
	for _, legalMove := range piece.LegalMoves(g, move.From) {
		if legalMove.toString() == move.toString() {
			return true
		}
	}
	return false
}
func (g *Game) switchPlayer() {
	g.toPlay = 1 - g.toPlay
}

func (g *Game) play(move *Move) *Game {
	if move == nil {
		return g
	}
	if !g.isLegal(move) {
		return g
	}
	if g.board[move.From[0]][move.From[1]].Type() == KING {
		g.castleRights[g.toPlay][0] = false
		g.castleRights[g.toPlay][1] = false
		if move.From[1]-move.To[1] == 2 {
			g.board[move.From[0]][columnD] = g.board[move.From[0]][columnA]
			g.board[move.From[0]][columnA] = nil

		} else if move.From[1]-move.To[1] == -2 {
			g.board[move.From[0]][columnF] = g.board[move.From[0]][columnH]
			g.board[move.From[0]][columnH] = nil

		}
	}
	if move.From[1] == columnA && move.From[0] == row1 {
		g.castleRights[White][1] = false
	}
	if move.From[1] == columnH && move.From[0] == row1 {
		g.castleRights[White][0] = false
	}
	if move.From[1] == columnA && move.From[0] == row8 {
		g.castleRights[Black][1] = false
	}
	if move.From[1] == columnH && move.From[0] == row8 {
		g.castleRights[Black][0] = false
	}

	g.board[move.To[0]][move.To[1]] = g.board[move.From[0]][move.From[1]]
	g.board[move.From[0]][move.From[1]] = nil
	g.switchPlayer()
	return g
}
func (g *Game) move(moveStr string) *Game {
	if len(moveStr) != 4 {
		return g
	}
	return g.play(newMove(moveStr[:2], moveStr[2:]))
}
func (g *Game) legalMovesFrom(square *Square) []Move {
	piece := g.board[square[0]][square[1]]
	if piece == nil || piece.Color() != g.toPlay {
		return make([]Move, 0)
	}
	return piece.LegalMoves(g, square)
}

func (g *Game) legalMoves() []Move {
	moves := make([]Move, 0)
	for row := row1; row <= row8; row++ {
		for column := columnA; column <= columnH; column++ {
			square := &Square{row, column}
			moves = append(moves, g.legalMovesFrom(square)...)
		}
	}
	return moves
}

func (g *Game) printBoard() {
	for row := row8; row >= row1; row-- {
		for column := columnA; column <= columnH; column++ {
			piece := g.board[row][column]
			if piece == nil {
				print(" ")
			} else {
				print(piece.Symbol())
			}
			print(" ")
		}
		println()
	}
}

func newGame() *Game {
	var game = new(Game)
	game.board[row1][columnA] = &Rook{White}
	game.board[row1][columnB] = &Knight{White}
	game.board[row1][columnC] = &Bishop{White}
	game.board[row1][columnD] = &Queen{White}
	game.board[row1][columnE] = &King{White}
	game.board[row1][columnF] = &Bishop{White}
	game.board[row1][columnG] = &Knight{White}
	game.board[row1][columnH] = &Rook{White}
	for column := columnA; column <= columnH; column++ {
		game.board[row2][column] = &Pawn{White}
	}
	for column := columnA; column <= columnH; column++ {
		game.board[row7][column] = &Pawn{Black}
	}
	game.board[row8][columnA] = &Rook{Black}
	game.board[row8][columnB] = &Knight{Black}
	game.board[row8][columnC] = &Bishop{Black}
	game.board[row8][columnD] = &Queen{Black}
	game.board[row8][columnE] = &King{Black}
	game.board[row8][columnF] = &Bishop{Black}
	game.board[row8][columnG] = &Knight{Black}
	game.board[row8][columnH] = &Rook{Black}
	game.toPlay = White
	game.castleRights = [2][2]bool{{true, true}, {true, true}}
	return game
}

func main() {
	g := newGame()
	for {
		g.printBoard()
		println()
		for _, move := range g.legalMoves() {
			print(move.toString(), " ")
		}
		println()
		println("Enter your move:")
		var move string
		_, _ = fmt.Scanln(&move)
		g.move(move)
	}
}
