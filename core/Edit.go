package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func EditBen(thisBen string) error {
	if !strings.HasSuffix(thisBen, ".ben") {
		thisBen += ".ben"
	}

	if _, err := os.Stat(filepath.Join(BenDir, thisBen)); os.IsNotExist(err) {
		return fmt.Errorf("Ben '%s' not found!", thisBen)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	fmt.Println("What to edit: [B]en or [S]cript?")
	var what string
	fmt.Scanln(&what)
	switch strings.ToLower(what) {
	case "ben", "b":
		cmd := exec.Command(editor, filepath.Join(BenDir, thisBen))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("Error opening %s! %v", thisBen)
		}
	case "script", "s":
		meta, err := ParseBen(BenDir, thisBen)
		if err != nil {
			fmt.Errorf("Error parsing %s! %v", thisBen, err)
		}

		cmd := exec.Command(editor, meta.Script)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("Error opening %s! %v", thisBen, err)
		}
	}

	return nil
}
