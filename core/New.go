package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewBen() error {
	CheckDir()

	var name string
	var version string
	var description string
	var script string

	fmt.Println("New Ben creator menu")
	fmt.Println("Name for a new Ben: ")
	fmt.Scanln(&name)
	fmt.Println("Version of a new Ben: ")
	fmt.Scanln(&version)
	if version == "@" {
		version = "rolling"
	}
	fmt.Println("Description of a new Ben: ")
	fmt.Scanln(&description)
	fmt.Println("Absolute script path: (format: /home/egor/desktop/script.sh)")
	fmt.Scanln(&script)

	fmt.Println("Please check if this is true: ")
	check := `%s@%s
%s
%s
`
	fmt.Printf(check, name, version, description, script)

	fmt.Println("Where to save Ben? [H]ere or [B]en directory")
	var where string
	var saveTo string
	fmt.Scanln(&where)
	switch strings.ToLower(where) {
	case "here", "h":
		here, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Error getting current dir! %v", err)
		}
		saveTo = here
	case "ben", "b", "ben directory", "ben dir":
		saveTo = BenDir
	default:
		return fmt.Errorf("User didnt choose a dir to write a new Ben")
	}

	if _, err := os.Stat(filepath.Join(saveTo, name+".ben")); err == nil {
		return fmt.Errorf("File already exists: %s", name+".ben")
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("Error checking file! %v", err)
	}

	benFile, err := os.Create(filepath.Join(saveTo, name+".ben"))
	if err != nil {
		return fmt.Errorf("Failed to create ben file! %v", err)
	}
	defer benFile.Close()

	_, err = fmt.Fprintf(benFile, `name: "%s"
version: "%s"
description: "%s"
script: "%s"
`,
		name,
		version,
		description,
		script)
	if err != nil {
		return fmt.Errorf("Failed to write ben file! %v", err)
	}

	if saveTo == BenDir {
		saveTo = "~/.benomad/"
	}
	fmt.Println("A new Ben file was successfully written to", filepath.Join(saveTo, name+".ben"))

	return nil
}
