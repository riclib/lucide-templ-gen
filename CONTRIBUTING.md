# Contributing to lucide-templ-gen

Thank you for your interest in contributing to lucide-templ-gen! This document provides guidelines and information for contributors.

## 🚀 Quick Start

1. **Fork the repository**
2. **Clone your fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/lucide-templ-gen.git
   cd lucide-templ-gen
   ```
3. **Set up development environment**:
   ```bash
   make dev-setup
   ```
4. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## 🛠️ Development Setup

### Prerequisites

- Go 1.21 or later
- Make (optional, but recommended)
- Git

### Environment Setup

```bash
# Install dependencies and tools
make dev-setup

# Or manually:
go mod tidy
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Building

```bash
# Build the CLI tool
make build

# Or manually:
go build -o bin/lucide-gen ./cmd/lucide-gen
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Test the CLI
./bin/lucide-gen -help
./bin/lucide-gen -dry-run -verbose
```

## 📝 Code Style

### Go Code

- Follow standard Go conventions
- Use `go fmt` for formatting
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and small

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
type(scope): brief description

Optional longer description explaining the change.

Fixes #123
```

Types:
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(generator): add support for custom icon prefixes
fix(cli): handle network timeouts gracefully
docs(readme): update installation instructions
```

## 🧪 Testing Guidelines

### Unit Tests

- Write tests for all public functions
- Use table-driven tests where appropriate
- Mock external dependencies (HTTP calls, file system)
- Aim for >80% code coverage

### Integration Tests

- Test the complete generation workflow
- Test CLI functionality
- Test with real Lucide icon data

### Example Test Structure

```go
func TestGenerateIcons(t *testing.T) {
    tests := []struct {
        name     string
        config   Config
        wantErr  bool
        wantFiles []string
    }{
        {
            name: "basic generation",
            config: Config{
                OutputDir:   "/tmp/test",
                PackageName: "icons",
            },
            wantFiles: []string{
                "icons_templ.go",
                "registry_templ.go",
                "categories_templ.go",
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Generate(tt.config)
            if (err != nil) != tt.wantErr {
                t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // Additional assertions...
        })
    }
}
```

## 📖 Documentation

### Code Documentation

- Add godoc comments for all exported types and functions
- Include examples in documentation when helpful
- Keep documentation up to date with code changes

### README Updates

- Update examples if you change the API
- Add new features to the feature list
- Update installation instructions if needed

### Changelog

Update `CHANGELOG.md` with your changes:

```markdown
## [Unreleased]

### Added
- New feature description

### Changed
- Changed feature description

### Fixed
- Bug fix description
```

## 🐛 Reporting Issues

### Bug Reports

Include:
- Go version
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Error messages/logs
- Minimal reproduction case

### Feature Requests

Include:
- Clear description of the feature
- Use case and motivation
- Proposed API (if applicable)
- Examples of how it would be used

## 📦 Submitting Changes

### Pull Request Process

1. **Ensure tests pass**:
   ```bash
   make test
   make lint
   ```

2. **Update documentation** if needed

3. **Add changelog entry** for user-facing changes

4. **Create pull request** with:
   - Clear title and description
   - Reference any related issues
   - Include screenshots for UI changes
   - List any breaking changes

### Pull Request Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass locally
- [ ] Added tests for new functionality
- [ ] Manual testing performed

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review of code completed
- [ ] Documentation updated
- [ ] Changelog updated
```

## 🔄 Release Process

Releases are automated via GitHub Actions:

1. **Update version** in `cmd/lucide-gen/main.go`
2. **Update CHANGELOG.md**
3. **Create and push tag**:
   ```bash
   git tag v1.2.3
   git push origin v1.2.3
   ```
4. **GitHub Actions** will create the release automatically

## 💡 Development Tips

### Working with Icon Data

- Use small test datasets during development
- Cache icon data locally to avoid repeated API calls
- Test with different icon categories

### Debugging Generation

- Use the `-dry-run` flag to test without creating files
- Use `-verbose` for detailed output
- Test with different configurations

### Performance Considerations

- Be mindful of API rate limits
- Implement appropriate timeouts
- Use concurrent processing where beneficial

## 🤝 Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow
- Follow the project's coding standards

## 📞 Getting Help

- **Issues**: Use GitHub issues for bugs and feature requests
- **Discussions**: Use GitHub discussions for questions and ideas
- **Email**: For security issues, email maintainers directly

## 🎉 Recognition

Contributors will be:
- Listed in the project README
- Credited in release notes
- Thanked in the community

Thank you for contributing to lucide-templ-gen! 🚀