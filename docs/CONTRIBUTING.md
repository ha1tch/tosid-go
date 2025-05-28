# Contributing to Semantic Computing with TOSID and KMAC

Thank you for your interest in contributing to the TOSID and KMAC project! This document provides guidelines for contributing to the codebase.

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, but recommended)

### Setting up the development environment

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/semantic-computing.git
   cd semantic-computing
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Install development tools:
   ```bash
   make install-tools
   ```

4. Run tests to verify setup:
   ```bash
   make test
   ```

## Development Workflow

### Before making changes

1. Create a new branch for your feature/bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make sure all tests pass:
   ```bash
   make check
   ```

### Making changes

1. Write your code following Go conventions
2. Add appropriate tests for new functionality
3. Update documentation if needed
4. Run the full test suite:
   ```bash
   make test
   ```

5. Run linting:
   ```bash
   make lint
   ```

6. Format your code:
   ```bash
   make fmt
   ```

### Testing

We use Go's built-in testing framework. Tests should be:

- Comprehensive (test happy path, edge cases, and error conditions)
- Fast (use mocks/stubs for external dependencies)
- Deterministic (no random failures)
- Well-named (test function names should describe what they test)

#### Running tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make benchmark

# Run tests for a specific package
go test ./pkg/tosid -v
```

### Code Style

We follow standard Go conventions:

- Use `gofmt` for formatting
- Follow the guidelines in [Effective Go](https://golang.org/doc/effective_go.html)
- Use meaningful variable and function names
- Write comments for exported functions and types
- Keep functions focused and reasonably sized

### Commit Messages

Use conventional commit format:

```
type(scope): description

body (optional)
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(tosid): add hierarchical navigation methods
fix(kmac): resolve assertion validation bug
docs(readme): update installation instructions
```

## Project Structure

```
semantic-computing/
├── cmd/examples/        # Example applications
├── internal/           # Internal packages (not exported)
│   ├── kmac/          # KMAC implementation
│   ├── tosid/         # TOSID implementation
│   └── integration/   # Integration utilities
├── pkg/               # Public API packages
│   ├── kmac/         # Public KMAC interface
│   ├── tosid/        # Public TOSID interface
│   └── semantic/     # Combined semantic API
└── docs/             # Documentation
```

## Adding New Features

When adding new features:

1. **Start with the public API** (`pkg/` packages) to define the interface
2. **Implement in internal packages** for the actual logic
3. **Add comprehensive tests** for both happy path and error cases
4. **Update documentation** including README and code comments
5. **Add examples** if the feature would benefit from demonstration

### TOSID Extensions

When extending TOSID functionality:

- Maintain backward compatibility with existing TOSID codes
- Follow the hierarchical taxonomy structure
- Add appropriate validation for new taxonomy codes
- Update the taxonomy documentation

### KMAC Extensions

When extending KMAC functionality:

- Ensure new statement types follow the existing patterns
- Add appropriate validation
- Consider how new statements integrate with the disassembler
- Update examples to demonstrate new capabilities

## Pull Request Process

1. **Create a pull request** with a clear title and description
2. **Reference any related issues** using GitHub's linking syntax
3. **Ensure CI passes** (tests, linting, etc.)
4. **Request review** from maintainers
5. **Address feedback** and make necessary changes
6. **Squash commits** if requested before merging

### Pull Request Checklist

- [ ] Tests pass locally
- [ ] New code has appropriate test coverage
- [ ] Documentation is updated
- [ ] Code follows project style guidelines
- [ ] Commit messages follow conventional format
- [ ] No breaking changes (or clearly documented)

## Reporting Issues

When reporting bugs or requesting features:

1. **Check existing issues** to avoid duplicates
2. **Use issue templates** when available
3. **Provide clear reproduction steps** for bugs
4. **Include relevant system information** (Go version, OS, etc.)
5. **Add labels** to categorize the issue

## Questions and Support

For questions about development or usage:

1. Check the documentation first
2. Search existing issues and discussions
3. Create a new discussion or issue if needed
4. Be specific about what you're trying to achieve

## Code of Conduct

This project follows the [Go Community Code of Conduct](https://golang.org/conduct). Please be respectful and inclusive in all interactions.

## License

By contributing to this project, you agree that your contributions will be licensed under the Apache License 2.0.