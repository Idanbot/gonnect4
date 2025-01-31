package game

import "fmt"

const (
    Rows    = 6
    Columns = 7
)

type Game struct {
    Board  [Rows][Columns]string
    Player string
}

func NewGame() *Game {
    g := &Game{Player: "X"}
    for i := range g.Board {
        for j := range g.Board[i] {
            g.Board[i][j] = "."
        }
    }
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
