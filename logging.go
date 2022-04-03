package servicego

import "github.com/kardianos/service"

// DefaultLogger is provides a default implementation of LoggerContainer for a Service
type DefaultLogger struct {
	logger service.Logger
}

func (d *DefaultLogger) Logger(logger service.Logger) {
	d.logger = logger
}

func (d *DefaultLogger) Log() service.Logger {
	return d.logger
}
