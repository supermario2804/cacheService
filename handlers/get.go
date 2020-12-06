package handlers

import (
	"cacheDataService/services"
	"cacheDataService/utils"
	"fmt"
	"net/http"
)

func GetTableCache(w http.ResponseWriter, r *http.Request) {
	utils.PrintFatal(fmt.Sprintf("Requested: %s", utils.Info(r.URL)))
	resp, statusCode := services.GetTableCache(w, r)
	resp.Status = statusCode
	w.WriteHeader(statusCode)
	utils.PrintWarn(fmt.Sprintf("Responded: %v", utils.Info(resp.Success, resp.Status, resp.Message)))
	utils.SendHTTPResponse(w, resp)
}

func GetPageCache(w http.ResponseWriter, r *http.Request) {
	utils.PrintFatal(fmt.Sprintf("Requested: %s", utils.Info(r.URL)))
	resp, statusCode := services.GetPageCache(w, r)
	resp.Status = statusCode
	w.WriteHeader(statusCode)
	utils.PrintWarn(fmt.Sprintf("Responded: %v", utils.Info(resp.Success, resp.Status, resp.Message)))
	utils.SendHTTPResponse(w, resp)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.PrintFatal(fmt.Sprintf("Requested: %s", utils.Info(r.URL)))
	resp := new(utils.ApiResponse)
	resp.Status = http.StatusOK
	resp.Message = "Service is healthy"
	resp.Success = true
	w.WriteHeader(http.StatusOK)
	utils.PrintWarn(fmt.Sprintf("Responded: %v", utils.Info(resp.Success, resp.Status, resp.Message)))
	utils.SendHTTPResponse(w, resp)
}
