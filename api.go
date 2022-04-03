package servicego

import "github.com/kardianos/service"

// Service extends the service.Interface for setup of service.Config and service.Logger
type Service interface {
	service.Interface
	ConfigProvider
	LoggerContainer
}

// ConfigProvider provides the service.Config for a Service
type ConfigProvider interface {
	// Config provides the configuration for service.New
	Config() *service.Config
}

// LoggerContainer makes the service.Logger available to the Service
type LoggerContainer interface {
	// Logger is called with the output
	Logger(logger service.Logger)
	// Log provides access to the DefaultLogger logger
	//
	// This should not be used until after Run is called
	Log() service.Logger
}
