package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Remove(this string) error {
	if _, err := os.Stat(filepath.Join(BenDir, this)); os.IsNotExist(err) {
		return fmt.Errorf("Script '%s' not found!", this)
	}

	fmt.Println("Removing script", this)
	fmt.Println("Do you want to continue? y/n")
	var do string
	fmt.Scanln(&do)
	switch strings.ToLower(do) {
	case "y", "yes":
		err := os.Remove(filepath.Join(BenDir, this))
		if err != nil {
			return fmt.Errorf("Error removing! %v", err)
		}

		fmt.Println("Script", this, "were removed successfully!")

		return nil
	default:
		fmt.Println("Removing cancelled")
	}

	return nil
}
