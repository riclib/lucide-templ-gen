# Changelog

All notable changes to lucide-templ-gen will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of lucide-templ-gen
- Type-safe Templ components for Lucide Icons
- CLI tool for icon generation
- Support for icon categories
- Programmatic API for custom generation
- GitHub Actions CI/CD pipeline
- Comprehensive documentation and examples

## [1.0.0] - 2024-01-XX

### Added
- **Core Features**
  - Generate type-safe Templ components from Lucide Icons
  - Fetch icons directly from Lucide GitHub repository
  - Support for all 1000+ Lucide icons
  - Automatic categorization of icons (navigation, actions, media, etc.)
  - Tree-shakeable output (only generate needed categories)

- **CLI Tool** (`lucide-gen`)
  - Command-line interface for icon generation
  - Support for custom output directories and package names
  - Dry-run mode for previewing generation
  - Verbose output for debugging
  - Category filtering support

- **Generated API**
  - Individual icon components with full type safety
  - Icon registry with string-to-component mapping
  - Category-based icon grouping
  - Icon existence checking utilities
  - Complete icon enumeration

- **Development Experience**
  - Zero runtime dependencies
  - Full Go module support
  - Integration with Taskfile and Make
  - Comprehensive examples and documentation
  - GitHub Actions workflows for CI/CD

- **Icon Organization**
  - 12 predefined categories: navigation, actions, media, communication, files, ui, data, devices, social, weather, transportation, business
  - Smart categorization based on icon names
  - Category-specific generation support
  - Fallback to "misc" category for uncategorized icons

### Technical Details
- **Go Version**: Requires Go 1.21 or later
- **Dependencies**: Zero runtime dependencies
- **Output**: Three generated files per run:
  - `icons_templ.go` - Individual icon components
  - `registry_templ.go` - Type-safe icon registry
  - `categories_templ.go` - Category-based access
- **Icon Source**: Fetched from https://github.com/lucide-icons/lucide
- **Format**: SVG icons with customizable attributes

### Usage Examples
```bash
# Generate all icons
lucide-gen -output ./icons

# Generate specific categories
lucide-gen -categories "navigation,actions" -output ./ui

# Dry run preview
lucide-gen -dry-run -verbose
```

### Breaking Changes
- None (initial release)

### Migration Guide
- None (initial release)

---

## Release Notes

### v1.0.0 - Initial Release

This is the first stable release of lucide-templ-gen, providing a complete solution for generating type-safe Templ components from Lucide Icons.

**Highlights:**
- 🎯 **Type Safety**: All icon names are checked at compile time
- 🚀 **Zero Runtime Deps**: Pure Go/Templ with no JavaScript required  
- 📦 **Tree Shakeable**: Only generate the icon categories you need
- 🔄 **Auto-Updating**: Fetch latest icons directly from Lucide repository
- 🏷️ **Well Organized**: Icons grouped into logical categories
- 🛠️ **Developer Friendly**: Full integration with Go toolchain

**Categories Supported:**
- Navigation (home, menu, arrows, etc.)
- Actions (edit, save, delete, etc.)
- Media (play, pause, volume, etc.)
- Communication (mail, phone, chat, etc.)
- Files (folder, download, archive, etc.)
- UI (eye, lock, search, etc.)
- Data (database, charts, code, etc.)
- Devices (phone, laptop, watch, etc.)
- Social (heart, star, share, etc.)
- Weather (sun, rain, cloud, etc.)
- Transportation (car, plane, bike, etc.)
- Business (briefcase, wallet, chart, etc.)

Ready for production use! 🎉