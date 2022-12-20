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

// TODO - this one may be done better
func SetShmupArchCoreSettings(retroarchCfgDirPath string) ([]Message, error) {
	t := timeStamp()
	report := make([]Message, 0, len(GameSettings)+1)

	// retroarch.cfg core settings
	report = append(report, setSettings(GlobalSettings, t, retroarchCfgDirPath, "", RETROARCH_CFG))

	// FBNeo Game Settings
	for gameName, cfgEntries := range GameSettings {
		report = append(report, setSettings(cfgEntries, t, retroarchCfgDirPath, FBNEO_CFG_DIR, gameName+".cfg"))
	}

	return report, nil
}

// CheckRetroarchCfgExists checks the given folder an existing retroarch.cfg
// to determine if the given path is valid
func CheckRetroarchCfgExists(retroArchCfgDirPath string) bool {
	info, err := os.Stat(filepath.Join(retroArchCfgDirPath, RETROARCH_CFG))
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
