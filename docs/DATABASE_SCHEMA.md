# OrderGuard Database Schema Design (MVP)

## 1. users
Stores trader information.
- `id` (UUID, Primary Key)
- `email` (VARCHAR, Unique)
- `password_hash` (VARCHAR)
- `created_at` (TIMESTAMP)

## 2. order_intents
Stores the user's initial request (First Line of Defense against Duplicate Clicks).
- `id` (UUID, Primary Key)
- `user_id` (UUID, FK to users)
- `client_order_ref` (VARCHAR) -- The Idempotency Key
- `symbol` (VARCHAR)
- `side` (VARCHAR) -- BUY or SELL
- `quantity` (DECIMAL)
- `price` (DECIMAL)
- `status` (VARCHAR) -- PENDING, PROCESSED, DUPLICATE_DETECTED
- `created_at` (TIMESTAMP)
- **UNIQUE CONSTRAINT:** `(user_id, client_order_ref)` (Prevents duplicate intents)

## 3. orders
The actual order sent to the broker.
- `id` (UUID, Primary Key)
- `intent_id` (UUID, FK to order_intents)
- `broker_order_id` (VARCHAR, Nullable) -- Returned by Broker
- `status` (VARCHAR) -- CREATED, SUBMITTING, ACKNOWLEDGED, OPEN, FILLED, UNKNOWN, REJECTED
- `filled_quantity` (DECIMAL, Default 0)
- `average_price` (DECIMAL)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

## 4. order_events
Immutable audit log of everything that happens to an order.
- `id` (UUID, Primary Key)
- `order_id` (UUID, FK to orders)
- `event_type` (VARCHAR) -- e.g., STATE_TRANSITION, BROKER_RESPONSE, TIMEOUT
- `old_state` (VARCHAR)
- `new_state` (VARCHAR)
- `metadata` (JSONB) -- Raw broker payload for debugging
- `created_at` (TIMESTAMP)

## 5. reconciliation_logs
(Senior Engineer's Advice Implementation)
Handles UNKNOWN states and background polling.
- `id` (UUID, Primary Key)
- `order_id` (UUID, FK to orders)
- `next_check_at` (TIMESTAMP) -- The Go worker polls this
- `resolve_by` (TIMESTAMP) -- Hard deadline. If missed, escalate to human.
- `retry_count` (INTEGER)
- `status` (VARCHAR) -- PENDING_CHECK, RESOLVED, ESCALATED