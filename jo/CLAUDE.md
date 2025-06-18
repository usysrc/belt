# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands
- Build: `go build`
- Run: `go run main.go [args]`
- Install: `go install`
- Test: `go test ./...`
- Test single file: `go test -v ./path/to/file_test.go`
- Lint: `golint ./...` (requires golint: `go install golang.org/x/lint/golint@latest`)
- Format: `gofmt -w .`

## Code Style Guidelines
- **Formatting**: Follow standard Go conventions with `gofmt`
- **Imports**: Group standard library imports first, then third-party, then local packages
- **Error Handling**: Always check errors and provide context in error messages
- **Comments**: Document exported functions and packages with meaningful comments
- **Naming**: Use camelCase for variables, PascalCase for exported functions/types
- **Types**: Prefer explicit types over interface{} when possible
- **Function Size**: Keep functions small and focused on a single responsibility
- **Error Propagation**: Use `fmt.Errorf("context: %w", err)` for error wrapping

## Version Control Workflow
This project uses Jujutsu (jj) for version control. When the user asks to "jj describe and push" or similar commit/push requests:

1. **Describe the change**: `jj describe -m "commit message"`
2. **Set main bookmark**: `jj bookmark set main` 
3. **Push to remote**: `jj git push --bookmark main`

**Commit Message Format**: Use conventional commit style:
- `feat[scope]: description` for new features
- `fix[scope]: description` for bug fixes  
- `docs[scope]: description` for documentation
- `refactor[scope]: description` for code refactoring
- `test[scope]: description` for test changes

**Example Usage**: When user says "jj describe and push", create a descriptive commit message based on the changes made, then execute the three-step workflow above.