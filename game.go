package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	dt "github.com/dylhunn/dragontoothmg"
	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/border"
	"github.com/maaslalani/gambit/pieces"
	"github.com/maaslalani/gambit/position"
)

type model struct {
	board           *dt.Board
	legalMoves      []dt.Move
	legalPieceMoves []dt.Move
	selected        string
}

func Model() tea.Model {
	board := dt.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	legalMoves := board.GenerateLegalMoves()
	return model{
		board:      &board,
		legalMoves: legalMoves,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) View() string {
	var s strings.Builder

	fen := m.board.ToFen()
	rs := strings.Split(fen, " ")[0]
	ranks := strings.Split(rs, "/")

	for r, rank := range ranks {
		if r == board.FirstRow {
			s.WriteString(border.Top())
		}

		count := 0
		for c, cell := range rank {
			if c == board.FirstCol {
				label := fmt.Sprintf(" %d ", board.LastRow-r+1)
				s.WriteString(Faint.Render(label) + border.Vertical)
			}

			if isNumeric(cell) {
				for i := 0; i < runeToInt(cell); i++ {
					display := "   "
					// Loop through all piece legal moves and see if this square matches any
					for _, move := range m.legalPieceMoves {
						if strings.HasSuffix(move.String(), position.ToSquare(board.LastRow-r, count)) {
							display = Cyan.Render(" . ")
							break
						}
					}
					s.WriteString(display + border.Vertical)
					count += 1
				}
			} else {
				var style lipgloss.Style
				if m.selected == position.ToSquare(board.LastRow-r, count) {
					style = Selected
				} else {
					style = lipgloss.NewStyle()
					for _, move := range m.legalPieceMoves {
						if strings.HasSuffix(move.String(), position.ToSquare(board.LastRow-r, count)) {
							style = Red
							break
						}
					}
				}
				s.WriteString(" " + style.Render(pieces.Display[string(cell)]) + " " + border.Vertical)
				count += 1
			}
		}
		s.WriteRune('\n')

		if r != board.LastRow {
			s.WriteString(border.Middle())
		} else {
			s.WriteString(border.Bottom() + Faint.Render(border.BottomLabels()))
		}
	}
	s.WriteRune('\n')

	return s.String()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}

		col, row := border.Cell(msg.X, msg.Y)

		if m.selected != "" {
			from := m.selected
			to := fmt.Sprintf("%s", position.ToSquare(row, col))

			for _, move := range m.legalMoves {
				if move.String() == from+to {
					// Perform move
					m.board.Apply(move)
					m.legalMoves = m.board.GenerateLegalMoves()

					// Unselect piece
					m.selected = ""
					m.legalPieceMoves = []dt.Move{}

					return m, nil
				}
			}

			// We didn't encounter a valid move
			m.selected = to
			m.legalPieceMoves = []dt.Move{}
			for _, move := range m.legalMoves {
				if strings.HasPrefix(move.String(), m.selected) {
					m.legalPieceMoves = append(m.legalPieceMoves, move)
				}
			}
		} else {
			m.selected = fmt.Sprintf("%s", position.ToSquare(row, col))
			m.legalPieceMoves = []dt.Move{}
			for _, move := range m.legalMoves {
				if strings.HasPrefix(move.String(), m.selected) {
					m.legalPieceMoves = append(m.legalPieceMoves, move)
				}
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// isNumeric returns true if the current rune is a number
func isNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

// runeToInt converts a rune to an int
func runeToInt(r rune) int {
	return int(r - '0')
}
