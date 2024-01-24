package builder

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func BundleCSS(templatePath string, generatedPath string) error {
	filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".css") {
			outFile := filepath.Join(generatedPath, filepath.Base(path))
			result := api.Build(api.BuildOptions{
				EntryPoints:  []string{path},
				Bundle:       true,
				MinifySyntax: true,
				Outfile:      outFile,
				Write:        true,
			})
			if len(result.Errors) > 0 {
				log.Fatalf("Failed to bundle CSS file %s: %v", path, result.Errors)
			}
			log.Printf("Bundled CSS file written to %s\n", outFile)
		}
		return nil
	})
	return nil
}
