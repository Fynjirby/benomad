package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func detectShell() string {
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", os.Getppid()), "-o", "comm=")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "unknown"
	}
	trimmed := strings.TrimSpace(out.String())
	splitted := strings.Split(trimmed, "/")
	shell := strings.ReplaceAll(splitted[len(splitted)-1], "-", "")
	return shell
}

func Path() error {
	CheckDir()
	ben := BenDir
	shell := detectShell()
	switch shell {
	case "bash", "zsh":
		fmt.Printf(`export PATH="%s:$PATH"`+"\n", ben)
	case "fish":
		fmt.Printf(`set -gx PATH %s $PATH`+"\n", ben)
	case "nu":
		fmt.Printf(`$env.PATH = ($env.PATH | prepend (%s))`+"\n", ben)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported shell: %s\n", shell)
	}
	return nil
}
