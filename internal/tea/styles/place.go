package styles

import "github.com/charmbracelet/lipgloss"

func PlaceCenter(width int, height int, view string) string {
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, view)
}
