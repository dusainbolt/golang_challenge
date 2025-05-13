package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ==== 1. Handler Function (Getting Started Handler) ====
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "👋 Xin chào từ handler cơ bản!")
}

// ==== 2. ServerMux ====
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📄 Đây là trang /about")
}

// ==== 3. Custom Handler ====
type CustomHandler struct{}

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🤖 Đây là custom handler bằng struct")
}

// ==== 4. Middleware ====
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[LOG] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("[DONE] Mất %s\n", time.Since(start))
	})
}

// ==== 5. Multiple Servers ====
func startAPIServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "📡 Pong từ API server")
	})
	log.Println("🚀 API server chạy ở :9000")
	log.Fatal(http.ListenAndServe(":9000", mux))
}

func main() {
	// Mux chính
	mux := http.NewServeMux()

	// 1. Handler cơ bản
	mux.HandleFunc("/", helloHandler)

	// 2. ServerMux route khác
	mux.HandleFunc("/about", aboutHandler)

	// 3. Custom handler bằng struct
	mux.Handle("/custom", CustomHandler{})

	// 4. Gắn middleware
	loggedMux := loggingMiddleware(mux)

	// 5. Chạy server phụ
	go startAPIServer()

	// 6. Server chính
	log.Println("🌐 Server chính chạy ở :8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
