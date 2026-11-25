<div align="center">
  <h1>??? OrderGuard</h1>
  <p><b>Every order deserves certainty.</b></p>

  <!-- Badges -->
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version"></a>
  <a href="https://www.rust-lang.org/"><img src="https://img.shields.io/badge/Rust-1.75+-000000?style=flat-square&logo=rust" alt="Rust Version"></a>
  <a href="https://www.postgresql.org/"><img src="https://img.shields.io/badge/PostgreSQL-15+-316192?style=flat-square&logo=postgresql" alt="PostgreSQL"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-green.svg?style=flat-square" alt="License"></a>
</div>

---

## ?? The Problem
In high-frequency trading, network timeouts create "Uncertain States." If a trader clicks "BUY" twice, and the broker network drops the first response, the trader risks duplicate financial exposure. Most systems treat a timeout as a failure, leading to accidental double orders.

## ??? The Solution
**OrderGuard** is a production-minded trading reliability and protection infrastructure. It uses a modular architecture to evaluate order safety, ensure durable idempotency, and reconcile uncertain states.

### Core Architecture
\\\	ext
[ Frontend (React/TS) ]
         ¦  (REST + Idempotency Key)
         ?
[ Go Backend (Orchestrator) ]
         ¦
         +--> [ PostgreSQL (Hosted) ] -- Durable state & Idempotency Records
         ¦
         +--> [ Rust Safety Library (cgo) ]
                  +-- Deterministic Duplicate Detection
\\\`n
## ?? Key Features
- **Idempotency Protection:** Database-enforced unique constraints prevent duplicate requests.
- **State Machine:** Explicit lifecycle (CREATED ? SUBMITTING ? UNKNOWN ? FILLED/REJECTED).
- **Zero-Trust FFI Boundary:** Go backend offloads deterministic safety rules to a memory-safe Rust library via C-ABI.
- **Background Reconciliation:** Database-driven deadlines (`next_check_at`) prevent stuck orders.

## ??? Tech Stack
- **Backend:** Go (Golang), `net/http`, `pgx/v5`
- **Safety Engine:** Rust (`cdylib`, C-ABI)
- **Database:** PostgreSQL (Neon)
- **Migrations:** `golang-migrate`

## ?? Local Development

### Prerequisites
- Go 1.21+
- Rust (stable)
- C Compiler (GCC/MinGW for Windows cgo)
- Hosted PostgreSQL (e.g., Neon, Supabase)

### Setup Instructions
1. Clone the repository
2. Create `.env` in the root and add your `DATABASE_URL`
3. Run database migrations:
   \\\ash
   migrate -path backend/migrations -database "" up
   \\\`n4. Build and run the backend:
   \\\ash
   cd backend
   go run main.go
   \\\`n
## ?? Documentation
- [Database Schema](docs/DATABASE_SCHEMA.md)
- [Architecture Design](assets/architecture.md)

---
*Disclaimer: This is an early-stage prototype and not a certified trading system. Do not use in production without independent security review.*

- [Development Log](DEV_LOG.md)

## Architecture Decisions
See DEV_LOG.md for a detailed engineering log.

- [x] Rust FFI integrated
