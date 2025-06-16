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

func InstallBen(url string) error {
	CheckDir()

	thisBen := filepath.Base(url)
	if !strings.HasSuffix(thisBen, ".ben") {
		return fmt.Errorf("%s is not a .ben file", thisBen)
	}

	if _, err := os.Stat(filepath.Join(BenDir, thisBen)); err == nil {
		return fmt.Errorf("File already exists: %s", thisBen)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("Error checking file! %v", err)
	}

	err := DownloadBen(url)
	if err != nil {
		return err
	}

	meta, err := ParseBen(filepath.Join(BenDir, "temp"), thisBen)
	if err != nil {
		return err
	}

	if meta.Version == "@" {
		meta.Version = "rolling"
	}

	fmt.Println("Installing package", meta.Name+"@"+meta.Version)
	fmt.Println("Do you want to continue installing? y/n")
	var do string
	fmt.Scanln(&do)
	switch strings.ToLower(do) {
	case "y", "yes":
		err := os.MkdirAll(filepath.Join(BenDir, "scripts", meta.Name), 0755)
		if err != nil {
			return fmt.Errorf("Failed to create Ben's script directory! %v", err)
		}
		target, err := os.Create(filepath.Join(BenDir, "scripts", meta.Name, filepath.Base(meta.Script)))
		if err != nil {
			return fmt.Errorf("Error creating a script file! %v", err)
		}
		defer target.Close()

		fmt.Println("Downloading script...")

		resp, err := http.Get(meta.Script)
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

		err = exec.Command("chmod", "+x", filepath.Join(BenDir, "scripts", meta.Name, filepath.Base(meta.Script))).Run()
		if err != nil {
			return fmt.Errorf("Failed making script executable! %v", err)
		}

		benFile, err := os.Create(filepath.Join(BenDir, meta.Name+".ben"))
		if err != nil {
			return fmt.Errorf("Failed to create ben file! %v", err)
		}
		defer benFile.Close()

		_, err = fmt.Fprintf(benFile, `name: "%s"
version: "%s"
description: "%s"
script: "%s"
`,
			meta.Name,
			meta.Version,
			meta.Description,
			filepath.Join(BenDir, "scripts", meta.Name, filepath.Base(meta.Script)))
		if err != nil {
			return fmt.Errorf("Failed to write ben file! %v", err)
		}

		err = os.Remove(filepath.Join(BenDir, "temp", thisBen))
		if err != nil {
			return fmt.Errorf("Error cleaning temp! %v", err)
		}

		fmt.Println("Successfully installed", meta.Name+"@"+meta.Version)
		return nil
	default:
		fmt.Println("Installation canceled")
		err := os.Remove(filepath.Join(BenDir, "temp", thisBen))
		if err != nil {
			return fmt.Errorf("Error cleaning temp! %v", err)
		}
		return nil
	}

	return nil
}
