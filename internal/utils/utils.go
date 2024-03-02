package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	// Copy the file from the HTTP response to the local file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
func CreateDirectory(path string) error {
	// Check if the directory exists
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	// If the directory does not exist, create its parent
	if os.IsNotExist(err) {
		err = CreateDirectory(filepath.Dir(path))
		if err != nil {
			return err
		}
		// Create the directory
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
