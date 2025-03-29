package tui

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/livensmi1e/jsonflower/internal/file"
)

type state int
type mode int

const (
	SelectingMode state = iota
	SelectingFile
	DisplayingJSON
	DisplayingYAML
)

const (
	ModeNone mode = iota
	ModeBeautify
	ModeMinify
	ModeJSONToYAML
)

type model struct {
	state        state
	mode         mode
	filePicker   filepicker.Model
	list         list.Model
	selectedFile string
	result       string
	quitting     bool
	// err          error
}

type listItem struct {
	title string
	mode  mode
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return "" }
func (i listItem) FilterValue() string { return i.title }

var (
	selectedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	enumeratorIcon = "> "
)

type customDelegate struct {
	list.DefaultDelegate
}

func (d customDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(listItem)
	if !ok {
		return
	}
	itemStyle := d.Styles.NormalTitle
	prefix := "  "
	if index == m.Index() {
		itemStyle = selectedStyle
		prefix = selectedStyle.Render(enumeratorIcon)
	}
	fmt.Fprint(w, prefix+itemStyle.Render(i.Title()))
}

func InitModel() model {
	items := []list.Item{
		listItem{"Beautify JSON", ModeBeautify},
		listItem{"Minify JSON", ModeMinify},
		listItem{"Convert JSON → YAML", ModeJSONToYAML},
	}
	delegate := customDelegate{}
	l := list.New(items, delegate, 40, 7)
	l.Title = "Select your choice:"
	l.Styles.Title = lipgloss.NewStyle()
	l.SetShowStatusBar(false)
	l.SetShowPagination(false)
	l.SetFilteringEnabled(false)
	fp := filepicker.New()
	dir, _ := filepath.Abs(".")
	fp.Height = file.GetFileCount(dir)
	fp.CurrentDirectory = dir
	m := model{state: SelectingMode, list: l, filePicker: fp}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "esc", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.state == SelectingMode {
				item, ok := m.list.SelectedItem().(listItem)
				if ok {
					m.mode = item.mode
					m.state = SelectingFile
					return m, m.filePicker.Init()
				}
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		// TODO: Fix to resize issue
	}
	var cmd tea.Cmd
	switch m.state {
	case SelectingMode:
		m.list, cmd = m.list.Update(msg)
	case SelectingFile:
		m.filePicker, cmd = m.filePicker.Update(msg)
		if didSelect, path := m.filePicker.DidSelectFile(msg); didSelect {
			m.selectedFile = path
			content, err := os.ReadFile(path)
			if err != nil {
				m.result = "❌ Error reading file"
			}
			switch m.mode {
			case ModeBeautify:
				m.result = beautifyJSON(string(content))
				m.state = DisplayingJSON
			case ModeJSONToYAML:
				m.result = convertJSON2YAML(string(content))
				m.state = DisplayingYAML
			}
		}
	}
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	switch m.state {
	case SelectingMode:
		return renderBox(m.list.View())
	case SelectingFile:
		return renderBox("  Pick your JSON file: \n\n" + m.filePicker.View())
	case DisplayingJSON, DisplayingYAML:
		return renderBox(m.result)
	}
	return ""
}
