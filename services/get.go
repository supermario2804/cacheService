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

func GetTableCache(w http.ResponseWriter, r *http.Request) (utils.ApiResponse, int) {
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
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

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
