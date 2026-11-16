# OrderGuard Roadmap

## Phase 0 — Environment & Requirements Audit
Verify tools; freeze MVP scope; pick hosted DB, migration tool, DB driver, Rust build mode; formalize idempotency scope.

## Phase 1 — Architecture & Design
- Product Requirements Document (PRD)
- System architecture
- Database ERD
- Order state machine
- Go–Rust FFI contract
- API spec
- Security design
- Failure scenario matrix
- Testing strategy

## Phase 2 — Rust Safety Library + FFI
Deterministic rules; C-ABI; unit tests; Go integration tests.

## Phase 3 — Go Backend
Config, API, models, validation, DB, Rust integration, order lifecycle.

## Phase 4 — Idempotency + Reconciliation
Durable idempotency records; concurrency protection; event history; UNKNOWN handling; reconciliation engine.

## Phase 5 — Simulated Broker
Accept, reject, delay, partial fill, full fill, cancel, network failure, lost response, duplicate events.

## Phase 6 — Frontend
Dashboard, confirmation modal, processing/unknown indicators, duplicate warning, history, status updates.

## Phase 7 — Testing & Reliability
Unit, integration, concurrency, failure injection, restart recovery, consistency, security, load.

## Phase 8 — Documentation & Demo
README, architecture, API, setup, test report, limitations, demo scenarios, presentation.

## Phase 9 — Future Broker Integration
Authorized broker adapters only; sandbox/paper first; security & compliance review.