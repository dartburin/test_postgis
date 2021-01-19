package middleware

import (
	"net/http"
	"time"

	lg "github.com/sirupsen/logrus"
)

type LogData struct {
	Log  lg.FieldLogger
	Name string
}

// Logger middleware function to handle logging REST call
func (l *LogData) MidLogger(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		// Calls the handler
		handler.ServeHTTP(w, r)
		// Logging
		l.Log.Printf("%s call: %s   [data: %v] [time: %s]\n", l.Name, r.Method, r.Body, time.Since(start))
	})
}
