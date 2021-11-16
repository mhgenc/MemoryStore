package handlers

import (
	"MemoryStore/app/models"
	"MemoryStore/constants"
	"encoding/json"
	"net/http"
)

//GetVal Returns the value of the key given to the URL as a parameter
func GetVal(w http.ResponseWriter, r *http.Request) {
	var returnData *models.ApiData

	if (r.Method) == http.MethodGet {
		if key, ok := r.URL.Query()["key"]; ok {
			value, err := GetStoreData(key[0])
			if err != nil {
				PrepareResponse(w, http.StatusNotFound, constants.ErrorNoDataFound, returnData)
				return
			} else {
				returnData = models.SetReturnData(key[0], value)
				PrepareResponse(w, http.StatusOK, "", returnData)
				return
			}
		} else {
			PrepareResponse(w, http.StatusBadRequest, constants.ErrorKeyNecessary, returnData)
			return
		}
	} else {
		PrepareResponse(w, http.StatusBadRequest, constants.ErrorOnlyGet, returnData)
		return
	}
}

//SetKey Fills the sent key value in the body to the memorystore
func SetKey(w http.ResponseWriter, r *http.Request) {
	var reqData models.ApiData
	if r.Method == http.MethodPost {
		if r.Body == nil {
			PrepareResponse(w, http.StatusBadRequest, constants.ErrorBodyCantEmpty, &reqData)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			PrepareResponse(w, http.StatusBadRequest, err.Error(), &reqData)
			return
		}

		if len(reqData.Key) > 0 && len(reqData.Value) > 0 {
			SetStoreData(&reqData)
			PrepareResponse(w, http.StatusOK, "", &reqData) //created can also be returned
			return
		} else {
			PrepareResponse(w, http.StatusBadRequest, constants.ErrorKeyValueCantEmpty, &reqData)
			return
		}
	} else {
		PrepareResponse(w, http.StatusBadRequest, constants.ErrorOnlyPost, &reqData)
		return
	}
}

//Flush Cleans memorystore and archives saved backup files (./savedfiles/archive/)
func Flush(w http.ResponseWriter, r *http.Request) {
	var data models.ApiData
	if r.Method == http.MethodDelete {
		FlushMemoryStore()
		//archive saved data
		err := ArchiveBackups()

		if err != nil {
			PrepareResponse(w, http.StatusInternalServerError, err.Error(), &data)
		} else {
			PrepareResponse(w, http.StatusOK, "", &data) // nocontent can also be returned
		}
		return
	} else {
		PrepareResponse(w, http.StatusBadRequest, constants.ErrorOnlyDelete, &data)
		return
	}
}

func InvalidPath(w http.ResponseWriter, r *http.Request) {
	PrepareResponse(w, http.StatusBadRequest, constants.ErrorInvalidPath, nil)
	return
}

func InvalidMethod(w http.ResponseWriter, r *http.Request) {
	PrepareResponse(w, http.StatusBadRequest, constants.ErrorInvalidMethod, nil)
	return
}
