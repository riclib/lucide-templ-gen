package lucidegen

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Config holds the configuration for icon generation
type Config struct {
	OutputDir   string   // Output directory path
	PackageName string   // Go package name
	Prefix      string   // Function name prefix
	Categories  []string // Icon categories to include (empty = all)
	DryRun      bool     // Preview without generating files
	Verbose     bool     // Enable verbose logging
}

// IconData represents a parsed Lucide icon
type IconData struct {
	Name       string `json:"name"`
	FuncName   string `json:"func_name"`
	ViewBox    string `json:"view_box"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	Keywords   []string `json:"keywords"`
	Deprecated bool   `json:"deprecated"`
}

// GenerationResult contains information about the generation process
type GenerationResult struct {
	IconsGenerated int      `json:"icons_generated"`
	FilesCreated   []string `json:"files_created"`
	Categories     []string `json:"categories"`
	Duration       time.Duration `json:"duration"`
}

// SVGElement represents the parsed SVG structure
type SVGElement struct {
	ViewBox string `xml:"viewBox,attr"`
	Width   string `xml:"width,attr"`
	Height  string `xml:"height,attr"`
	Content string `xml:",innerxml"`
}

// GitHubFile represents a file in GitHub API response
type GitHubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

var (
	// Icon categories mapping
	iconCategories = map[string][]string{
		"navigation": {
			"home", "menu", "chevron-up", "chevron-down", "chevron-left", "chevron-right",
			"arrow-up", "arrow-down", "arrow-left", "arrow-right", "arrow-up-right",
			"arrow-down-right", "arrow-down-left", "arrow-up-left", "corner-up-left",
			"corner-up-right", "corner-down-left", "corner-down-right", "move",
			"move-diagonal", "move-horizontal", "move-vertical", "navigation",
			"compass", "map", "map-pin", "route", "signpost",
		},
		"actions": {
			"plus", "minus", "edit", "edit-2", "edit-3", "trash", "trash-2", "save",
			"copy", "clipboard", "clipboard-copy", "clipboard-list", "cut", "scissors",
			"undo", "redo", "refresh-cw", "refresh-ccw", "rotate-cw", "rotate-ccw",
			"flip-horizontal", "flip-vertical", "maximize", "minimize", "maximize-2",
			"minimize-2", "zoom-in", "zoom-out", "search", "filter", "sort-asc",
			"sort-desc", "more-horizontal", "more-vertical", "settings", "sliders",
		},
		"media": {
			"play", "pause", "stop", "skip-forward", "skip-back", "fast-forward",
			"rewind", "volume", "volume-1", "volume-2", "volume-x", "mic", "mic-off",
			"video", "video-off", "camera", "camera-off", "image", "film", "music",
			"headphones", "speaker", "radio", "tv", "monitor", "smartphone", "tablet",
		},
		"communication": {
			"mail", "send", "inbox", "outbox", "phone", "phone-call", "phone-incoming",
			"phone-outgoing", "phone-off", "message-circle", "message-square",
			"chat", "users", "user", "user-plus", "user-minus", "user-check",
			"user-x", "at-sign", "bell", "bell-off", "bell-ring", "notification",
		},
		"files": {
			"file", "file-text", "file-plus", "file-minus", "file-x", "files",
			"folder", "folder-open", "folder-plus", "folder-minus", "folder-x",
			"hard-drive", "download", "upload", "import", "export", "paperclip",
			"link", "link-2", "external-link", "archive", "package", "package-2",
		},
		"ui": {
			"eye", "eye-off", "lock", "unlock", "key", "shield", "shield-check",
			"shield-alert", "shield-x", "check", "check-circle", "check-circle-2",
			"x", "x-circle", "alert-triangle", "alert-circle", "alert-octagon",
			"info", "help-circle", "question-mark", "loader", "loader-2", "circle",
			"square", "triangle", "diamond", "heart", "star", "bookmark",
		},
		"data": {
			"database", "server", "cloud", "cloud-snow", "cloud-rain", "cloud-lightning",
			"wifi", "wifi-off", "signal", "activity", "trending-up", "trending-down",
			"bar-chart", "bar-chart-2", "bar-chart-3", "bar-chart-4", "pie-chart",
			"line-chart", "area-chart", "git-branch", "git-commit", "git-merge",
			"git-pull-request", "github", "gitlab", "code", "code-2", "terminal",
		},
		"devices": {
			"smartphone", "tablet", "laptop", "monitor", "tv", "watch", "gamepad",
			"gamepad-2", "keyboard", "mouse", "printer", "scanner", "usb", "bluetooth",
			"battery", "battery-charging", "battery-full", "battery-low", "plug",
			"power", "power-off", "cpu", "memory-stick", "hard-drive", "disc",
		},
		"social": {
			"heart", "star", "thumbs-up", "thumbs-down", "share", "share-2", "rss",
			"bookmark", "flag", "award", "trophy", "medal", "gift", "cake", "party-popper",
			"smile", "frown", "meh", "laugh", "angry", "surprised", "wink", "kiss",
			"facebook", "twitter", "instagram", "linkedin", "youtube", "twitch",
		},
		"weather": {
			"sun", "moon", "star", "cloud", "cloud-drizzle", "cloud-rain", "cloud-snow",
			"cloud-lightning", "umbrella", "wind", "tornado", "rainbow", "sunrise",
			"sunset", "thermometer", "thermometer-sun", "thermometer-snowflake",
			"droplets", "waves", "zap", "flame", "snowflake", "tree-pine",
		},
		"transportation": {
			"car", "truck", "bus", "train", "plane", "ship", "bike", "scooter",
			"taxi", "fuel", "map", "map-pin", "route", "compass", "navigation",
			"anchor", "sail", "wheel", "tire", "traffic-cone", "construction",
		},
		"business": {
			"briefcase", "building", "building-2", "factory", "store", "bank",
			"credit-card", "wallet", "coins", "banknote", "receipt", "calculator",
			"presentation", "chart", "graph", "analytics", "target", "goal",
			"handshake", "deal", "contract", "signature", "stamp", "scale",
		},
	}
)

// Generate creates Lucide icon components based on the provided configuration
func Generate(config Config) (*GenerationResult, error) {
	start := time.Now()
	
	if config.Verbose {
		fmt.Printf("Starting Lucide icon generation...\n")
		fmt.Printf("Output directory: %s\n", config.OutputDir)
		fmt.Printf("Package name: %s\n", config.PackageName)
	}

	// Set defaults
	if config.PackageName == "" {
		config.PackageName = "icons"
	}

	// Fetch icons from GitHub
	icons, err := fetchLucideIcons(config.Verbose)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch icons: %w", err)
	}

	// Filter by categories if specified
	if len(config.Categories) > 0 {
		icons = filterIconsByCategories(icons, config.Categories)
	}

	// Sort icons by name for consistent output
	sort.Slice(icons, func(i, j int) bool {
		return icons[i].Name < icons[j].Name
	})

	if config.Verbose {
		fmt.Printf("Found %d icons to generate\n", len(icons))
	}

	result := &GenerationResult{
		IconsGenerated: len(icons),
		Categories:     getUniqueCategories(icons),
		Duration:       time.Since(start),
	}

	if config.DryRun {
		fmt.Printf("DRY RUN: Would generate %d icons\n", len(icons))
		for _, icon := range icons {
			fmt.Printf("  - %s (%s)\n", icon.FuncName, icon.Category)
		}
		return result, nil
	}

	// Create output directory
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate files
	files, err := generateFiles(icons, config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate files: %w", err)
	}

	result.FilesCreated = files
	result.Duration = time.Since(start)

	if config.Verbose {
		fmt.Printf("Generation completed in %v\n", result.Duration)
		fmt.Printf("Created files:\n")
		for _, file := range files {
			fmt.Printf("  - %s\n", file)
		}
	}

	return result, nil
}

// fetchLucideIcons retrieves icon data from the Lucide GitHub repository
func fetchLucideIcons(verbose bool) ([]IconData, error) {
	if verbose {
		fmt.Println("Fetching icon list from GitHub...")
	}

	// Get file list from GitHub API
	resp, err := http.Get("https://api.github.com/repos/lucide-icons/lucide/contents/icons")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file list: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var files []GitHubFile
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, fmt.Errorf("failed to decode file list: %w", err)
	}

	var icons []IconData
	for i, file := range files {
		if !strings.HasSuffix(file.Name, ".svg") {
			continue
		}

		if verbose && i%50 == 0 {
			fmt.Printf("Processing icon %d/%d...\n", i+1, len(files))
		}

		iconName := strings.TrimSuffix(file.Name, ".svg")
		
		// Download and parse SVG
		iconData, err := downloadAndParseSVG(file.DownloadURL, iconName)
		if err != nil {
			if verbose {
				fmt.Printf("Warning: failed to process %s: %v\n", iconName, err)
			}
			continue
		}

		icons = append(icons, *iconData)
	}

	return icons, nil
}

// downloadAndParseSVG downloads and parses an SVG file
func downloadAndParseSVG(url, name string) (*IconData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	svgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var svg SVGElement
	if err := xml.Unmarshal(svgData, &svg); err != nil {
		return nil, err
	}

	// Determine category
	category := categorizeIcon(name)

	return &IconData{
		Name:     name,
		FuncName: toFunctionName(name),
		ViewBox:  svg.ViewBox,
		Content:  strings.TrimSpace(svg.Content),
		Category: category,
	}, nil
}

// categorizeIcon determines the category of an icon based on its name
func categorizeIcon(iconName string) string {
	for category, icons := range iconCategories {
		for _, icon := range icons {
			if icon == iconName {
				return category
			}
		}
	}
	return "misc"
}

// filterIconsByCategories filters icons to only include specified categories
func filterIconsByCategories(icons []IconData, categories []string) []IconData {
	categorySet := make(map[string]bool)
	for _, cat := range categories {
		categorySet[strings.ToLower(cat)] = true
	}

	var filtered []IconData
	for _, icon := range icons {
		if categorySet[icon.Category] {
			filtered = append(filtered, icon)
		}
	}

	return filtered
}

// getUniqueCategories returns a sorted list of unique categories
func getUniqueCategories(icons []IconData) []string {
	categorySet := make(map[string]bool)
	for _, icon := range icons {
		categorySet[icon.Category] = true
	}

	var categories []string
	for cat := range categorySet {
		categories = append(categories, cat)
	}

	sort.Strings(categories)
	return categories
}

// toFunctionName converts an icon name to a valid Go function name
func toFunctionName(name string) string {
	// Convert kebab-case to PascalCase
	parts := strings.Split(name, "-")
	var result strings.Builder
	
	for _, part := range parts {
		if len(part) > 0 {
			// Capitalize first letter, keep rest as-is
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(part[1:])
			}
		}
	}
	
	// Ensure it starts with a letter
	funcName := result.String()
	if len(funcName) > 0 && !isLetter(rune(funcName[0])) {
		funcName = "Icon" + funcName
	}
	
	return funcName
}

// isLetter checks if a rune is a letter
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// toConstantName converts an icon name to a constant name
func toConstantName(name, prefix string) string {
	funcName := toFunctionName(name)
	if prefix != "" {
		return prefix + funcName
	}
	return "Icon" + funcName
}

// generateFiles creates all the template files
func generateFiles(icons []IconData, config Config) ([]string, error) {
	var createdFiles []string

	// Generate main icons file
	iconsFile := filepath.Join(config.OutputDir, "icons.templ")
	if err := generateIconsFile(icons, config, iconsFile); err != nil {
		return nil, fmt.Errorf("failed to generate icons file: %w", err)
	}
	createdFiles = append(createdFiles, iconsFile)

	// Generate registry file
	registryFile := filepath.Join(config.OutputDir, "registry.templ")
	if err := generateRegistryFile(icons, config, registryFile); err != nil {
		return nil, fmt.Errorf("failed to generate registry file: %w", err)
	}
	createdFiles = append(createdFiles, registryFile)

	// Generate categories file
	categoriesFile := filepath.Join(config.OutputDir, "categories.templ")
	if err := generateCategoriesFile(icons, config, categoriesFile); err != nil {
		return nil, fmt.Errorf("failed to generate categories file: %w", err)
	}
	createdFiles = append(createdFiles, categoriesFile)

	return createdFiles, nil
}