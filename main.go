package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/krlohnes/cooked-goose/internal/processor"
)

func main() {
	// Define flags
	up := flag.Bool("up", false, "Filter to process only *.up.sql files.")
	down := flag.Bool("down", false, "Filter to process only *.down.sql files.")
	overwrite := flag.Bool("overwrite", false, "Overwrite the output directory if it exists.")
	outputDir := flag.String("output-dir", "", "Specify a custom output directory.")
	help := flag.Bool("help", false, "Display this help message.")

	// Parse flags
	flag.Parse()

	// Handle --help flag
	if *help {
		printHelp()
		return
	}

	// Ensure input directory is provided
	if flag.NArg() == 0 {
		fmt.Println("Error: No input directory provided.")
		printHelp()
		os.Exit(1)
	}

	inputDir := flag.Arg(0)

	// Set default output directory if not provided
	defaultOutputDir := inputDir + "_cooked"
	if *outputDir == "" {
		*outputDir = defaultOutputDir
	}

	// Determine processing filter
	filter := ""
	if *up {
		filter = "up"
	} else if *down {
		filter = "down"
	}

	// Process directory using the processor package
	err := processor.ProcessDirectory(inputDir, *outputDir, filter, *overwrite)
	if err != nil {
		fmt.Printf("Error processing directory: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Processing completed successfully.")
}

func printHelp() {
	fmt.Println(`Usage: cooked-goose [directory] [flags]

A command-line utility for processing goose migrations SQL files in a given directory with env sub.
It can filter files, interpolate environment variables, and output to a structured '_cooked' directory.
This can be useful for things like sqlc.

Flags:
  --up             Filter to process only *.up.sql files.
  --down           Filter to process only *.down.sql files.
  --overwrite      Overwrite the output directory if it exists.
  --output-dir     Specify a custom output directory (default is [directory]_cooked).
  --help           Display this help message.

Examples:
  # Process all SQL files in the "migrations" directory
  cooked-goose migrations

  # Process only *.up.sql files and overwrite conflicts
  cooked-goose migrations --up --overwrite

  # Process only *.down.sql files
  cooked-goose migrations --down

  # Process SQL files and output to a custom directory
  cooked-goose migrations --output-dir custom_cooked`)
}
