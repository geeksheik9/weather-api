package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ErrorResponse
// struct representation of what gets returned by the RespondWithError function.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

// RespondWithError Utility function to convert an error message into a JSON response.
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	// Clean up all quote marks for readability, marshal adds additional "\" escape char in the response JSON.
	RespondWithJSON(w, code, map[string]string{"error": strings.Replace(msg, `"`, ``, -1)})
}

// RespondNoContent Utility function to send a response without any content.
func RespondNoContent(w http.ResponseWriter, code int) {
	if w != nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(code)
	}
}

// RespondWithJSON Utility function to convert the payload into a JSON response.
// ORIGINAL:
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if w != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Errorf("Error in RespondWithJSON marshal: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(response)
	}
}

// GetJSONRequestBody function to return the request body as string
func GetJSONRequestBody(r *http.Request) (requestBodyJSON string) {
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	requestBodyJSON = string(bodyBytes)

	return requestBodyJSON
}

// CheckError checks the err message and returns a code based on the message.
func CheckError(err error) int {
	var code int
	if err == nil {
		code = http.StatusOK
	} else if strings.Contains(err.Error(), "no documents in result") ||
		strings.Contains(err.Error(), "out of bounds") ||
		strings.Contains(err.Error(), "not found") {
		code = http.StatusNotFound
	} else if strings.Contains(err.Error(), "E11000 duplicate key error") ||
		strings.Contains(err.Error(), "E11001 duplicate key error") {
		code = http.StatusConflict
	} else if strings.Contains(err.Error(), "E10334") ||
		strings.Contains(err.Error(), "Invalid request payload, unable to marshal into json, err: ") {
		code = http.StatusBadRequest
	} else {
		code = http.StatusInternalServerError
	}

	return code
}
