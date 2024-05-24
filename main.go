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

func (s Square) ToString() string {
	return string('a'+s[1]) + string('1'+s[0])
}
func NewSquare(name string) *Square {
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

func (m Move) ToString() string {
	return m.From.ToString() + m.To.ToString()
}

func NewMove(start string, end string) *Move {
	move := new(Move)
	move.From = NewSquare(start)
	move.To = NewSquare(end)
	if move.From == nil || move.To == nil {
		return nil
	}
	return move
}

type Game struct {
	board  [8][8]Piece
	toPlay Color
	// castleRights stores the rights to castle for each player. The first index is the color of the player and the second index is the side of the board. The first element is for short castling and the second element is for long castling.
	castleRights [2][2]bool
}

func (g *Game) CanShortCastle(color Color) bool {
	return g.castleRights[color][0]
}
func (g *Game) CanLongCastle(color Color) bool {
	return g.castleRights[color][1]
}

func (g *Game) IsLegal(move *Move) bool {
	if move == nil {
		return false
	}
	piece := g.board[move.From[0]][move.From[1]]
	if piece == nil || piece.Color() != g.toPlay {
		return false
	}
	for _, legalMove := range piece.LegalMoves(g, move.From) {
		if legalMove.ToString() == move.ToString() {
			return true
		}
	}
	return false
}
func (g *Game) switchPlayer() {
	g.toPlay = 1 - g.toPlay
}

func (g *Game) Play(move *Move) *Game {
	if move == nil {
		return g
	}
	if !g.IsLegal(move) {
		return g
	}
	piece := g.board[move.From[0]][move.From[1]]

	g.changeCastlingRights(move, piece)

	if piece.Type() == KING && (move.From[1]-move.To[1] == 2 || move.From[1]-move.To[1] == -2) {
		g.castle(move)
	}

	g.board[move.To[0]][move.To[1]] = piece
	g.board[move.From[0]][move.From[1]] = nil
	g.switchPlayer()
	return g
}
func (g *Game) playWithoutChecking(move *Move) *Game {
	if move == nil {
		return g
	}
	piece := g.board[move.From[0]][move.From[1]]
	g.changeCastlingRights(move, piece)
	if piece.Type() == KING && (move.From[1]-move.To[1] == 2 || move.From[1]-move.To[1] == -2) {
		g.castle(move)
	}
	g.board[move.To[0]][move.To[1]] = piece
	g.board[move.From[0]][move.From[1]] = nil
	g.switchPlayer()
	return g
}

func (g *Game) changeCastlingRights(move *Move, piece Piece) {
	if piece.Type() == KING {
		g.castleRights[g.toPlay][0] = false
		g.castleRights[g.toPlay][1] = false
	} else if piece.Type() == ROOK {
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
	}
}

func (g *Game) castle(move *Move) {
	if move.From[1]-move.To[1] == 2 {
		g.board[move.From[0]][columnD] = g.board[move.From[0]][columnA]
		g.board[move.From[0]][columnA] = nil

	} else if move.From[1]-move.To[1] == -2 {
		g.board[move.From[0]][columnF] = g.board[move.From[0]][columnH]
		g.board[move.From[0]][columnH] = nil

	}
}
func (g *Game) Move(moveStr string) *Game {
	if len(moveStr) != 4 {
		return g
	}
	return g.Play(NewMove(moveStr[:2], moveStr[2:]))
}
func (g *Game) LegalMovesFrom(square *Square) []Move {
	piece := g.board[square[0]][square[1]]
	if piece == nil || piece.Color() != g.toPlay {
		return make([]Move, 0)
	}
	return piece.LegalMoves(g, square)
}

func (g *Game) LegalMoves() []Move {
	moves := make([]Move, 0)
	for row := row1; row <= row8; row++ {
		for column := columnA; column <= columnH; column++ {
			square := &Square{row, column}
			moves = append(moves, g.LegalMovesFrom(square)...)
		}
	}
	return moves
}
func (g *Game) SquaresAttacked(attacker Color) map[string]bool {
	squares := make(map[string]bool)
	for row := row1; row <= row8; row++ {
		for column := columnA; column <= columnH; column++ {
			piece := g.board[row][column]
			if piece != nil && piece.Color() == attacker {
				for _, square := range piece.Attacks(g, &Square{row, column}) {
					squares[square.ToString()] = true

				}
			}
		}
	}
	return squares
}

// IsCheck returns true if the king of the given color is in check.
func (g *Game) IsCheck(color Color) bool {
	attacker := 1 - color
	for row := row1; row <= row8; row++ {
		for column := columnA; column <= columnH; column++ {
			piece := g.board[row][column]
			if piece != nil && piece.Color() == color && piece.Type() == KING {
				kingSquare := &Square{row, column}
				return g.SquaresAttacked(attacker)[kingSquare.ToString()]
			}
		}
	}
	return false
}

func (g *Game) IsCheckmate() bool {
	return g.IsCheck(g.toPlay) && len(g.LegalMoves()) == 0
}
func (g *Game) IsStalemate() bool {
	return !g.IsCheck(g.toPlay) && len(g.LegalMoves()) == 0
}

func (g *Game) PrintBoard() {
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
func (g *Game) Copy() *Game {
	var game = new(Game)
	for row := row1; row <= row8; row++ {
		for column := columnA; column <= columnH; column++ {
			game.board[row][column] = g.board[row][column]
		}
	}
	game.toPlay = g.toPlay
	game.castleRights = g.castleRights
	return game

}

func NewGame() *Game {
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
	g := NewGame()
	for {
		g.PrintBoard()
		println()
		if g.IsCheckmate() {
			println("Checkmate")
			break
		}
		if g.IsStalemate() {
			println("Stalemate")
			break
		}
		for _, move := range g.LegalMoves() {
			print(move.ToString(), " ")
		}
		println()
		var color string
		if g.toPlay == White {
			color = "White"
		} else {
			color = "Black"
		}
		fmt.Printf("%s to Play, enter your Move : \n", color)
		var move string
		fmt.Scanln(&move)
		g.Move(move)
	}
	fmt.Scanln()
}
