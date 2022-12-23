package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (r RetroArchChanger) setSettings(cfgEntries []cfgEntry, subDir, fileName string) Message {
	err := r.backupCfg(subDir, fileName)
	if err != nil && !os.IsNotExist(err) {
		return Message{Success: false, Text: fmt.Sprintf("%s (%s); Couldn't backup existing cfg, therefore skipping: %v", fileName, subDir, err.Error())}
	}
	err = updateCfg(filepath.Join(r.retroarchCfgDirPath, subDir, fileName), cfgEntries)
	if err != nil {
		return Message{Success: false, Text: fmt.Sprintf("%s (%s); Couldn't optimize config: %v", fileName, subDir, err)}
	}
	return Message{Success: true, Text: fmt.Sprintf("%v (%s)", fileName, subDir)}
}

func (r RetroArchChanger) backupCfg(subDir, fileName string) error {
	backupPath := filepath.Join(r.retroarchCfgDirPath, "_shmuparchify_backup", r.timeStamp, subDir)
	err := os.MkdirAll(backupPath, 0755)
	if err != nil {
		return err
	}

	cfgFilePath := filepath.Join(r.retroarchCfgDirPath, subDir, fileName)
	backupFilePath := filepath.Join(backupPath, fileName)

	// make copy
	original, err := os.Open(cfgFilePath)
	if err != nil {
		return err
	}
	defer original.Close()

	new, err := os.Create(backupFilePath)
	if err != nil {
		return err
	}
	defer new.Close()

	_, err = io.Copy(new, original)
	if err != nil {
		return err
	}
	return nil
}

func updateCfg(cfgPath string, entries []cfgEntry) error {
	cfg, err := readCfg(cfgPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.MkdirAll(filepath.Dir(cfgPath), 0755)
		} else {
			return fmt.Errorf("could not read existing config: %v", err)
		}
	}

	cfg = patchAndAppendEntries(cfg, entries)

	err = writeCfg(cfgPath, cfg)
	if err != nil {
		return err
	}

	return nil
}

func patchAndAppendEntries(cfgRows []string, entries []cfgEntry) []string {
	for _, n := range entries {
		replaced := false
		for i, row := range cfgRows {
			rowSplit := strings.Split(row, " = ")
			if strings.TrimSuffix(strings.TrimPrefix(rowSplit[0], "\""), "\"") == n.option {
				cfgRows[i] = n.String()
				replaced = true
				break
			}
		}
		if !replaced {
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
