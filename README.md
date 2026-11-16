# OrderGuard

> Every order deserves certainty.

OrderGuard is a trading order reliability and protection infrastructure. It protects traders from accidental duplicate orders and unclear order execution states caused by repeated clicks, delayed broker updates, network timeouts, and unsafe API retries.

**Status:** Early-stage prototype. Not a certified trading system. Not connected to any live broker.

---

## Architecture (planned)
Frontend (React+TS) → Go Backend → Rust Safety Library (FFI)
│
├── PostgreSQL (hosted, durable source of truth)
└── Broker Adapter → Simulated Broker


- **Go**: API, orchestration, persistence, broker communication.
- **Rust**: Deterministic safety rules via a small C-ABI FFI. No I/O, no DB.
- **PostgreSQL**: Durable idempotency, order events, audit trail.

---

## Project phases

See [`docs/ROADMAP.md`](docs/ROADMAP.md).

---

## Development

See [`docs/SETUP.md`](docs/SETUP.md).

---

## Known limitations & disclaimers

See [`docs/LIMITATIONS.md`](docs/LIMITATIONS.md).
