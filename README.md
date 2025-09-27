![Go CI](https://github.com/<your-username>/fullcov/actions/workflows/ci.yml/badge.svg)

# Full Test Coverage in Go ✅

This project demonstrates how to write Go code with **100% test coverage**, including:
- Unit tests
- Example tests (for documentation)
- GitHub Actions CI enforcing 100% coverage

## Run tests locally
```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
