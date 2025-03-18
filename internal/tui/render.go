package tui

import "github.com/charmbracelet/lipgloss"

func renderBox(content string) string {
	var boxStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("220"))

	var headerStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("220")).
		Foreground(lipgloss.Color("214")).
		Bold(true).
		Align(lipgloss.Center).
		PaddingLeft(2).
		PaddingRight(2)

	var header = headerStyle.Render("🌻 JSONFlower - JSON Beautifier Tool 🌻")

	var contentStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		PaddingRight(2)

	var formattedContent = header + "\n" +
		contentStyle.Render(content)

	return boxStyle.Render(formattedContent)
}
