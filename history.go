package main

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	DirName = ".curry_history"
)

func buildHistoryDirPath() string {
	return filepath.Join(os.Getenv("HOME"), DirName)
}

func buildHistoryFilePath(baseCommand []string) string {
	return filepath.Join(buildHistoryDirPath(), baseCommand[0])
}

func initHistory(baseCommand []string) error {
	dir := buildHistoryDirPath()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	filePath := buildHistoryFilePath(baseCommand)

	if _, err := os.Stat(filePath); err == os.ErrNotExist {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}

func readHistory(baseCommand []string) ([]string, error) {
	file, err := os.ReadFile(buildHistoryFilePath(baseCommand))
	if err != nil {
		return nil, err
	}

	return strings.Split(string(file), "\n"), nil
}

func saveHistory(baseCommand []string, input string) error {
	file, err := os.OpenFile(buildHistoryFilePath(baseCommand), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(input + "\n"); err != nil {
		return err
	}

	return nil
}
