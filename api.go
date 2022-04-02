package servicego

import "github.com/kardianos/service"

type Service interface {
	service.Interface
	ConfigProvider
	LoggerContainer
}

type ConfigProvider interface {
	Config() *service.Config
}

type LoggerContainer interface {
	Logger(logger service.Logger)
	Log() service.Logger
}
