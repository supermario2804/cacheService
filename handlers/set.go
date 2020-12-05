package handlers

import (
	"cacheDataService/services"
	"cacheDataService/utils"
	"fmt"
	"net/http"
)

func SetTableCache(w http.ResponseWriter, r *http.Request) {
	utils.PrintFatal(fmt.Sprintf("Requested: %s", utils.Info(r.URL)))
	resp, statusCode := services.SetTableCache(w, r)
	resp.Status = statusCode
	w.WriteHeader(statusCode)
	utils.PrintWarn(fmt.Sprintf("Responded: %v", utils.Info(resp.Success, resp.Status, resp.Message)))
	utils.SendHTTPResponse(w, resp)
}

func SetPageCache(w http.ResponseWriter, r *http.Request) {
	utils.PrintFatal(fmt.Sprintf("Requested: %s", utils.Info(r.URL)))
	resp, statusCode := services.SetPageCache(w, r)
	resp.Status = statusCode
	w.WriteHeader(statusCode)
	utils.PrintWarn(fmt.Sprintf("Responded: %v", utils.Info(resp.Success, resp.Status, resp.Message)))
	utils.SendHTTPResponse(w, resp)
}
