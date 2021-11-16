package main

import (
	"MemoryStore/app"
	"MemoryStore/app/batch"
	"MemoryStore/app/handlers"
	"MemoryStore/config"
	"MemoryStore/logger"
	"net/http"
)

func main() {

	defer logger.LogFile.Close()

	/*
		config.ReadConfigFile()
		config.GetEnv("MEMDB_API_PORT", &config.API_PORT)
		config.GetEnv("LOG_FILE", &config.LOG_FILE)
		config.GetEnv("STORE_FILE", &config.STORE_FILE)
		config.GetEnv("SAVE_INTERVAL", &config.SAVE_INTERVAL)
	*/

	logger.ConfigureLogger()
	handlers.InitMemoryStore()
	handlers.ReadBackup()
	app.SetRouters()
	batch.StartSaveBatch()
	logger.Info("MemoryStore API starting...")
	logger.Info("API serving at :" + config.API_PORT)
	logger.Fatal(http.ListenAndServe(":"+config.API_PORT, nil))
}
