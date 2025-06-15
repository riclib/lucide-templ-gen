package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	lucidegen "github.com/riclib/lucide-templ-gen"
)

const version = "1.1.0"

func main() {
	var (
		outputDir   = flag.String("output", ".", "Output directory")
		packageName = flag.String("package", "icons", "Package name")
		prefix      = flag.String("prefix", "", "Function name prefix")
		categories  = flag.String("categories", "", "Comma-separated categories to include (empty = all)")
		dryRun      = flag.Bool("dry-run", false, "Show what would be generated without creating files")
		verbose     = flag.Bool("verbose", false, "Enable verbose output")
		showVersion = flag.Bool("version", false, "Show version information")
		help        = flag.Bool("help", false, "Show help information")
	)

	flag.Parse()

	if *help {
		showHelp()
		return
	}

	if *showVersion {
		fmt.Printf("lucide-gen version %s\n", version)
		return
	}

	// Parse categories
	var categoryList []string
	if *categories != "" {
		categoryList = strings.Split(*categories, ",")
		for i, cat := range categoryList {
			categoryList[i] = strings.TrimSpace(cat)
		}
	}

	// Create config
	config := lucidegen.Config{
		OutputDir:   *outputDir,
		PackageName: *packageName,
		Prefix:      *prefix,
		Categories:  categoryList,
		DryRun:      *dryRun,
		Verbose:     *verbose,
	}

	// Validate config
	if err := validateConfig(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Generate icons
	result, err := lucidegen.Generate(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Generation failed: %v\n", err)
		os.Exit(1)
	}

	// Show results
	if !config.DryRun {
		fmt.Printf("✅ Successfully generated %d icons in %v\n", result.IconsGenerated, result.Duration)
		if config.Verbose {
			fmt.Printf("📁 Files created:\n")
			for _, file := range result.FilesCreated {
				fmt.Printf("   %s\n", file)
			}
			fmt.Printf("📂 Categories: %s\n", strings.Join(result.Categories, ", "))
		}
	}
}

func validateConfig(config lucidegen.Config) error {
	if config.OutputDir == "" {
		return fmt.Errorf("output directory cannot be empty")
	}

	if config.PackageName == "" {
		return fmt.Errorf("package name cannot be empty")
	}

	// Validate package name
	if !isValidPackageName(config.PackageName) {
		return fmt.Errorf("invalid package name: %s", config.PackageName)
	}

	return nil
}

func isValidPackageName(name string) bool {
	if len(name) == 0 {
		return false
	}

	// Must start with letter or underscore
	if !isLetter(rune(name[0])) && name[0] != '_' {
		return false
	}

	// Rest can be letters, digits, or underscores
	for _, r := range name[1:] {
		if !isLetter(r) && !isDigit(r) && r != '_' {
			return false
		}
	}

	return true
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func showHelp() {
	fmt.Printf(`lucide-gen - Generate type-safe Templ components for Lucide Icons

Usage:
  lucide-gen [options]

Options:
  -output string      Output directory (default ".")
  -package string     Package name (default "icons")
  -prefix string      Function name prefix (default "")
  -categories string  Comma-separated categories to include (default: all)
  -dry-run           Show what would be generated without creating files
  -verbose           Enable verbose output
  -version           Show version information
  -help              Show this help message

Examples:
  # Generate all icons in current directory
  lucide-gen

  # Generate specific categories
  lucide-gen -categories "navigation,actions,media"

  # Generate with custom package and prefix
  lucide-gen -output ./icons -package icons -prefix Lucide

  # Dry run to preview
  lucide-gen -dry-run -verbose

Categories:
  navigation      - home, menu, chevron-*, arrow-*, etc.
  actions         - plus, minus, edit, trash, save, etc.
  media          - play, pause, stop, volume, etc.
  communication  - mail, phone, message, bell, etc.
  files          - file, folder, download, upload, etc.
  ui             - eye, lock, search, check, etc.
  data           - database, server, cloud, chart, etc.
  devices        - smartphone, laptop, monitor, etc.
  social         - heart, star, share, thumbs-up, etc.
  weather        - sun, moon, cloud, rain, etc.
  transportation - car, plane, bike, etc.
  business       - briefcase, building, wallet, etc.

For more information, visit: https://github.com/riclib/lucide-templ-gen
`)
}