package services

import (
	"cacheDataService/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
)

type setRequest struct {
	Table      string      `json:"table"`
	PrimaryKey string      `json:"pk"`
	Data       interface{} `json:"data"`
}

type setPageRequest struct {
	Title      string        `json:"title"`
	PageNumber string        `json:"pageNumber"`
	SortBy     string        `json:"sortby"`
	Data       []interface{} `json:"data"`
}

func SetTableCache(w http.ResponseWriter, r *http.Request) (utils.ApiResponse, int) {
	var reqdata setRequest

	apiResp := utils.ApiResponse{
		Success: false,
		Status:  http.StatusInternalServerError,
		Message: "Something went wrong",
		Data:    make(map[string]interface{}),
	}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		utils.HandleError(readErr)
		return apiResp, http.StatusBadRequest
	}
	jsonErr := json.Unmarshal(body, &reqdata)
	if jsonErr != nil {
		utils.HandleError(jsonErr, "Error while unmarshaling body")
		return apiResp, http.StatusBadRequest
	}
	defer r.Body.Close()

	if reqdata.PrimaryKey == "" || reqdata.Table == "" {
		utils.HandleError(fmt.Errorf("PrimaryKey or Table field is missing"), reqdata.PrimaryKey, reqdata.Table)
		return apiResp, http.StatusBadRequest
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()
	pong, connErr := client.Ping().Result()
	if connErr != nil {
		utils.HandleError(connErr)
		return apiResp, http.StatusInternalServerError
	}
	utils.PrintSuccess(pong)

	jsonData, jsonErr := json.Marshal(reqdata.Data)
	if jsonErr != nil {
		utils.HandleError(jsonErr)
		return apiResp, http.StatusInternalServerError
	}

	setErr := client.Set(reqdata.Table+"_"+reqdata.PrimaryKey, jsonData, 0).Err()
	if setErr != nil {
		utils.HandleError(setErr)
		return apiResp, http.StatusInternalServerError
	}
	apiResp.Success = true
	apiResp.Status = http.StatusOK
	apiResp.Message = "Cache set successfully"
	return apiResp, http.StatusOK
}

func SetPageCache(w http.ResponseWriter, r *http.Request) (utils.ApiResponse, int) {
	var reqdata setPageRequest

	apiResp := utils.ApiResponse{
		Success: false,
		Status:  http.StatusInternalServerError,
		Message: "Something went wrong",
		Data:    make(map[string]interface{}),
	}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		utils.HandleError(readErr)
		return apiResp, http.StatusBadRequest
	}
	jsonErr := json.Unmarshal(body, &reqdata)
	if jsonErr != nil {
		utils.HandleError(jsonErr, "Error while unmarshaling body")
		return apiResp, http.StatusBadRequest
	}
	defer r.Body.Close()

	if reqdata.Title == "" || reqdata.PageNumber == "" || reqdata.SortBy == "" {
		utils.HandleError(fmt.Errorf("Title or Pagenumber or Sortby field is missing"), reqdata.Title, reqdata.PageNumber, reqdata.SortBy)

		return apiResp, http.StatusBadRequest
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()
	pong, connErr := client.Ping().Result()
	if connErr != nil {
		utils.HandleError(connErr)
		return apiResp, http.StatusInternalServerError
	}
	utils.PrintSuccess(pong)

	jsonData, jsonErr := json.Marshal(reqdata.Data)
	if jsonErr != nil {
		utils.HandleError(jsonErr)
		return apiResp, http.StatusInternalServerError
	}

	setErr := client.Set(reqdata.Title+"_"+reqdata.PageNumber+"_"+reqdata.SortBy, jsonData, 0).Err()
	if setErr != nil {
		utils.HandleError(setErr)
		return apiResp, http.StatusInternalServerError
	}
	apiResp.Success = true
	apiResp.Status = http.StatusOK
	apiResp.Message = "Page Cache set successfully"
	return apiResp, http.StatusOK
}
