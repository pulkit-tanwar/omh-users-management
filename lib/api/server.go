package api

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pulkit-tanwar/omh-users-management/lib/config"
	log "github.com/sirupsen/logrus"
)

// Server - Structure for server
type Server struct {
	*config.Config
	router *echo.Echo
}

// NewServer - Constructor function for server
func NewServer(cfg *config.Config) *Server {
	server := &Server{
		Config: cfg,
		router: echo.New(),
	}

	server.router.GET(path.Join(cfg.APIPath, "/api/v1/ping"), server.Ping)
	server.router.POST(path.Join(cfg.APIPath, "/api/v1/users"), server.CreateUser)
	server.router.GET(path.Join(cfg.APIPath, "/api/v1/users/:userName"), server.FetchUserByUserName)
	server.router.Use(middleware.Logger())

	return server
}

// Start - This function will start the echo server
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	log.Infof("Listening on %s", address)
	return s.router.Start(address)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { // For unit testing only
	s.router.ServeHTTP(w, r)
}

// Stop - This function will stop the echo server
func (s *Server) Stop(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
