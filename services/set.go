package services

import (
	"cacheDataService/utils"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
	//"_ github.com/gomodule/redigo/redis"
)

var ctx = context.Background()

type setRequest struct {
	Table      string      `json:"table"`
	PrimaryKey string      `json:"pk"`
	Data       interface{} `json:"data"`
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
		utils.HandleError(fmt.Errorf("PrimaryKey or Table field is missing"),reqdata.PrimaryKey,reqdata.Table)
		return apiResp, http.StatusBadRequest
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, connErr := client.Ping(ctx).Result()
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

	setErr := client.Set(ctx, reqdata.Table+"_"+reqdata.PrimaryKey, jsonData, 0).Err()
	if setErr != nil {
		utils.HandleError(setErr)
		return apiResp, http.StatusInternalServerError
	}
	fmt.Println(string(jsonData))
	apiResp.Success = true
	apiResp.Status = http.StatusOK
	apiResp.Message = "Cache set successfully"
	return apiResp, http.StatusOK
}
