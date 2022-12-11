package bubl

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	STEP_RETROARCH_INPUT = 1
	STEP_CONFIRM         = 2
	STEP_RUNNING         = 3
	STEP_DONE            = 4
)

type model struct {
	step int

	retroarchCfgDirInput textinput.Model
	dirHasCfg            bool

	tmp_running string
	success     bool

	spinner  spinner.Model
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func Run() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	retroarchCfgDirInput := textinput.New()
	//retroarchCfgDirInput.PlaceholderStyle = myStyle.styleInactiveText
	//retroarchCfgDirInput.TextStyle = myStyle.styleActiveText
	retroarchCfgDirInput.Placeholder = "type ..."
	retroarchCfgDirInput.Prompt = ""
	retroarchCfgDirInput.Focus()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		step:                 STEP_RETROARCH_INPUT,
		retroarchCfgDirInput: retroarchCfgDirInput,
		tmp_running:          "Now would go running",

		// Our to-do list is a grocery list
		choices: []string{"ShmupArch Core", "Bezels", "Yoko/Tate", "Button Config"},
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
		spinner:  s,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.step == STEP_RETROARCH_INPUT {
				m.step = STEP_CONFIRM
				cmd = makeCheckDirCommand(m.retroarchCfgDirInput.Value())
				return m, cmd
			}
			if m.step == STEP_CONFIRM {
				m.step = STEP_RUNNING
				cmd := makeDoCoreSettingsCommand(m.retroarchCfgDirInput.Value())
				return m, cmd
			}

		case "esc":
			if m.step == STEP_CONFIRM {
				m.step = STEP_RETROARCH_INPUT
			}

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}

		// update input
		if m.step == STEP_RETROARCH_INPUT {
			m.retroarchCfgDirInput, cmd = m.retroarchCfgDirInput.Update(msg)
		}

	case cfgDirContainsCfgMsg:
		m.dirHasCfg = bool(msg)

	case doneWithSettingsMsg:
		m.step = STEP_DONE
		if doneWithSettingsMsg(msg).err == nil {
			m.success = true
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	}

	return m, cmd
}

func (m model) View() string {
	indicator := indicatorTextStyle.Render(">")

	// The header
	s := "Hey, to shmupify your RetrArch config directory..."
	s += "\n"

	if m.step == STEP_RETROARCH_INPUT {
		s += indicator
	}
	s += m.retroarchCfgDirInput.View()
	s += "\n\n"

	if m.step >= STEP_CONFIRM {
		// print results of sanity checks
		if m.dirHasCfg {
			s += goodTextStyle.Render("OK   ") + "Found existing RetroArch config in the given directory"
		} else {
			s += badTextStyle.Render("WARN ") + "Could not find existing RetroArch config in the given directory"
		}
		s += "\n"

		// assessment
		how := badTextStyle.Render("bad")
		if m.dirHasCfg {
			how = goodTextStyle.Render("good")
		}
		s += fmt.Sprintf("Looks %s...", how)
		s += "\n\n"
	}

	if m.step == STEP_CONFIRM {
		s += indicator
		s += "What now?\n"
		s += indicatorTextStyle.Render("Enter") + " - Let's go; Shmupify my RetroArch\n"
		s += indicatorTextStyle.Render("ESC") + "   - Edit path"
		s += "\n\n"
	}

	if m.step == STEP_RUNNING {
		s += m.spinner.View() + " " + goodTextStyle.Render("I'm on it, hang tight!")
		s += "\n\n"
	}

	if m.step == STEP_DONE {
		s += goodTextStyle.Render("DONE; Now go, shoot'em up.")
		s += "\n\n"
	}

	// s += "What should we buy at the market?\n\n"
	// // Iterate over our choices
	// for i, choice := range m.choices {

	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}

	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}

	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	// }

	// The footer
	s += dimmedTextStyle.Render("Ctrl + C to quit.") + "\n"

	return docStyle.Render(s)
}
