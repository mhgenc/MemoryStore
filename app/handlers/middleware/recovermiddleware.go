package middleware

import (
	"MemoryStore/constants"
	"MemoryStore/logger"
	"errors"
	"net/http"
)

//Recover is middleware. It's recover api panics
func Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New(constants.ErrorUnknown)
				}
				logger.Error("Recover : " + err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
