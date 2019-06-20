package logger

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

// Traffic logs all network requests to the console
func Traffic(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		elapsedTime := time.Since(start)

		message := fmt.Sprintf(
			"%s\t%s\t%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			elapsedTime,
			r.Header.Get("X-Forwarded-For"),
			r.RemoteAddr,
		)
		log.Println(message)
		})
}

// Echo prints logs to the terminal
func Echo(message interface{}) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	log.Println(f.Name() + ": " + message.(string))
}

// Store prints logs to a file
func Store(message interface{}) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	// f := runtime.FuncForPC(pc[0])
	// Log to file
}

// Log prints logs to both
func Log(message interface{}) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	// Log to both
	log.Println(f.Name() + ": " + message.(string))
}
