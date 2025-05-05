package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mfridman/interpolate"
)

type envWrapper struct{}

var _ interpolate.Env = (*envWrapper)(nil)

func (e *envWrapper) Get(key string) (string, bool) {
	_, b := os.LookupEnv(key)
	if !b {
		//Panicking because the interpolate library doesn't quite work as
		//expected and this is fairly necessary with this tool
		panic(fmt.Sprintf("%s not set in environment", key))
	}
	return os.LookupEnv(key)
}

func ProcessDirectory(inputDir, outputDir string, filter string, overwrite bool) error {
	err := filepath.WalkDir(inputDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// Filter files based on flags
		if (filter == "up" && filepath.Ext(path) != ".up.sql") ||
			(filter == "down" && filepath.Ext(path) != ".down.sql") {
			return nil
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			return err
		}

		// Process content
		processed, err := processFileContent(string(content))
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", path, err)
			return err // Continue processing other files
		}

		// Write to output directory
		relPath, _ := filepath.Rel(inputDir, path)
		outputPath := filepath.Join(outputDir, relPath)
		os.MkdirAll(filepath.Dir(outputPath), 0755)

		// Check if the file exists and handle overwrite logic
		if !overwrite {
			if _, err := os.Stat(outputPath); err == nil {
				fmt.Printf("Skipping existing file %s (overwrite disabled)\n", outputPath)
				return nil
			}
		}

		err = os.WriteFile(outputPath, []byte(processed), 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", outputPath, err)
		}

		return nil
	})

	return err
}

// processFileContent parses the input SQL content and interpolates environment variables
// wrapped in `-- +goose ENVSUB ON` and `-- +goose ENVSUB OFF`.
// If `-- +goose ENVSUB ON` is found without a corresponding `OFF`, it interpolates until the end of the file.
func processFileContent(content string) (string, error) {
	var result strings.Builder
	var isInterpolating bool

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Check for ENVSUB directives
		if trimmedLine == "-- +goose ENVSUB ON" {
			isInterpolating = true
			result.WriteString(line + "\n") // Write the directive line to the output
			continue
		} else if trimmedLine == "-- +goose ENVSUB OFF" {
			isInterpolating = false
			result.WriteString(line + "\n") // Write the directive line to the output
			continue
		}

		// If within an ENVSUB block, interpolate variables
		if isInterpolating {
			fmt.Println("Interpolating " + line)
			interpolatedLine, err := interpolate.Interpolate(&envWrapper{}, line)
			fmt.Println("ERR ", err)
			if err != nil {
				return "", fmt.Errorf("error interpolating line %q: %w", line, err)
			}
			fmt.Println("New line is ", interpolatedLine)
			result.WriteString(interpolatedLine + "\n")
		} else {
			// Write the line as is
			result.WriteString(line + "\n")
		}
	}

	return result.String(), nil
}
