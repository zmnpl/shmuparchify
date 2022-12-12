package bubl

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zmnpl/shmuparchify/core"
)

type cfgDirContainsCfgMsg bool

func makeCheckDirCommand(path string) tea.Cmd {
	return func() tea.Msg {
		//text := "Doesn't look like a RetroArch config dir... Are you sure?"
		cfgExists := core.CheckRetroarchCfgExists(path)
		return cfgDirContainsCfgMsg(cfgExists)
	}
}

type doneWithSettingsMsg struct {
	report []core.Message
	err    error
}

func makeDoCoreSettingsCommand(path string) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		report, err := core.SetShmupArchCoreSettings(path)
		return doneWithSettingsMsg{report: report, err: err}
	}
}
