package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

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

type Job func() Message

// RetroArch Transformer
type RetroArchChanger struct {
	retroarchCfgDirPath string
	romPath             string
	timeStamp           string
	withBezels          bool
}

// WithBezels sets the option to also download bezels
func WithBezels(r RetroArchChanger) RetroArchChanger {
	r.withBezels = true
	return r
}

func NewRATransformer(retroarchCfgDirPath string, options ...func(RetroArchChanger) RetroArchChanger) RetroArchChanger {
	r := RetroArchChanger{
		retroarchCfgDirPath: retroarchCfgDirPath,
		timeStamp:           timeStamp(),
	}

	for _, opt := range options {
		r = opt(r)
	}

	return r
}

// TODO - this one may be done better
func (r RetroArchChanger) SetShmupArchCoreSettings() ([]Message, error) {
	report := make([]Message, 0, len(GameSettings)+1)

	// retroarch.cfg core settings
	report = append(report, r.setSettings(GlobalSettings, "", RETROARCH_CFG, true))

	// FBNeo Game Settings
	for gameName, cfgEntries := range GameSettings {
		report = append(report, r.setSettings(cfgEntries, FBNEO_CFG_DIR, gameName+".cfg", true))
	}

	return report, nil
}

func (r RetroArchChanger) GetShmupArchJobs() []Job {
	jobs := make([]Job, 0, len(GameSettings)+1)
	// retroarch.cfg core settings
	jobs = append(jobs, func() Message {
		return r.setSettings(GlobalSettings, "", RETROARCH_CFG, true)
	})

	// FBNeo Game Settings
	for g := range GameSettings {
		// need to copy into new variables to use in closure (otherwise pointer to loop var is used)
		// details see: https://github.com/golang/go/wiki/CommonMistakes
		gameCfg := g + ".cfg"
		settings := GameSettings[g]
		// create job
		j := func() Message { return r.setSettings(settings, FBNEO_CFG_DIR, gameCfg, true) }
		jobs = append(jobs, j)
	}

	return jobs
}

func (r RetroArchChanger) GetBezelJobs() []Job {
	jobs := make([]Job, 0, len(GameSettings))

	for g := range GameSettings {
		// need to copy into new variables to use in closure (otherwise pointer to loop var is used)
		// details see: https://github.com/golang/go/wiki/CommonMistakes
		game := g

		job := func() Message {
			// TODO: add bezel config here as well

			err := r.DownloadBezel(game)
			if err != nil {
				return Message{Success: false, Text: fmt.Sprintf("%s; Failed to download bezel: %v", game, err)}
			}
			return Message{Success: true, Text: fmt.Sprintf("%s; Bezel download successful", game)}
		}
		jobs = append(jobs, job)
	}

	return jobs
}

func (r RetroArchChanger) DownloadBezel(game string) error {
	time.Sleep(250 * time.Millisecond)
	return nil
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

// CheckRetroarchCfgExists checks the given folder an existing retroarch.cfg
// to determine if the given path is valid
func (r RetroArchChanger) CheckRetroarchCfgExists() bool {
	info, err := os.Stat(filepath.Join(r.retroarchCfgDirPath, RETROARCH_CFG))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CheckPermissions returns true if a test write/deletion succeeded
func (r RetroArchChanger) CheckPermissions() bool {
	testFile := filepath.Join(r.retroarchCfgDirPath, "shmuparchify_rw.test")

	err := os.WriteFile(testFile, []byte("testing read/write"), 0755)
	if err != nil {
		return false
	}

	err = os.Remove(testFile)
	if err != nil {
		return false
	}

	return true
}
