package core

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadBen(url string) error {
	CheckDir()

	if !strings.HasSuffix(filepath.Base(url), ".ben") {
		return fmt.Errorf("%s is not a .ben file", filepath.Base(url))
	}
	if _, err := os.Stat(filepath.Join(BenDir, "temp", filepath.Base(url))); err == nil {
		fmt.Println("File", filepath.Base(url), "is already downloaded")
		return nil
	}
	out, err := os.Create(filepath.Join(BenDir, "temp", filepath.Base(url)))
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Printf("Downloading %s...\n", filepath.Base(url))

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status! %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
