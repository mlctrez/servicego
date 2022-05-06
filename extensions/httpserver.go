package extensions

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/kardianos/service"
	"github.com/mlctrez/servicego"
)

type HttpServer struct {
	server *http.Server
	logger service.Logger
}

func (hs *HttpServer) Start(address string, handler http.Handler, lc servicego.LoggerContainer) error {

	hs.logger = lc.Log()

	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	hs.server = &http.Server{Handler: handler}

	go func() {
		serverErr := hs.server.Serve(listen)

		if serverErr != nil && serverErr != http.ErrServerClosed {
			_ = hs.logger.Error(serverErr)
		}
	}()

	hs.logger("listening on %q", listen.Addr().String())

	return nil
}

func (hs *HttpServer) Stop(timeout time.Duration) error {

	if timeout == 0 {
		timeout = 5 * time.Second
	}

	if hs.server != nil {
		if err := runWithTimeout(hs.server.Shutdown, timeout); err != nil {
			return err
		}
	}
	_ = hs.logger.Info("normal exit")

	return nil
}

func runWithTimeout(task func(ctx context.Context) error, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return task(ctx)
}
