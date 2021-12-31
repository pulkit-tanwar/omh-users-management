package api_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/pulkit-tanwar/omh-users-management/lib/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	rr := serve(t, get("/api/v1/ping"), config.DefaultConfig())
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", strings.ToLower(rr.HeaderMap["Content-Type"][0]))

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))

	assert.Equal(t, map[string]interface{}{"Ping": "OK"}, response)
}
