#
```markdown
# TaskFlow Go – Full-Stack Task Management App

A modern, performant task management application built with **Go** on the backend and **React + TypeScript + TailwindCSS** on the frontend. Fully containerized with Docker.

Live demo (when deployed): coming soon  
GitHub: https://github.com/yusufmalik2008/go-backend

## Features
- Create, read, update, delete tasks
- Responsive and beautiful UI with TailwindCSS
- RESTful API in Go
- Single-command startup with Docker Compose
- Hot-reloading in development (via scripts)
- Clean separation: backend → `web/`, frontend assets → built with Go script

## Tech Stack
| Layer        | Technology                                 |
|--------------|--------------------------------------------|
| Backend      | Go (net/http or your framework of choice) |
| Frontend     | React + TypeScript + TailwindCSS           |
| Build tool   | Custom Go script (`scripts/build-tailwind.go`) |
| Container    | Docker + docker-compose                    |
| Package mgr  | npm (frontend) + Go modules                |

## Project Structure
```
.
├── docker-compose.yml      → spins up backend + (optional DB)
├── go.mod                  → Go dependencies
├── Makefile                → handy shortcuts (build, run, etc.)
├── scripts/
│   └── build-tailwind.go   → compiles Tailwind → single CSS file
├── web/                    → all Go server code + templates
├── middleware/             → OpenTelemetry, auth, logging, etc.
├── package.json            → Tailwind + React dependencies
├── tailwind.config.js      → Tailwind configuration
└── README.md               → you are here
```

## Quick Start (Development)

```bash
# 1. Clone and enter
git clone https://github.com/yusufmalik2008/go-backend.git
cd go-backend

# 2. Start everything (backend + Tailwind watcher)
docker-compose up --build
# or without Docker:
make dev
```

App will be available at http://localhost:8080 (or whatever port you configured).

## Available Makefile commands
```bash
make dev        # runs air + tailwind watcher (hot reload)
make build      # builds the final binary
make run        # runs the compiled binary
make clean      # removes binaries
make tailwind   # builds CSS once
```

## Docker
```bash
# Development
docker-compose up --build

# Production (smaller image)
docker build -t taskflow-go .
docker run -p 8080:8080 taskflow-go
```

## Contributing
Feel free to open issues or PRs! This is a learning/practice project that’s turning into something real.

## License
MIT © yusufmalik2008 – see [LICENSE](LICENSE) file
```

