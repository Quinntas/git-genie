package models

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/quinntas/git-genie/internal/tea/bindings"
	"github.com/quinntas/git-genie/internal/tea/styles"
	"github.com/quinntas/git-genie/internal/utils/structs"
)

type MainModel struct {
	loaded   bool
	quitting bool

	spinner spinner.Model
	structs.Dimensions

	help help.Model
}

func NewMainModel() *MainModel {
	helpModel := help.New()
	helpModel.ShowAll = true

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &MainModel{
		loaded:     false,
		quitting:   false,
		spinner:    s,
		Dimensions: structs.NewDimensions(0, 0),
		help:       helpModel,
	}
}

func (m *MainModel) Init() tea.Cmd {
	if m.loaded || !m.quitting {
		return nil
	} else {
		return m.spinner.Tick
	}
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Dimensions.Width = msg.Width
		m.Dimensions.Height = msg.Height

		m.help.Width = msg.Width

		m.loaded = true
		return m, tea.Batch()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, bindings.Keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}

	case error:
		return m, nil

	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	m.help, cmd = m.help.Update(msg)

	return m, cmd
}

func (m *MainModel) View() string {
	if !m.loaded {
		str := fmt.Sprintf("\n\n   %s Loading...   \n\n", m.spinner.View())
		return styles.PlaceCenter(m.Width, m.Height, str)
	}

	if m.quitting {
		str := fmt.Sprintf("\n\n   %s Quitting...   \n\n", m.spinner.View())
		return styles.PlaceCenter(m.Width, m.Height, str)
	}

	helpView := m.help.View(bindings.Keys)

	return styles.PlaceCenter(m.Width, m.Height, helpView)
}
