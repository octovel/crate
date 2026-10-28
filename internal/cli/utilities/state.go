package utilities

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/octovel/crate/internal/cli/models"
)

func GetStatePath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			return "", fmt.Errorf(
				"'LOCALAPPDATA' environmental variable not set.",
			)
		}

		return filepath.Join(localAppData, "Octovel", "Crate", "state.json"), nil

	case "darwin":
		configDir, err := os.UserConfigDir()
		if err != nil {
			return "", fmt.Errorf(
				"Could not find config directory: %w",
				err,
			)
		}

		return filepath.Join(configDir, "Octovel", "Crate", "state.json"), nil

	default: // linux distros
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf(
				"Could not find user home directory: %w",
				err,
			)
		}

		return filepath.Join(
			home, ".local", "state", "octovel", "crate", "state.json",
		), nil
	}
}

func GetState() *models.State {
	path, err := GetStatePath()
	if err != nil {
		return nil
	}

	var state models.State
	fileBytes, _ := os.ReadFile(path)
	err = json.Unmarshal(fileBytes, &state)
	if err != nil {
		return nil
	}

	return &state
}
