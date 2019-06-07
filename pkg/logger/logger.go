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

func FileLog(object interface{}) error {
	return nil
}

// Log outputs an object to the console
func Log(objects ...interface{}) {
	depth := 1

	pc, _, line, _ := runtime.Caller(depth)

	fn := runtime.FuncForPC(pc)
	var functionName string
	if fn == nil {
		functionName = "?()"
	} else {
		functionName = fn.Name()
	}

	log.Println(functionName, line, ":", objects)
}
