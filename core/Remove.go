package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RemoveBen(thisBen string) error {
	if !strings.HasSuffix(thisBen, ".ben") {
		thisBen += ".ben"
	}
	if _, err := os.Stat(filepath.Join(BenDir, thisBen)); os.IsNotExist(err) {
		return fmt.Errorf("Ben '%s' not found!", thisBen)
	}

	meta, err := ParseBen(BenDir, thisBen)
	if err != nil {
		return err
	}
	fmt.Println("Removing package", meta.Name+"@"+meta.Version)
	fmt.Println("Do you want to continue removing? y/n")
	var do string
	fmt.Scanln(&do)
	switch strings.ToLower(do) {
	case "y", "yes":
		if _, err := os.Stat(meta.Script); err == nil {
			if err := os.Remove(filepath.Join(BenDir, "scripts", meta.Name)); err != nil {
				fmt.Printf("Warning: could not remove script! %v\n", err)
			}
		} else {
			fmt.Println("Script not found, skipping...")
		}
		err = os.Remove(filepath.Join(BenDir, thisBen))
		if err != nil {
			return fmt.Errorf("Error removing Ben! %v", err)
		}

		fmt.Println("Ben and script of", meta.Name+"@"+meta.Version, "were removed successfully!")

		return nil
	default:
		fmt.Println("Removing cancelled")
	}

	return nil
}
