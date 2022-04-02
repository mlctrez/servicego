package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kardianos/service"
	sg "github.com/mlctrez/servicego"
)

func main() {
	sg.Run(&exampleService{})
}

var _ sg.Service = (*exampleService)(nil)

type exampleService struct {
	sg.Defaults
	server *http.Server
}

func (e *exampleService) Config() *service.Config {
	// example of how to provide a custom config, this just delegates to the default config
	config := e.DefaultConfig.Config()
	config.Description = "override the default description"
	return config
}

func (e *exampleService) Start(s service.Service) error {
	e.Log().Info("starting")
	return e.startHttp()
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if strings.TrimPrefix(request.URL.Path, "/") == "" {
		writer.Write([]byte("OK"))
		return
	}
	writer.WriteHeader(http.StatusNotFound)
}

func (e *exampleService) startHttp() error {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	e.server = &http.Server{Handler: mux}

	go func() {
		serverErr := e.server.Serve(listen)
		if serverErr != nil && serverErr != http.ErrServerClosed {
			e.Log().Error(serverErr)
		}
	}()

	return nil
}

func (e *exampleService) Stop(s service.Service) error {
	e.Log().Info("stopping")
	if e.server == nil {
		return nil
	}

	stopContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := e.server.Shutdown(stopContext); err != nil {
		e.Log().Error("server shutdown failed :%e", err)
		os.Exit(-1)
	}
	e.Log().Info("normal exit")

	return nil
}
