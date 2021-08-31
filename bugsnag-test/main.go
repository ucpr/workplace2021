package main

import (
	"fmt"
	"net/http"

	"github.com/bugsnag/bugsnag-go/v2"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(status int) {
	s.status = status
	s.ResponseWriter.WriteHeader(status)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		wRecorder := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(wRecorder, r)
		fmt.Println(wRecorder.status)
		fmt.Println("[END] middleware1")
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world"))
}

func panicFunc(w http.ResponseWriter, req *http.Request) {
	panic("パニックパニック！")
}

func resp500(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println("internal error")
	w.Write([]byte("Internal error"))
}

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "",
		ReleaseStage: "development",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/ucpr/workspace2021/bugsnag-test"},
		// more configuration options
	})

	// rest of your program.
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(index))
	mux.Handle("/panic", middleware(http.HandlerFunc(panicFunc)))
	mux.Handle("/500", middleware(http.HandlerFunc(resp500)))

	srv := http.Server{
		Addr:    ":8080",
		Handler: bugsnag.Handler(mux),
	}

	srv.ListenAndServe()
}
