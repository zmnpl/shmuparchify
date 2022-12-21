package core

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

const (
	BACKUP_DIR = "_backup"

	MESSAGE_GOOD = "SUCCESS"
	MESSAGE_INFO = "INFO"
	MASSGE_ERROR = "ERROR"
)

type Message struct {
	Success bool
	Text    string
}

type SAIFI struct {
	retroarchCfgDirPath string
	timeStamp           string
}

func NewSAIFI(retroarchCfgDirPath string, options ...func(SAIFI) SAIFI) SAIFI {
	sai := SAIFI{
		retroarchCfgDirPath: retroarchCfgDirPath,
		timeStamp:           timeStamp(),
	}

	for _, opt := range options {
		sai = opt(sai)
	}

	return sai
}

// TODO - this one may be done better
func (saifi SAIFI) SetShmupArchCoreSettings() ([]Message, error) {
	report := make([]Message, 0, len(GameSettings)+1)

	// retroarch.cfg core settings
	report = append(report, saifi.setSettings(GlobalSettings, "", RETROARCH_CFG))

	// FBNeo Game Settings
	for gameName, cfgEntries := range GameSettings {
		report = append(report, saifi.setSettings(cfgEntries, FBNEO_CFG_DIR, gameName+".cfg"))
	}

	return report, nil
}

// CheckRetroarchCfgExists checks the given folder an existing retroarch.cfg
// to determine if the given path is valid
func (saifi SAIFI) CheckRetroarchCfgExists() bool {
	info, err := os.Stat(filepath.Join(saifi.retroarchCfgDirPath, RETROARCH_CFG))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// TryFindRetroarchCFGDir gives a d
func TryFindRetroarchCFGDir() string {
	home, _ := homedir.Dir()

	os := runtime.GOOS
	switch os {
	case "windows":
		return ""
	case "darwin":
		return filepath.Join(home, ".config/retroarch")
	case "linux":
		return filepath.Join(home, ".config/retroarch")
	default:
		return ""
	}
}
