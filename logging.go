package servicego

import "github.com/kardianos/service"

type DefaultLogger struct {
	logger service.Logger
}

func (d *DefaultLogger) Logger(logger service.Logger) {
	d.logger = logger
}

func (d *DefaultLogger) Log() service.Logger {
	return d.logger
}
