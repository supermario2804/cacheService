package services

import (
	"cacheDataService/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
	"_ github.com/gomodule/redigo/redis"
)

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

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, connErr := client.Ping().Result()
	if connErr != nil {
		utils.HandleError(connErr)
	}
return apiResp,http.StatusInternalServerError
}
