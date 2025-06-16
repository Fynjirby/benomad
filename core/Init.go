package core

import (
	"os"
	"path/filepath"
)

// Init variables and consts file
var (
	HomeDir, _ = os.UserHomeDir()
	BenDir     = filepath.Join(HomeDir, ".benomad")
)

// Type Meta is in core/Parse.go!
