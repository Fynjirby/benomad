package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(this string, args []string) error {
	cmd := exec.Command("/bin/bash", append([]string{filepath.Join(BenDir, this)}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error running %s! %v", this, err)
	}

	return nil
}
