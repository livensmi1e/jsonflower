package beautify

import "github.com/charmbracelet/lipgloss"

var (
	keyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("214")) // Vàng
	strStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("35"))  // Xanh lá
	numStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("141")) // Tím
	boolStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Đỏ
	nullStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244")) // Xám
)
