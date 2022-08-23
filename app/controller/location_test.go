package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"scan/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocations(t *testing.T) {
	c, _ := GetTestController()
	req, _ := http.NewRequest("GET", "/api/v1/locations", nil)
	w := httptest.NewRecorder()
	c.Server.ServeHTTP(w, req)

	var locations []model.Location
	json.Unmarshal(w.Body.Bytes(), &locations)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, locations)
}
