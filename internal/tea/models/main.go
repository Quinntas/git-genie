package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/quinntas/git-genie/internal/tea/bindings"
	"github.com/quinntas/git-genie/internal/utils/structs"
)

type MainModel struct {
	loaded   bool
	quitting bool
	structs.Dimensions

	help help.Model
}

func NewMainModel() *MainModel {
	helpModel := help.New()
	helpModel.ShowAll = true
	return &MainModel{
		loaded:     false,
		quitting:   false,
		Dimensions: structs.NewDimensions(0, 0),
		help:       helpModel,
	}
}

func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	}

	var cmd tea.Cmd
	m.help, cmd = m.help.Update(msg)

	return m, cmd
}

func (m *MainModel) View() string {
	if !m.loaded {
		return "Loading..."
	}

	if m.quitting {
		return "Quitting..."
	}

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, "")
}
