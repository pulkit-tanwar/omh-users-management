package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping - This function will ping the echo server
func (s *Server) Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Ping": "OK"})
}
