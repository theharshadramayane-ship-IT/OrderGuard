# OrderGuard — Progress Tracker

Last updated: Phase 0, Step 0.3

## Status Legend
- NOT STARTED
- IN PROGRESS
- BLOCKED
- TESTING
- VERIFIED COMPLETE

## Phase Overview

| Phase | Name | Status |
|---|---|---|
| 0 | Environment & Requirements Audit | IN PROGRESS |
| 1 | Architecture & Design Docs | NOT STARTED |
| 2 | Rust Safety Library + FFI | NOT STARTED |
| 3 | Go Backend | NOT STARTED |
| 4 | Idempotency + Reconciliation | NOT STARTED |
| 5 | Simulated Broker | NOT STARTED |
| 6 | Frontend | NOT STARTED |
| 7 | Testing & Reliability | NOT STARTED |
| 8 | Documentation & Demo | NOT STARTED |
| 9 | Future Broker Integration | NOT STARTED |

## Phase 0 — Detailed Checklist

| Item | Status | Notes |
|---|---|---|
| OS / tool versions verified | VERIFIED COMPLETE | Go 1.27.4, Rust 1.99.1 MSVC, Node 22.23.2, Git 2.56.0, Docker CLI 29.7.2 |
| Project folder chosen | VERIFIED COMPLETE | `C:\dev\orderguard` |
| GitHub repo created | VERIFIED COMPLETE | `theharshadramayane-ship-IT/OrderGuard` |
| `.gitignore` (secrets) | IN PROGRESS | This step |
| Initial docs skeleton | IN PROGRESS | This step |
| C toolchain for cgo | BLOCKED (deferred to Phase 2 prep) | `CGO_ENABLED=0`, no gcc/clang |
| Hosted PostgreSQL provider selected | NOT STARTED | Recommendation: Neon |
| Migration tool decided | NOT STARTED | Recommendation: golang-migrate |
| Go DB driver decided | NOT STARTED | Recommendation: pgx/v5 |
| Rust build mode decided | NOT STARTED | Recommendation: cdylib |
| Idempotency scope decided | NOT STARTED | Recommendation: (user_id, client_order_ref) |
| Redis in MVP? | NOT STARTED | Recommendation: OUT |
| MVP scope frozen | NOT STARTED | Awaiting decisions above |

## Known Risks

1. **cgo on Windows needs GCC-compatible compiler.** Go's cgo does NOT support MSVC. Plan: MSYS2 + MinGW-w64 at Phase 2 start.
2. **Docker daemon currently stopped.** Not needed until Phase 5.
3. **No local PostgreSQL.** All DB work via hosted provider + `DATABASE_URL`.

## Last Action
- GitHub repo created; cloning and bootstrapping docs.
