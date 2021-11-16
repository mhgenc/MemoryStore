package models

type ApiResponse struct {
	Status int      `json:"status,omitempty"`
	Error  string   `json:"error,omitempty"`
	Data   *ApiData `json:"data,omitempty"`
}

type ApiData struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

//SetApiResponse Set ApiResponse type
func SetApiResponse(status int, error string, data *ApiData) *ApiResponse {
	apiResponse := new(ApiResponse)

	apiResponse.Status = status
	apiResponse.Error = error
	apiResponse.Data = data

	return apiResponse
}

//GetStatus return api response status
func (a *ApiResponse) GetStatus() int {
	return a.Status
}

//SetReturnData set return data struct
func SetReturnData(key, value string) *ApiData {
	returnData := new(ApiData)
	returnData.Key = key
	returnData.Value = value

	return returnData
}
