package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const (
	BACKUP_DIR = "_backup"

	MESSAGE_GOOD = "SUCCESS"
	MESSAGE_INFO = "INFO"
	MASSGE_ERROR = "ERROR"
)

type Message struct {
	status string
	text   string
}

var (
	Report []Message
)

// TODO - this one may be done better
func SetShmupArchCoreSettings(retroarchCfgDirPath string) error {
	err := UpdateCfg(filepath.Join(retroarchCfgDirPath, RETROARCH_CFG), GlobalSettings)
	if err != nil {
		return err
	}

	gameErrors := ""
	for gameName, cfgEntries := range GameSettings {
		err = UpdateCfg(filepath.Join(retroarchCfgDirPath, FBNEO_CFG_DIR, gameName+".cfg"), cfgEntries)
		if err != nil {
			gameErrors += err.Error()
		}
	}
	if len(gameErrors) > 0 {
		return fmt.Errorf(gameErrors)
	}

	return nil
}

func UpdateCfg(cfgPath string, entries []cfgEntry) error {
	//fmt.Printf("\nOptimizing config: %s\n", cfgPath)

	cfg, err := readCfg(cfgPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			//fmt.Println("Config does not exist; Will create a new one...")
			os.MkdirAll(filepath.Dir(cfgPath), 0744)
		} else {
			return fmt.Errorf("SKIPING: Could not read existing config: %v", err)
		}
	}

	cfg = patchAndAppendEntries(cfg, entries)

	writeCfg(cfgPath, cfg)

	return nil
}

func patchAndAppendEntries(cfgRows []string, entries []cfgEntry) []string {
	for _, n := range entries {
		replaced := false
		for i, row := range cfgRows {
			rowSplit := strings.Split(row, " = ")
			if rowSplit[0] == n.option {
				cfgRows[i] = n.String()
				//fmt.Printf("Replacing %v with %v\n", n.option, n.value)
				//oldVal := strings.Trim(rowSplit[1], "\"")
				//fmt.Printf("%-40s %v -> %v\n", fmt.Sprintf("Replacing %v", n.option), oldVal, n.value)
				replaced = true
				break
			}
		}
		if !replaced {
			//fmt.Printf("Appending %v\n", n.String())
			cfgRows = append(cfgRows, n.String())
		}
	}

	return cfgRows
}

// readCfg returns the lines of a cfg file as slice of string
func readCfg(cfgPath string) ([]string, error) {
	cfg, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer cfg.Close()

	scanner := bufio.NewScanner(cfg)
	scanner.Split(bufio.ScanLines)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows, nil
}

func writeCfg(cfgPath string, rows []string) error {
	f, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer f.Close()
	cfgContent := strings.Join(rows, "\n")
	f.Write([]byte(cfgContent))
	return nil
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
