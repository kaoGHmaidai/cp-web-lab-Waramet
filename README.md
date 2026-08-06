# CP Web Application Starter Template

This repository contains a full-stack starter with a Go Fiber backend and a Next.js frontend. The backend is organized around dependency-injected services and router-scoped endpoint packages, while the frontend uses Next.js App Router.

## Repository structure

```text
.
├── docker-compose.yml
├── README.md
├── src/
│   ├── backend/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── main_test.go
│   │   └── internal/
│   │       ├── endpoints/
│   │       │   ├── healthz/
│   │       │   │   ├── endpoint.go
│   │       │   │   └── router.go
│   │       │   └── foobar/
│   │       │       ├── endpoint.go
│   │       │       └── router.go
│   │       ├── service/
│   │       │   ├── service.go
│   │       │   ├── service_mock.go
│   │       │   └── service_test.go
│   └── frontend/
│       ├── app/
│       ├── components/
│       ├── package.json
│       └── ...
└── .github/workflows/
    ├── ci.yml
    ├── backend-quality-gate.yml
    ├── backend-tests.yml
    ├── frontend-quality-gate.yml
    └── frontend-tests.yml
```

## Prerequisites

Install the following before running the project locally:

- Go 1.26.5 or newer
- Node.js 20+ (CI currently uses Node 24)
- npm

## Start the backend (Go)

From the repository root:

```bash
cd src/backend
go mod download
go run .
```

The backend listens on port 3000 and exposes:

- GET /healthz
- GET /foo/bar

You can verify it by opening http://localhost:3000/healthz.

### Backend tests

```bash
cd src/backend
go test ./...
```

## Start the frontend (Next.js)

Open a second terminal and run:

```bash
cd src/frontend
npm ci
npm run dev -- --port 3001
```

Then open http://localhost:3001 in your browser.

### Frontend tests and checks

```bash
cd src/frontend
npm run lint
npm run test:coverage
```

## Run everything with Docker Compose

The repository also includes a Docker Compose setup in [docker-compose.yml](docker-compose.yml) for running both services together.

From the repository root:

```bash
docker compose up --build
```

This starts:

- Backend on http://localhost:3000
- Frontend on http://localhost:3001

To stop the services:

```bash
docker compose down
```

To remove the volumes as well:

```bash
docker compose down -v
```

## CI overview

GitHub Actions runs the checks from [.github/workflows/ci.yml](.github/workflows/ci.yml) on pushes and pull requests.

### What runs in CI

- Pre-checks: formatting and shared setup from [.github/workflows/pre.yml](.github/workflows/pre.yml)
- Frontend quality gates: linting, documentation checks, duplicate-code detection, dependency checks, and production-only audit from [.github/workflows/frontend-quality-gate.yml](.github/workflows/frontend-quality-gate.yml)
- Frontend tests: Jest unit tests with coverage thresholds from [.github/workflows/frontend-tests.yml](.github/workflows/frontend-tests.yml)
- Backend quality gates: formatting, vet, static analysis, and vulnerability scanning from [.github/workflows/backend-quality-gate.yml](.github/workflows/backend-quality-gate.yml)
- Backend tests: Go test coverage from [.github/workflows/backend-tests.yml](.github/workflows/backend-tests.yml)

### Current thresholds

- Backend coverage threshold: 80%
- Frontend coverage thresholds:
  - lines: 80%
  - functions: 80%
  - branches: 70%
  - statements: 80%

If you are contributing, make sure the relevant local checks pass before opening a pull request.

## For maintainer
When open MR, the CI will also check the MR whether it's matched with the convention or not. Please refer to this convention when naming the MR (Referenced convention)[https://commitlint.js.org/concepts/commit-conventions.html]

## Optional pre-commit setup

You can install a Git pre-commit hook locally so common checks run automatically before each commit.

```bash
pip install pre-commit
pre-commit install
```

What the hook helps with:

- `go-fmt` and `go-vet` for the backend
- `eslint` and related frontend quality checks
- basic formatting and whitespace hygiene
- catching obvious issues before they reach CI

If you want to run the hooks manually at any time, use:

```bash
pre-commit run --all-files
```

To clean up later, you can remove the hook installation and undo any generated artifacts:

```bash
pre-commit uninstall
rm -rf .git/hooks/pre-commit
```

If you also want to remove local build and cache output created during development, you can delete the generated directories:

```bash
rm -rf src/frontend/.next src/frontend/coverage src/backend/coverage
```