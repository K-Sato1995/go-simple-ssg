package builder

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyStaticFiles(sourceDir, destinationDir string) error {
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destinationDir, 0755); err != nil {
			return err
		}
	}
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".css")) {
			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()
			destinationFile, err := os.Create(filepath.Join(destinationDir, info.Name()))
			if err != nil {
				return err
			}
			defer destinationFile.Close()
			_, err = io.Copy(destinationFile, sourceFile)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
