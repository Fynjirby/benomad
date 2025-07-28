package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Edit(this string) error {
	if _, err := os.Stat(filepath.Join(BenDir, this)); os.IsNotExist(err) {
		return fmt.Errorf("Script '%s' not found!", this)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	cmd := exec.Command(editor, filepath.Join(BenDir, this))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error opening %s! %v", this, err)
	}

	return nil
}
