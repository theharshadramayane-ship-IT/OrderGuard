# OrderGuard — Known Limitations

OrderGuard is a prototype for engineering and demonstration purposes.

## Not a certified trading system
- Not audited or certified for real-money trading.
- Do not use in production without independent security and compliance review.

## No live broker integration in the MVP
- OrderGuard connects only to a simulated broker.
- Live broker adapters are Phase 9 (future), require authorized APIs, and are out of MVP scope.

## Idempotency is not exactly-once execution
- An idempotency key prevents duplicate *submissions* from the client, on the OrderGuard side.
- It cannot guarantee that an external broker will not execute an order twice across networks.
- Design accepts *uncertain external outcomes* and reconciles them.

## UNKNOWN ≠ REJECTED
- An UNKNOWN order state means the outcome is not yet known.
- OrderGuard never auto-retries an UNKNOWN order; it reconciles first.

## Windows cgo constraint
- Go's cgo requires GCC/clang on Windows (not MSVC).
- `CGO_ENABLED` must be `1` and a GCC-compatible compiler installed before the Rust library can be linked.