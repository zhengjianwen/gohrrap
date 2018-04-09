package middleware

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger inherits from log.Logger used to log messages with the Logger middleware
	*log.Logger
}

// NewLogger returns a new Logger instance
func NewLogger(out io.Writer) *Logger {
	l := log.New(out, "", 0)
	l.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return &Logger{l}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, r)
	use := time.Since(start)
	if use > time.Second {
		res := rw.(negroni.ResponseWriter)
		l.Printf("%s %v %s in %v", r.Method, res.Status(), r.URL.Path, use)
	}
}
