package core

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunBen(thisBen string) error {
	if !strings.HasSuffix(thisBen, ".ben") {
		thisBen += ".ben"
	}

	meta, err := ParseBen(BenDir, thisBen)
	if err != nil {
		fmt.Errorf("Error parsing %s! %v", thisBen, err)
	}

	cmd := exec.Command("/bin/bash", meta.Script)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Error running %s! %v", thisBen, err)
	}

	return nil
}
