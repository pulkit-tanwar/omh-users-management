package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/pulkit-tanwar/omh-users-management/lib/api"
	"github.com/pulkit-tanwar/omh-users-management/lib/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func serve(t *testing.T, request *http.Request, cfg *config.Config) *httptest.ResponseRecorder {
	server, rr := makeServer(t, cfg)
	server.ServeHTTP(rr, request)
	return rr
}

func makeServer(t *testing.T, cfg *config.Config) (*api.Server, *httptest.ResponseRecorder) {
	rr := httptest.NewRecorder()
	server := api.NewServer(cfg)
	return server, rr
}

func TestStartAndStop(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.Host = "localhost"
	cfg.Port = 9999

	server, _ := makeServer(t, cfg)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Start()
		require.NotNil(t, err)
		assert.Equal(t, "http: Server closed", err.Error())
	}()

	time.Sleep(50 * time.Millisecond) // Giving some time, so that server can start up

	err := server.Stop(context.Background())
	require.Nil(t, err)
	wg.Wait()
}

func get(path string) *http.Request {
	return httptest.NewRequest(echo.GET, path, nil)
}
