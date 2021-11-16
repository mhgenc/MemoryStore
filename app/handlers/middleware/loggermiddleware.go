package middleware

import (
	"MemoryStore/logger"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//loggingResponseWriter http.ResponseWriter is an interface
type loggingResponseWriter struct {
	status int
	body   string
	http.ResponseWriter
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *loggingResponseWriter) Write(body []byte) (int, error) {
	w.body = string(body)
	return w.ResponseWriter.Write(body)
}

//Logger is a middleware. Logs request and response
func Logger(h http.Handler) http.Handler {
	var logMessage string
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingRW := &loggingResponseWriter{
			ResponseWriter: w,
		}
		body, _ := ioutil.ReadAll(r.Body)

		logMessage = fmt.Sprintf("METHOD=%s, URI=%s, BODY=%s", r.Method, r.RequestURI, body)
		logger.Info(logMessage)

		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		h.ServeHTTP(loggingRW, r)

		logMessage = fmt.Sprintf("STATUS=%d, RESPONSE=%s", loggingRW.status, loggingRW.body)
		logger.Info(logMessage)
	})
}
