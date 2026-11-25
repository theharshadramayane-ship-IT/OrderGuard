package main

/*
#cgo CFLAGS: -I../rust/orderguard-safety/src
#cgo LDFLAGS: -L../rust/orderguard-safety/target/release -lorderguard_safety
#include <stdlib.h>

extern int evaluate_order_safety(const char* client_order_ref, double quantity);
*/
import "C"
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"unsafe"

	"orderguard/backend/internal/database" 
)

func main() {
	// 1. Connect to Database
	dbPool := database.ConnectDB()
	defer dbPool.Close()

	// 2. Define API Handler
	http.HandleFunc("/api/test-safety", func(w http.ResponseWriter, r *http.Request) {
		ref := "req-12345"
		qty := 500.0 // Safe quantity

		// Convert Go string to C string (Rust ke liye)
		cRef := C.CString(ref)
		defer C.free(unsafe.Pointer(cRef))

		// CALL RUST FUNCTION
		result := C.evaluate_order_safety(cRef, C.double(qty))

		fmt.Printf("Rust evaluated order ref=%s qty=%.2f => Result: %d\n", ref, qty, result)

		if result == 1 {
			w.Write([]byte("Rust says: SAFE (1)"))
		} else if result == 0 {
			w.Write([]byte("Rust says: UNSAFE (0)"))
		} else {
			w.Write([]byte("Rust says: ERROR"))
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 OrderGuard Backend starting on http://localhost:%s", port)
	log.Printf("Test Endpoint: http://localhost:%s/api/test-safety", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
