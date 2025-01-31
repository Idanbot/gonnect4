package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"unicode"
)

const (
	Players = 2
	Rows    = 7
	Columns = 7
)

type Game struct {
	Board   [Rows][Columns]string
	Players [2]string
	Turn    int
}

func GameLoop() {
	g := NewGame()
	for {
		fmt.Println("Turn: Player", g.GetTurnSymbol())
		g.PrintBoard()
		g.Play()
		g.SwitchTurn()
		fmt.Println()
		if g.CheckWinner() {
			g = NewGame()
		}
		fmt.Println()
	}
}

func NewGame() *Game {
	g := &Game{Players: [2]string{"X", "O"}, Turn: getRandomTurn()}
	for i := range g.Board {
		for j := range g.Board[i] {
			g.Board[i][j] = "."
		}
	}

	g.Board[6][0] = "₁"
	g.Board[6][1] = "₂"
	g.Board[6][2] = "₃"
	g.Board[6][3] = "₄"
	g.Board[6][4] = "₅"
	g.Board[6][5] = "₆"
	g.Board[6][6] = "₇"
	return g
}

func (g *Game) PrintBoard() {
	for _, row := range g.Board {
		for _, cell := range row {
			fmt.Print(cell + " ")
		}
		fmt.Println()
	}
}

func (g *Game) Play() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a column number to play (Q to Exit):")

	r, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading input. Try again.")
		return
	}

	if r == 'Q' || r == 'q' {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}

	column, err := strconv.Atoi(string(r))
	if err != nil || column < 1 || column > Columns || !unicode.IsDigit(r) {
		fmt.Println("Invalid input, please enter a number between 1-7.")
		return
	}
	column -= 1
	for i := Rows - 2; i >= 0; i-- {
		if g.Board[i][column] == "." {
			g.Board[i][column] = g.GetTurnSymbol()
			break
		}
	}
}

func (g *Game) CheckWinner() bool {
	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			if g.Board[i][j] != "." {
				player := g.Board[i][j]
				if g.checkDirection(i, j, 1, 0, player) ||
					g.checkDirection(i, j, 0, 1, player) ||
					g.checkDirection(i, j, 1, 1, player) ||
					g.checkDirection(i, j, 1, -1, player) {
					fmt.Println("Player", player, "wins!")
					fmt.Println("Press R to restart or any other key to exit.")
					reader := bufio.NewReader(os.Stdin)
					r, _, _ := reader.ReadRune()
					if r == 'R' || r == 'r' {
						return true
					}

					fmt.Println("Goodbye!")
					os.Exit(0)
				}
			}
		}
	}

	return false
}

func (g *Game) checkDirection(row, col, dRow, dCol int, player string) bool {
	for k := 1; k < 4; k++ {
		r, c := row+(dRow*k), col+(dCol*k)
		if r < 0 || r >= Rows || c < 0 || c >= Columns || g.Board[r][c] != player {
			return false
		}
	}
	return true
}

func getRandomTurn() int {
	return rand.Intn(Players)
}

func (g *Game) GetTurnSymbol() string {
	return g.Players[g.Turn]
}

func (g *Game) SwitchTurn() {
	g.Turn = (g.Turn + 1) % Players
}
