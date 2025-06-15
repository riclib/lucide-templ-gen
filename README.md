# Lucide Templ Generator

A Go code generator that creates type-safe [Templ](https://templ.guide) components for [Lucide Icons](https://lucide.dev).

## Features

- 🎯 **Type-safe**: All icon names are compile-time checked
- 🚀 **Zero runtime dependencies**: Pure Go/Templ components
- 📦 **Tree-shakeable**: Only used icons are included
- 🔄 **Auto-updating**: Fetch latest icons from Lucide repository
- 🏷️ **Categorized**: Icons organized by logical categories
- 🎨 **Customizable**: Full control over attributes and styling
- 📝 **Well-documented**: Generated code includes documentation

## Installation

### Install CLI Tool

```bash
go install github.com/riclib/lucide-templ-gen/cmd/lucide-gen@latest
```

### Add as Dependency

```bash
go get github.com/riclib/lucide-templ-gen
```

### Download Binary

Download pre-built binaries from the [releases page](https://github.com/riclib/lucide-templ-gen/releases):

- **Linux**: `lucide-gen-linux-amd64.tar.gz`
- **macOS**: `lucide-gen-darwin-amd64.tar.gz` (Intel) or `lucide-gen-darwin-arm64.tar.gz` (Apple Silicon)
- **Windows**: `lucide-gen-windows-amd64.zip`

## Quick Start

### Command Line Usage

```bash
# Generate all icons
lucide-gen -output ./icons

# Generate specific categories
lucide-gen -output ./icons -categories "navigation,actions,media"

# Generate with custom package name
lucide-gen -output ./ui/icons -package icons -prefix Icon

# Dry run to see what would be generated
lucide-gen -output ./icons -dry-run
```

### Programmatic Usage

```go
package main

import (
    "log"
    lucidegen "github.com/riclib/lucide-templ-gen"
)

func main() {
    config := lucidegen.Config{
        OutputDir:   "./web/components/icons",
        PackageName: "icons",
        Prefix:      "Lucide",
        Categories:  []string{"navigation", "actions"},
    }
    
    if err := lucidegen.Generate(config); err != nil {
        log.Fatal(err)
    }
}
```

### Integration with Taskfile

```yaml
# Taskfile.yml
tasks:
  icons:generate:
    desc: Generate Lucide icon components
    cmds:
      - lucide-gen -output web/templates/components -package components
      - templ generate web/templates/components/icons_templ.go

  icons:update:
    desc: Update to latest Lucide icons
    cmds:
      - task: icons:generate
```

## Usage in Templates

After generation, use the type-safe icon components:

```go
package main

import "your-project/web/templates/components"

// Direct icon usage
templ HomePage() {
    <div class="nav">
        @components.Home(templ.Attributes{"class": "nav-icon"})
        @components.User(templ.Attributes{"class": "user-icon", "size": "24"})
    </div>
}

// Dynamic icon usage with registry
templ IconButton(iconName components.IconName, label string) {
    <button class="btn">
        @components.Icon(iconName, templ.Attributes{"class": "btn-icon"})
        { label }
    </button>
}

// Usage with type safety
templ Dashboard() {
    @IconButton(components.IconHome, "Home")
    @IconButton(components.IconSettings, "Settings")
}
```

## Generated API

The generator creates several files:

### `icons_templ.go`
Individual icon components:
```go
// Home renders the home Lucide icon
templ Home(attrs templ.Attributes) { /* ... */ }

// User renders the user Lucide icon  
templ User(attrs templ.Attributes) { /* ... */ }
```

### `registry_templ.go`
Type-safe registry and utilities:
```go
type IconName string

const (
    IconHome IconName = "home"
    IconUser IconName = "user"
)

// Icon renders any icon by name
templ Icon(name IconName, attrs templ.Attributes) { /* ... */ }

// IconExists checks if an icon is available
func IconExists(name string) bool { /* ... */ }

// AllIcons returns all available icons
func AllIcons() []IconName { /* ... */ }
```

### `categories_templ.go`
Categorized icon access:
```go
// NavigationIcons returns all navigation-related icons
func NavigationIcons() []IconName { /* ... */ }

// ActionIcons returns all action-related icons  
func ActionIcons() []IconName { /* ... */ }
```

## Configuration

### Command Line Options

```
Usage: lucide-gen [options]

Options:
  -output string      Output directory (default ".")
  -package string     Package name (default "icons")
  -prefix string      Function name prefix (default "")
  -categories string  Comma-separated categories to include (default: all)
  -dry-run           Show what would be generated without creating files
  -version           Show version information
  -help              Show this help message
```

### Programmatic Config

```go
type Config struct {
    OutputDir    string   // Output directory path
    PackageName  string   // Go package name
    Prefix       string   // Function name prefix
    Categories   []string // Icon categories to include
    DryRun       bool     // Preview without generating
    Verbose      bool     // Enable verbose output
}
```

## Icon Categories

Icons are organized into logical categories:

- **Navigation**: home, menu, chevron-*, arrow-*, etc.
- **Actions**: plus, minus, edit, trash, save, copy, etc.
- **Media**: play, pause, stop, volume, music, video, etc.
- **Communication**: mail, phone, message, bell, etc.
- **Files**: file, folder, download, upload, etc.
- **UI**: eye, eye-off, lock, unlock, search, etc.
- **Data**: database, server, cloud, wifi, etc.
- **Devices**: smartphone, laptop, monitor, etc.
- **Social**: heart, star, share, thumbs-up, etc.
- **Weather**: sun, moon, cloud, umbrella, etc.

## Examples

Check the [examples](./examples) directory for complete integration examples:

- [Basic Usage](./examples/basic) - Simple icon generation
- [Custom Categories](./examples/categories) - Category-specific generation  
- [Advanced Integration](./examples/advanced) - Full project integration
- [Taskfile Integration](./examples/taskfile) - Build system integration

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Lucide Icons](https://lucide.dev) for the beautiful icon set
- [Templ](https://templ.guide) for the excellent templating system