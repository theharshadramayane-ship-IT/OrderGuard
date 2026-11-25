use std::ffi::CStr;
use std::os::raw::c_char;

/// Evaluates if an order is safe to submit.
/// Returns:
///  1  = Safe (Allow)
///  0  = Unsafe (Reject)
/// -1  = Invalid Arguments (Null pointer)
/// -99 = Internal Panic
#[no_mangle]
pub extern "C" fn evaluate_order_safety(
    client_order_ref: *const c_char,
    quantity: f64,
) -> i32 {
    // Protect the FFI boundary from Rust panics
    let result = std::panic::catch_unwind(|| {
        if client_order_ref.is_null() {
            return -1;
        }

        let c_str = unsafe { CStr::from_ptr(client_order_ref) };
        let ref_str = match c_str.to_str() {
            Ok(s) => s,
            Err(_) => return -1,
        };

        // Rule 1: Idempotency key cannot be empty
        if ref_str.trim().is_empty() {
            return 0; 
        }

        // Rule 2: Quantity must be strictly positive
        if quantity <= 0.0 {
            return 0; 
        }

        // Rule 3: Maximum allowed single order limit (Example: 1000 shares)
        if quantity > 1000.0 {
            return 0;
        }

        1 // Safe
    });

    match result {
        Ok(code) => code,
        Err(_) => -99,
    }
}