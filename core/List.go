package core

import (
	"fmt"
	"os"
)

func List() error {
	entries, err := os.ReadDir(BenDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}

	return nil
}
