package handlers

import (
	"MemoryStore/app/models"
	"MemoryStore/config"
	"MemoryStore/constants"
	"MemoryStore/logger"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"time"
)

//MemoryStore is our db
var memoryStore map[string]string

//InitMemoryStore Initialize memory store map[string]string
func InitMemoryStore() {
	memoryStore = make(map[string]string)
}

//FlushMemoryStore Clear memory store
func FlushMemoryStore() {
	InitMemoryStore()
}

//SetStoreData Set key/value to memory store
func SetStoreData(data *models.ApiData) {
	memoryStore[data.Key] = data.Value
}

//GetStoreData Returns the value of the key given as a parameter
func GetStoreData(key string) (string, error) {
	if val, ok := memoryStore[key]; ok {
		return val, nil
	} else {
		return "", errors.New(constants.ErrorNoDataFound)
	}
}

//PrepareResponse Prepare the api response body and http status
func PrepareResponse(w http.ResponseWriter, status int, errorMessage string, data *models.ApiData) {
	w.Header().Set("Content-Type", "application/json")

	response := models.SetApiResponse(status, errorMessage, data)

	w.WriteHeader(response.GetStatus())
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		logger.Error("PrepareResponse - ", err)
	}
}

// ReadBackup read backup file and fill MemoryStore
func ReadBackup() {

	err := os.MkdirAll(config.DIR_SAVED_FILES, os.ModePerm)
	if err != nil {
		logger.Error("ReadBackUp-", err)
		return
	}

	files, err := ioutil.ReadDir(config.DIR_SAVED_FILES)

	if err != nil {
		logger.Error("ReadBackUp - ", err)
		return
	}

	// if backup file exists then read and save to memorystore
	if len(files) > 1 {

		//order by created date
		sort.Slice(files, func(i, j int) bool {
			return (files[i].ModTime().Before(files[j].ModTime()))
		})

		//find most new file
		filePath := path.Join(config.DIR_SAVED_FILES, files[len(files)-1].Name())

		f, err := ioutil.ReadFile(filePath)
		if err != nil {
			logger.Fatal("ReadBackUp ", err)
		}

		err = json.Unmarshal(f, &memoryStore)
		if err != nil {
			logger.Error("ReadBackUp ", err)
		} else {
			logger.Info("Store db read from " + filePath)
		}
	}
}

//ArchiveBackups Archives backup files when the flush api is called
func ArchiveBackups() error {
	// if dir exists do nothing, else create dir
	err := os.MkdirAll(config.DIR_SAVED_FILES, os.ModePerm)
	if err != nil {
		logger.Error("ArchiveBackups-", config.DIR_SAVED_FILES, err)
		return err
	}

	err = os.MkdirAll(config.DIR_ARCHIVE_FILES, os.ModePerm)
	if err != nil {
		logger.Error("ArchiveBackups-", config.DIR_ARCHIVE_FILES, err)
		return err
	}

	files, err := ioutil.ReadDir(config.DIR_SAVED_FILES)

	if err != nil {
		logger.Error("archiveBackups - " + err.Error())
		return err
	}

	for i := range files {
		if !files[i].IsDir() {
			oldFilePath := path.Join(config.DIR_SAVED_FILES, files[i].Name())
			newFilePath := path.Join(config.DIR_ARCHIVE_FILES, files[i].Name())
			err := os.Rename(oldFilePath, newFilePath)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

// SetBackup Saves the current version of the memorystore to the backup file
func SetBackup() {

	err := os.MkdirAll(config.DIR_SAVED_FILES, os.ModePerm)
	if err != nil {
		logger.Error(err)
		return
	}

	if len(memoryStore) > 0 {
		filePath := path.Join(config.DIR_SAVED_FILES, time.Now().Format("02-01-2006-15:04:05.000000")+config.STORE_FILE)

		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Error(err)
		}
		defer f.Close()

		jsonString, err := json.Marshal(memoryStore)
		if err != nil {
			logger.Error(err)
		}

		byteSlice := []byte(jsonString)
		_, err = f.Write(byteSlice)
		if err != nil {
			logger.Error(err)
		} else {
			logger.Info("Memorystore state saved to " + filePath)
		}
	} else {
		logger.Info("Memorystore db is empty")
	}
}
