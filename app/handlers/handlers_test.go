package handlers

import (
	"MemoryStore/app/models"
	"MemoryStore/config"
	"MemoryStore/constants"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func TestGetValKeyNeccessary(t *testing.T) {

	var response models.ApiResponse

	req, err := http.NewRequest("GET", "/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVal)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusBadRequest {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusBadRequest)
	}

	if response.Error != constants.ErrorKeyNecessary {
		t.Errorf("Erroe message wrong: got %v want %v",
			response.Error, constants.ErrorKeyNecessary)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetValNotFound(t *testing.T) {
	var response models.ApiResponse

	req, err := http.NewRequest("GET", "/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("key", "notexist")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVal)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusNotFound {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusNotFound)
	}

	if response.Error != constants.ErrorNoDataFound {
		t.Errorf("Erroe message wrong: got %v want %v",
			response.Error, constants.ErrorNoDataFound)
	}

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestGetValKeyOnlyGetMethodAllowed(t *testing.T) {

	var response models.ApiResponse

	req, err := http.NewRequest("POST", "/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVal)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusBadRequest {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusBadRequest)
	}

	if response.Error != constants.ErrorOnlyGet {
		t.Errorf("Erroe message wrong: got %v want %v",
			response.Error, constants.ErrorOnlyGet)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetVal(t *testing.T) {
	InitMemoryStore()
	SetStoreData(models.SetReturnData("age", "27"))

	var response models.ApiResponse

	req, err := http.NewRequest("GET", "/api/get", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("key", "age")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVal)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusOK {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusOK)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedKey := "age"
	if response.Data.Key != expectedKey {
		t.Errorf("handler returned unexpected body data: got %v want %v",
			response.Data.Key, expectedKey)
	}

	expectedValue := "27"
	if response.Data.Value != expectedValue {
		t.Errorf("handler returned unexpected body data: got %v want %v",
			response.Data.Key, expectedValue)
	}
}

func TestSetBodyCantEmpty(t *testing.T) {

	InitMemoryStore()

	var response models.ApiResponse

	req, err := http.NewRequest("POST", "/api/set", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SetKey)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusBadRequest {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusBadRequest)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestSet(t *testing.T) {

	InitMemoryStore()

	var jsonStr = []byte(`{"key":"firstname", "value":"mehmet"}`)

	var response models.ApiResponse

	req, err := http.NewRequest("POST", "/api/set", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SetKey)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusOK {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusOK)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedKey := "firstname"
	if response.Data.Key != expectedKey {
		t.Errorf("handler returned unexpected body key: got %v want %v",
			rr.Body.String(), expectedKey)
	}

	expectedValue := "mehmet"
	if response.Data.Value != expectedValue {
		t.Errorf("handler returned unexpected body key: got %v want %v",
			rr.Body.String(), expectedValue)
	}

	if val, _ := GetStoreData(expectedKey); val != expectedValue {
		t.Errorf("Memory db has invalid value : got %v want %v",
			val, expectedValue)
	}
}

func TestSetOnlyPostMethodAllowed(t *testing.T) {

	var response models.ApiResponse

	req, err := http.NewRequest("HEAD", "/api/set", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SetKey)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusBadRequest {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusBadRequest)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	if response.Error != constants.ErrorOnlyPost {
		t.Errorf("Erroe message wrong: got %v want %v",
			response.Error, constants.ErrorOnlyPost)
	}

}

func TestFlush(t *testing.T) {
	var response models.ApiResponse

	config.DIR_SAVED_FILES = t.TempDir()
	config.DIR_ARCHIVE_FILES = path.Join(config.DIR_SAVED_FILES, "archive")

	req, err := http.NewRequest("DELETE", "/api/flush", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Flush)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusOK {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusOK)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if response.Data.Key != "" && response.Data.Value != "" {
		t.Errorf("handler returned unexpected body: got %v want empty data",
			rr.Body.String())
	}
}

func TestFlushOnlyDeleteMethodAllowed(t *testing.T) {
	var response models.ApiResponse

	req, err := http.NewRequest("PUT", "/api/flush", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Flush)
	handler.ServeHTTP(rr, req)

	err = json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Error("Response body not valid")
	}

	if response.Status != http.StatusBadRequest {
		t.Errorf("Body has wrong status code: got %v want %v",
			response.Status, http.StatusBadRequest)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	if response.Error != constants.ErrorOnlyDelete {
		t.Errorf("Erroe message wrong: got %v want %v",
			response.Error, constants.ErrorOnlyDelete)
	}

}
