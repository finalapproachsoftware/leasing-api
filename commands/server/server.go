package server

import (
	"net/http"
	"time"

	"github.com/finalapproachsoftware/leasing-api/api"
	"github.com/finalapproachsoftware/leasing-api/logging"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     *logrus.Entry
	addr       string
	api        *api.Api
	httpRouter *mux.Router
}

func newServer() (*Server, error) {
	logger := logging.Log().WithFields(logrus.Fields{
		"source": "server",
	})

	api, err := api.NewApi()
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()

	server := &Server{
		logger:     logger,
		addr:       "localhost:9000",
		api:        api,
		httpRouter: router,
	}

	server.registerUserRoutes()

	return server, nil

}

func (s *Server) start() error {
	srv := &http.Server{
		Handler:      s.httpRouter,
		Addr:         s.addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.logger.Info("Starting HTTP server")
	srv.ListenAndServe()

	return nil
}
