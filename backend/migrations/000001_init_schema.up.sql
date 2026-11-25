-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 1. Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. Order Intents (Idempotency Protection)
CREATE TABLE order_intents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    client_order_ref VARCHAR(255) NOT NULL,
    symbol VARCHAR(50) NOT NULL,
    side VARCHAR(10) NOT NULL, -- BUY or SELL
    quantity DECIMAL(18,4) NOT NULL,
    price DECIMAL(18,4),
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_client_ref UNIQUE (user_id, client_order_ref)
);

-- 3. Orders (Actual Broker Orders)
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    intent_id UUID NOT NULL REFERENCES order_intents(id) ON DELETE RESTRICT,
    broker_order_id VARCHAR(255),
    status VARCHAR(50) NOT NULL DEFAULT 'CREATED',
    filled_quantity DECIMAL(18,4) NOT NULL DEFAULT 0,
    average_price DECIMAL(18,4),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 4. Order Events (Audit & Immutable State History)
CREATE TABLE order_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    event_type VARCHAR(50) NOT NULL, 
    old_state VARCHAR(50),
    new_state VARCHAR(50) NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 5. Reconciliation Logs (For UNKNOWN states & Background Worker)
CREATE TABLE reconciliation_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    next_check_at TIMESTAMP WITH TIME ZONE NOT NULL,
    resolve_by TIMESTAMP WITH TIME ZONE NOT NULL,
    retry_count INTEGER NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING_CHECK',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_order_intents_user ON order_intents(user_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_reconciliation_next_check ON reconciliation_logs(next_check_at) WHERE status = 'PENDING_CHECK';
-- Indexes designed for background reconciliation polling.
