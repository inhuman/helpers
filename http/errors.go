package http

import (
	"net/http"
	"encoding/json"
	"github.com/inhuman/helpers/common"
)

type httpError struct {
	Error string `json:"error"`
}


// CheckErrorHTTP is used for check error and write it to http.ResponseWriter, if it not nil.
// Return true if error not nil
func CheckErrorHTTP(err error, w http.ResponseWriter, code int) bool {

	if err != nil {
		var er httpError
		er.Error = err.Error()
		jsn, jsnErr := json.Marshal(er)
		common.CheckError(jsnErr)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(code)
		w.Write([]byte(jsn))
		return true
	}
	return false
}