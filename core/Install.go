package core

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Install(url string) error {
	CheckDir()

	this := filepath.Base(url)

	if _, err := os.Stat(filepath.Join(BenDir, this)); err == nil {
		return fmt.Errorf("File already exists: %s", this)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("Error checking file! %v", err)
	}

	fmt.Println("Installing script", this)
	fmt.Println("Do you want to continue installing? y/n")
	var do string
	fmt.Scanln(&do)
	switch strings.ToLower(do) {
	case "y", "yes":
		target, err := os.Create(filepath.Join(BenDir, this))
		if err != nil {
			return fmt.Errorf("Error creating a script file! %v", err)
		}
		defer target.Close()

		fmt.Println("Downloading script...")

		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error downloading script! %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Bad status! %s", resp.Status)
		}

		_, err = io.Copy(target, resp.Body)
		if err != nil {
			return fmt.Errorf("Error copying! %v", err)
		}

		fmt.Println("Making script executable...")

		err = exec.Command("chmod", "+x", filepath.Join(BenDir, this)).Run()
		if err != nil {
			return fmt.Errorf("Failed making script executable! %v", err)
		}

		fmt.Println("Successfully installed", this)
		return nil
	default:
		fmt.Println("Installation canceled")
	}

	return nil
}
