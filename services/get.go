package services

import (
	"cacheDataService/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
)

type getRequest struct {
	Table      string `json:"table"`
	PrimaryKey string `json:"pk"`
}

type getPageRequest struct {
	Title      string `json:"title"`
	PageNumber string `json:"pageNumber"`
	SortBy     string `json:"sortby"`
}

func GetTableCache(w http.ResponseWriter, r *http.Request) (utils.ApiResponse, int) {
	var reqdata getRequest

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
	utils.PrintInfo(pong)

	val, getErr := client.Get(reqdata.Table + "_" + reqdata.PrimaryKey).Result()
	if getErr != nil {
		utils.HandleError(getErr)
		return apiResp, http.StatusInternalServerError
	}

	apiResp.Success = true
	apiResp.Status = http.StatusOK
	apiResp.Message = "Cache fetched successfully"
	apiResp.Data["tableData"] = val
	return apiResp, http.StatusOK
}

func GetPageCache(w http.ResponseWriter, r *http.Request) (utils.ApiResponse, int) {
	var reqdata getPageRequest

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

	if reqdata.Title == "" || reqdata.SortBy == "" || reqdata.PageNumber == "" {
		utils.HandleError(fmt.Errorf("Title or Sortby or Pagenumber field is missing"), reqdata.Title, reqdata.PageNumber, reqdata.SortBy)
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
	utils.PrintInfo(pong)

	val, getErr := client.Get(reqdata.Title + "_" + reqdata.PageNumber + "_" + reqdata.SortBy).Result()
	if getErr != nil {
		utils.HandleError(getErr)
		return apiResp, http.StatusInternalServerError
	}

	apiResp.Success = true
	apiResp.Status = http.StatusOK
	apiResp.Message = "Cache fetched successfully"
	apiResp.Data["pageData"] = val
	return apiResp, http.StatusOK
}
