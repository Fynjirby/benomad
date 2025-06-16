package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// The utils file
// here mostly sub functions and other little things

func CheckDir() {
	if err := os.MkdirAll(filepath.Join(BenDir, "temp"), 0755); err != nil {
		fmt.Println("Failed to create temp directory!", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Join(BenDir, "scripts"), 0755); err != nil {
		fmt.Println("Failed to create scripts directory!", err)
		os.Exit(1)
	}
}
