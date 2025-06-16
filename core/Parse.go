package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// some things in type, needa even more
type Meta struct {
	Name        string
	Version     string
	Description string
	Script      string
}

func ParseBen(path string, file string) (*Meta, error) {
	meta := &Meta{}

	if !strings.HasSuffix(file, ".ben") {
		file += ".ben"
	}
	data, err := os.ReadFile(filepath.Join(path, file))
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		if len(val) >= 2 && val[0] == '"' && val[len(val)-1] == '"' {
			val = val[1 : len(val)-1]
		}

		switch key {
		case "name":
			meta.Name = val
		case "version":
			meta.Version = val
		case "description":
			meta.Description = val
		case "script":
			meta.Script = val
		}
	}

	var missingFields []string
	if meta.Name == "" {
		missingFields = append(missingFields, "name")
	}
	if meta.Version == "" {
		missingFields = append(missingFields, "version")
	}
	if meta.Description == "" {
		missingFields = append(missingFields, "description")
	}
	if meta.Script == "" {
		missingFields = append(missingFields, "script")
	}

	if len(missingFields) > 0 {
		return nil, fmt.Errorf("Missing values! %s\n Can't parse this ben", strings.Join(missingFields, ", "))
	}

	return meta, nil
}
