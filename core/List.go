package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListBen() error {
	files, err := os.ReadDir(BenDir)
	if err != nil {
		return fmt.Errorf("Error reading ben directory! %v", err)
	}

	hasBens := false
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ben" {
			hasBens = true
			break
		}
	}

	if !hasBens {
		return fmt.Errorf("No Bens found in ~/.benomad!")
	}

	fmt.Println()
	fmt.Printf("  Installed packages in ~/.benomad/\n")

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ben" {
			meta, err := ParseBen(BenDir, file.Name())
			if err != nil {
				fmt.Printf("     %s (corrupted!)\n", strings.TrimSuffix(file.Name(), ".ben"))
				continue
			}

			fmt.Printf("  -> %s@%s\n", meta.Name, meta.Version)
		}
	}

	return nil
}
