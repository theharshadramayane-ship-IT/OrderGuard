# OrderGuard: Engineering Development Log

This document tracks the daily engineering progress, architectural decisions, and problem-solving approach for OrderGuard.

## [November 17, 2026] - Database Setup & Migrations
**Problem:** Needed a durable, crash-safe database schema to handle idempotency and unknown states without relying on in-memory timers.
**Solution:** 
- Selected Hosted PostgreSQL (Neon) for serverless, reliable DB hosting.
- Created tables: `users`, `order_intents`, `orders`, `order_events`, and `reconciliation_logs`.
- Implemented `next_check_at` and `resolve_by` columns for background polling, directly incorporating senior engineering feedback from Milen Hristov.
**Tools Used:** `golang-migrate`, PostgreSQL, pgx/v5.

## [November 18, 2026] - Go HTTP Server & Connection Pooling
**Problem:** Need an API layer that can safely handle concurrent requests without exhausting database connections.
**Solution:**
- Built a modular Go backend (`internal/order`, `internal/database`).
- Configured `pgxpool` for robust PostgreSQL connection pooling with a max connection limit.
- Established initial REST endpoints.

## [November 19, 2026] - Windows CGO Environment Setup
**Problem:** Integrating a Rust library into Go requires `cgo`, which needs a C compiler on Windows. The MSVC linker was failing.
**Solution:**
- Installed MSYS2 and GCC (UCRT64) to provide a GNU-based C toolchain.
- Configured Rust to use the `stable-x86_64-pc-windows-gnu` toolchain to match the GCC compiler.
- Resolved disk space issues (Disk Cleanup) to allow the Rust compiler to install.

## [November 25, 2026] - Rust Safety Library & FFI Integration
**Problem:** Need a memory-safe, deterministic rule engine to evaluate duplicate orders and invalid quantities.
**Solution:**
- Built a Rust `cdylib` library (`orderguard-safety`).
- Implemented a C-ABI compatible FFI function (`evaluate_order_safety`).
- Protected the FFI boundary using `std::panic::catch_unwind` to prevent Rust panics from crashing the Go server.
- Linked the compiled `.dll` to the Go backend via `cgo`.
**Result:** Go successfully calls Rust for order safety evaluation.

## [November 26, 2026] - Refactoring Go Modules & CI Readiness
**Problem:** Go module paths were failing because `go.mod` was initialized in the wrong directory, causing `orderguard/backend/internal/database` to be searched in the system's Go path.
**Solution:**
- Consolidated the project into a single root module (`go mod init orderguard`).
- Fixed the import paths in `main.go`.
- Prepared the repository for CI/CD pipelines.

## Next Steps
- Write Go API endpoints that call the Rust safety engine before DB insertion.
- Implement database-driven background poller using `FOR UPDATE SKIP LOCKED` (as advised by Milen Hristov).
- Connect React frontend.