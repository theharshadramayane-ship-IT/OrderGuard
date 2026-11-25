# OrderGuard API Design

## POST /api/orders
- **Description:** Submits a new order.
- **Payload:** client_order_ref, symbol, side, quantity, price.
- **Safety:** Evaluated by Rust FFI before DB insertion.
