package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func detectShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	return filepath.Base(shell)
}

func Path() error {
	CheckDir()
	shell := detectShell()
	switch shell {
	case "bash", "zsh":
		fmt.Printf(`export PATH="%s:$PATH"`+"\n", BenDir)
	case "fish":
		fmt.Printf(`set -gx PATH %s $PATH`+"\n", BenDir)
	case "nu":
		fmt.Printf(`$env.PATH = ($env.PATH | prepend (%s))`+"\n", BenDir)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported shell: %s\n", shell)
	}
	return nil
}
