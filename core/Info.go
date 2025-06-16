package core

import (
	"fmt"
	"path/filepath"
	"strings"
)

func InfoBen(thisBen string) error {
	if !strings.HasSuffix(thisBen, ".ben") {
		thisBen += ".ben"
	}
	meta, err := ParseBen(BenDir, thisBen)
	if err != nil {
		fmt.Printf("   %s (corrupted!)\n", strings.TrimSuffix(thisBen, ".ben"))
	}

	relScriptPath, err := filepath.Rel(BenDir, meta.Script)
	if err != nil {
		return fmt.Errorf("Failed getting relative path to %s! %v", meta.Script, err)
	}

	fmt.Println()
	fmt.Printf("  %s@%s\n", meta.Name, meta.Version)
	fmt.Printf("  %s\n", meta.Description)
	fmt.Printf("  Script: %s\n", relScriptPath)

	return nil
}
