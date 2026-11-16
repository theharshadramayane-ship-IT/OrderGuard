# OrderGuard — Local Setup

## Prerequisites (verified in Phase 0)

- Windows 10 22H2
- Go 1.27.4 (windows/amd64)
- Rust 1.99.1 (stable-x86_64-pc-windows-msvc)
- Cargo 1.99.1
- Node 22.23.2
- npm 10.9.8
- Git 2.56.0
- Docker CLI 29.7.2 (daemon currently stopped — not required for Phase 0/1)
- **No local PostgreSQL.** All DB access via hosted `DATABASE_URL`.
- **cgo currently disabled (`CGO_ENABLED=0`); a GCC-compatible compiler must be installed before Phase 2.**

## Folder layout
orderguard/
├── backend/ # Go backend
├── frontend/ # React + TypeScript
├── rust/ # orderguard-safety (cdylib)
├── docs/ # Design & operations docs
├── scripts/ # Helper scripts
├── PROGRESS.md
└── README.md


## Environment variables

Copy `.env.example` → `.env` and fill in:

- `DATABASE_URL` — hosted PostgreSQL connection string (TLS required)
- `JWT_SECRET` — dev-only signing secret
- `SIMULATED_BROKER_*` — see Phase 5

**Never commit `.env`.**