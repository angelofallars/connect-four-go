package main

import (
	"fmt"
	"os"

	B "connect-four-go/board"

	tea "github.com/charmbracelet/bubbletea"

	lipgloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor        int
	gameBoard     B.GameBoard
	currentPlayer B.Player
}

func initialModel() model {
	return model{
		gameBoard:     B.NewGameBoard(),
		cursor:        0,
		currentPlayer: B.PLAYER_ONE,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}

		case "right", "l":
			if m.cursor < B.GRID_WIDTH-1 {
				m.cursor++
			}

		case "enter", "j":
			err := m.gameBoard.DropDisc(uint8(m.cursor), m.currentPlayer)

			if err == B.ErrFilledBoard {
				break
			}

			switch m.currentPlayer {
			case B.PLAYER_ONE:
				m.currentPlayer = B.PLAYER_TWO
			case B.PLAYER_TWO:
				m.currentPlayer = B.PLAYER_ONE
			}
		}
	}

	return m, nil
}

var colorPlayerOne = "#E8E373"
var colorPlayerTwo = "#E87A73"
var colorWhite = "#BBBBBB"
var colorBlack = "#333333"

var styleTopBar = lipgloss.NewStyle().
	Bold(true).
	Padding(0, 1).
	MarginBottom(1).
	MarginLeft(10).
	Foreground(lipgloss.Color(colorBlack))

var stylePlayerOne = styleTopBar.Copy().
	Background(lipgloss.Color(colorPlayerOne))

var stylePlayerTwo = styleTopBar.Copy().
	Background(lipgloss.Color(colorPlayerTwo))

var styleCell = lipgloss.NewStyle().
	Background(lipgloss.Color(colorWhite))

var styleCellPlayerOne = styleCell.Copy().
	Foreground(lipgloss.Color(colorPlayerOne))

var styleCellPlayerTwo = styleCell.Copy().
	Foreground(lipgloss.Color(colorPlayerTwo))

var styleCellBorder = styleCell.Copy().
	Foreground(lipgloss.Color(colorBlack))

var styleBoard = styleCell.Copy().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderBackground(lipgloss.Color(colorWhite)).
	BorderForeground(lipgloss.Color(colorBlack))

func (m model) View() string {
	s := ""

	switch m.currentPlayer {
	case B.PLAYER_ONE:
		s += stylePlayerOne.Render("PLAYER ONE TURN")
	case B.PLAYER_TWO:
		s += stylePlayerTwo.Render("PLAYER TWO TURN")
	}

	s += "\n"

	s += "  "
	for i := 0; i < m.cursor; i++ {
		s += "     "
	}
	s += "🔽🔽  "
	s += "\n"

	board := ""
	for i := 0; i < B.GRID_HEIGHT; i++ {
		for k := 0; k < 3; k++ {
			for j := 0; j < B.GRID_WIDTH; j++ {
				board += styleCellBorder.Render("┃")
				if k != 2 {
					switch m.gameBoard.Grid[i][j] {
					case B.BLANK:
						board += styleCell.Render("    ")
					case B.YELLOW:
						board += styleCellPlayerOne.Render("████")
					case B.RED:
						board += styleCellPlayerTwo.Render("████")
					}
				} else {
					board += styleCellBorder.Render("━━━━")
				}
			}

			if k != 2 {
				board += "\n"
			}
		}

		if i < B.GRID_HEIGHT-1 {
			board += "\n"
		}
	}

	board = styleBoard.Render(board)

	s += board

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
