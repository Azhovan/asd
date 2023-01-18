package config

import (
	"fmt"
	"os"
)

// GetDefaultASDDirectory returns the default directory where all notes exist.
func GetDefaultASDDirectory() (string, error) {
	userhome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.asd", userhome), nil
}
