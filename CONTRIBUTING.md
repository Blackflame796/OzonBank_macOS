# Contributing Guidelines

## Development Setup

1. Ensure you have Go 1.26.1 or later installed
2. Clone the repository
3. Install dependencies: `go mod download`

## Project Structure

- `cmd/ozonbank/`: Application entry point
- `internal/app/`: Core application logic
- `internal/ui/`: UI components
- `internal/models/`: Data structures
- `internal/utils/`: Utility functions
- `scripts/`: Build and utility scripts
- `tests/`: Integration tests

## Code Style

- Follow Go conventions and best practices
- Use meaningful variable and function names
- Add comments for exported types and functions
- Keep functions small and focused on a single responsibility

## Testing

Before submitting a PR, ensure:
1. Code compiles without errors
2. All tests pass
3. No obvious issues with the code

## Commit Messages

Use clear and descriptive commit messages:
- `feat: add new feature description`
- `fix: fix bug description`
- `docs: update documentation`
- `refactor: refactor component name`

## Pull Request Process

1. Create a feature branch from `main`
2. Make your changes
3. Ensure code follows project conventions
4. Submit a pull request with a clear description
5. Respond to review comments

## Adding New Features

When adding new features:
1. Add new types/structures to `internal/models/`
2. Implement logic in appropriate `internal/` package
3. Add logging using the `Logger` utility
4. Update documentation in `docs/` or `README.md`
5. Add tests if applicable

## Questions or Issues

If you have questions about the architecture or need clarification, please open an issue or contact the maintainers.
