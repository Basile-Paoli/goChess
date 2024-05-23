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
	board  [8][8]Piece
	toPlay Color
}

func (g *Game) isLegal(move *Move) bool {
	if move == nil {
		return false
	}
	piece := g.board[move.From[0]][move.From[1]]
	if piece != nil && piece.Color() == g.toPlay {
		return true
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
	g.board[move.To[0]][move.To[1]] = g.board[move.From[0]][move.From[1]]
	g.board[move.From[0]][move.From[1]] = nil
	g.switchPlayer()
	return g
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
	return game
}

func main() {
	game := newGame()
	game.play(newMove("e2", "e4")).printBoard()
}
