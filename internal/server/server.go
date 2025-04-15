package server

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *config.ServerConfig
	logger *logrus.Logger
	router *mux.Router
}

func New(config *config.ServerConfig) *Server {
	return &Server {
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	err := s.configureLogger()
	if err != nil {
		return err
	}
	
	s.configureRouter()
	
	s.logger.Info("Starting server on address: http://localhost:8000")
	
	return http.ListenAndServe(s.config.Addr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	
	s.logger.SetLevel(level)
	
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/hello", handlers.HandleHello())
}