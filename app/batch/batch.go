package batch

import (
	"MemoryStore/app/handlers"
	"MemoryStore/config"
	"strconv"
	"time"
)

//SaveDataBatch It saves the data in the memory according to the minute value read from the config.
func SaveDataBatch() {
	duration, _ := strconv.Atoi(config.SAVE_INTERVAL)

	ticker := time.NewTicker(time.Duration(duration) * time.Minute)

	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			handlers.SetBackup()
		}
	}
	ticker.Stop()
	done <- true
}

func StartSaveBatch() {
	go SaveDataBatch()
}
