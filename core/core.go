package core

import (
	"fmt"
	"io/ioutil"
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

type Job func() Message

// RetroArch Transformer
type RetroArchChanger struct {
	retroarchCfgDirPath string
	romPath             string
	timeStamp           string
	withOverlays        bool
}

// WithOverlays sets the option to also download overlays
func WithOverlays(r RetroArchChanger) RetroArchChanger {
	r.withOverlays = true
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

func (r RetroArchChanger) GetShmupArchJobs() []Job {
	jobs := make([]Job, 0, len(ShmupArchGameSettings)+1)
	// retroarch.cfg core settings
	jobs = append(jobs, func() Message {
		err := r.setSettings(GlobalSettings, "", RETROARCH_CFG, true)
		if err != nil {
			return Message{Success: false, Text: fmt.Sprintf("retroarch.cfg; Could not apply global settings: %v", err)}
		}
		return Message{Success: true, Text: "retroarch.cfg; Applied global settings"}
	})

	// FBNeo Shmup Settings
	for g := range ShmupArchGameSettings {
		// need to copy into new variables to use in closure (otherwise pointer to loop var is used)
		// details see: https://github.com/golang/go/wiki/CommonMistakes
		gameCfg := g + ".cfg"
		settings := ShmupArchGameSettings[g]
		// create job
		j := func() Message {
			err := r.setSettings(settings, FBNEO_CFG_DIR, gameCfg, true)
			if err != nil {
				return Message{Success: false, Text: fmt.Sprintf("%s; (%s) Could not apply settings: %v", gameCfg, FBNEO_CFG_DIR, err)}
			}
			return Message{Success: true, Text: fmt.Sprintf("%s; (%s) Applied settings", gameCfg, FBNEO_CFG_DIR)}
		}
		jobs = append(jobs, j)
	}

	// FBNeo Non-Shmups Settings
	for g := range GameSettingsNonShmups {
		// need to copy into new variables to use in closure (otherwise pointer to loop var is used)
		// details see: https://github.com/golang/go/wiki/CommonMistakes
		gameCfg := g + ".cfg"
		settings := GameSettingsNonShmups[g]
		// create job
		j := func() Message {
			err := r.setSettings(settings, FBNEO_CFG_DIR, gameCfg, true)
			if err != nil {
				return Message{Success: false, Text: fmt.Sprintf("%s; (%s) Could not apply settings: %v", gameCfg, FBNEO_CFG_DIR, err)}
			}
			return Message{Success: true, Text: fmt.Sprintf("%s; (%s) Applied settings", gameCfg, FBNEO_CFG_DIR)}
		}
		jobs = append(jobs, j)
	}

	return jobs
}

func (r RetroArchChanger) GetOverlayJobs() []Job {
	jobs := make([]Job, 0, len(ShmupArchGameSettings))

	// Only for Shmups as overlays don't really distract here
	for g := range ShmupArchGameSettings {
		// need to copy into new variables to use in closure (otherwise pointer to loop var is used)
		// details see: https://github.com/golang/go/wiki/CommonMistakes
		game := g

		// TODO: Actually, since this is now not a closure but a maker function, explicit copy above is not necessary
		jobs = append(jobs, r.MakeOverlayJob(game))
	}

	return jobs
}

func (r RetroArchChanger) DownloadOverlay(game string) error {
	os.MkdirAll(filepath.Join(r.retroarchCfgDirPath, OVERLAY_PATH), 0755)
	err := downloadFile(fmt.Sprintf(OVERLAY_DOWNLOAD_URL, game+".cfg"), filepath.Join(r.retroarchCfgDirPath, OVERLAY_PATH, game+".cfg"))
	if err != nil {
		return err
	}
	err = downloadFile(fmt.Sprintf(OVERLAY_DOWNLOAD_URL, game+".png"), filepath.Join(r.retroarchCfgDirPath, OVERLAY_PATH, game+".png"))
	if err != nil {
		return err
	}

	return nil
}

// TryFindRetroarchCFGDir gives a d
func TryFindRetroarchCFGDir() string {
	home, _ := homedir.Dir()

	os := runtime.GOOS
	switch os {
	case "windows":
		return "C:\\RetroArch-Win64"
	case "darwin":
		return filepath.Join(home, ".config/retroarch")
	case "linux":
		return filepath.Join(home, ".config/retroarch")
	default:
		return ""
	}
}

// CheckPathExists just checks if the given folder exists in the file system
func (r RetroArchChanger) CheckPathExists() bool {
	_, err := os.Stat(r.retroarchCfgDirPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckRetroarchCfgExists checks the given folder an existing retroarch.cfg
// to determine if the given path is valid
func (r RetroArchChanger) CheckRetroarchCfgExists() bool {
	info, err := os.Stat(filepath.Join(r.retroarchCfgDirPath, RETROARCH_CFG))
	if os.IsNotExist(err) {
		//err = r.CopyDefaultCfg()
		//if err != nil {
		//	return false
		//}
		//return true
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

// CopyDefaultCfg tries a simple copy of the retroarch default config
func (r RetroArchChanger) CopyDefaultCfg() error {
	_, err := os.Stat(filepath.Join(r.retroarchCfgDirPath, RETROARCH_CFG))
	if os.IsNotExist(err) {
		// config does not exist, but maybe default config in same dir (windows zip budle comes like that)
		// try very naive copy of default cfg to make the dir ready
		_, err := os.Stat(filepath.Join(r.retroarchCfgDirPath, RETROARCH_DEFAULT_CFG))
		if err != nil {
			return err
		}
		defaultCfg, err := ioutil.ReadFile(filepath.Join(r.retroarchCfgDirPath, RETROARCH_DEFAULT_CFG))
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(r.retroarchCfgDirPath, RETROARCH_CFG), defaultCfg, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
