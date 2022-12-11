package bubl

import "github.com/charmbracelet/lipgloss"

const (
	// TERMINAL THEME COLORS
	BLACK   = lipgloss.Color("0")
	RED     = lipgloss.Color("1")
	GREEN   = lipgloss.Color("2")
	YELLOW  = lipgloss.Color("3")
	BLUE    = lipgloss.Color("4")
	MAGENTA = lipgloss.Color("5")
	CYAN    = lipgloss.Color("6")
	WHITE   = lipgloss.Color("7")

	LIGHT_BLACK   = lipgloss.Color("8")
	LIGHT_RED     = lipgloss.Color("9")
	LIGTH_GREEN   = lipgloss.Color("10")
	LIGHT_YELLOW  = lipgloss.Color("11")
	LIGHT_BLUE    = lipgloss.Color("12")
	LIGHT_MAGENTA = lipgloss.Color("13")
	LIGHT_CYAN    = lipgloss.Color("14")

	OUTER_MARGIN = 1
)

var (
	docStyle           lipgloss.Style
	goodTextStyle      lipgloss.Style
	badTextStyle       lipgloss.Style
	dimmedTextStyle    lipgloss.Style
	indicatorTextStyle lipgloss.Style
)

func init() {
	docStyle = lipgloss.NewStyle().Margin(OUTER_MARGIN)

	goodTextStyle = lipgloss.NewStyle().Foreground(GREEN)
	badTextStyle = lipgloss.NewStyle().Foreground(RED)

	indicatorTextStyle = lipgloss.NewStyle().Foreground(BLUE).PaddingRight(1)

	dimmedTextStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})
}
