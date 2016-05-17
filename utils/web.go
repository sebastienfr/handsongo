package utils

import (
	"encoding/json"
	logger "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	// ResponseHeaderContentTypeKey is the key used for response content type
	ResponseHeaderContentTypeKey = "Content-Type"
	// ResponseHeaderContentTypeJSONUTF8 is the key used for UTF8 JSON response
	ResponseHeaderContentTypeJSONUTF8 = "application/json; charset=UTF-8"

	resourceNotFound = "Resource not found"
	errorMsg         = "Error"
)

// SendJSONWithHTTPCode outputs JSON with an HTTP code
func SendJSONWithHTTPCode(w http.ResponseWriter, d interface{}, code int) {
	w.Header().Set(ResponseHeaderContentTypeKey, ResponseHeaderContentTypeJSONUTF8)
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		logger.WithField("body", d).WithField("code", code).Error("error while encoding JSON body of response")
		// panic will cause the http.StatusInternalServerError to be send to users
		panic(err)
	}
}

// SendJSON outputs a JSON with http.StatusOK code
func SendJSONOk(w http.ResponseWriter, d interface{}) {
	SendJSONWithHTTPCode(w, d, http.StatusOK)
}

// SendError sends error with a custom message and error code
func SendJSONError(w http.ResponseWriter, error string, code int) {
	SendJSONWithHTTPCode(w, map[string]string{errorMsg: error}, code)
}

// SendNotFound produces a http.StatusNotFound response with the following JSON, '{"Error":"Resource not found"}'
func SendJSONNotFound(w http.ResponseWriter) {
	SendJSONError(w, resourceNotFound, http.StatusNotFound)
}

// NotFoundHandler return a JSON implementation of the not found handler
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendJSONNotFound(w)
	}
}

// ParamAsString returns an URL parameter /{name} as a string
func ParamAsString(name string, r *http.Request) string {
	vars := mux.Vars(r)
	value := vars[name]
	return value
}
