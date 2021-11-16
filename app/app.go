package app

import (
	"MemoryStore/app/handlers"
	"MemoryStore/app/handlers/middleware"
	"net/http"
)

//SetRouters Set api routes based on radix tree
func SetRouters() {

	http.Handle("/api/set", middleware.Recover(middleware.Logger(http.HandlerFunc(handlers.SetKey))))
	http.Handle("/api/get", middleware.Recover(middleware.Logger(http.HandlerFunc(handlers.GetVal))))
	http.Handle("/api/flush", middleware.Recover(middleware.Logger(http.HandlerFunc(handlers.Flush))))
}
