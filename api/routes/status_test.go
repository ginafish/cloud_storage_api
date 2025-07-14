package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type header struct {
	Key   string
	Value string
}

// Test that the status endpoint exists and returns a status json object
func TestGetStatus(t *testing.T) {
	r := getRouter(true)
	r.GET("/api/v1/status", getAppStatus)

	w := getHTTPResponse(r, "GET", "/api/v1/status")

	statusOK := w.Code == http.StatusOK
	if !statusOK {
		t.Errorf("Status code was %d", w.Code)
	}
	resp, err := ioutil.ReadAll(w.Body)
	var appStatuses []status
	json.Unmarshal(resp, &appStatuses)
	reqSuccess := err == nil
	if !reqSuccess {
		t.Errorf("An error occurred: %d", err)
	}
	gotExpectedData := len(appStatuses) == 1 && appStatuses[0].app == "api" && appStatuses[0].status == "UP"
	if !gotExpectedData {
		t.Log(resp)
		t.Log(appStatuses)
		t.Error("Status data not formatted as expected")
	}
}
