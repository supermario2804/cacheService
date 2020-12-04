package handlers
import (
	"net/http"
	""
)




/*DeleteBill Receives a bill id and soft deletes it from database*/
func SetTableCache(w http.ResponseWriter, r *http.Request) {
	utils.PrintFatal(fmt.Sprintf("Requested: %s", utils.Info(r.URL)))
	resp, statusCode := services.SetTableCache(w, r)
	resp.Status = statusCode
	w.WriteHeader(statusCode)
	utils.PrintWarn(fmt.Sprintf("Responded: %v", utils.Info(resp.Success, resp.Status, resp.Message)))
	utils.SendHTTPResponse(w, resp)
}