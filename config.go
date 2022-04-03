package servicego

import "github.com/kardianos/service"

type DefaultConfig struct{}

func (d *DefaultConfig) Config() *service.Config {

	options := service.KeyValue{
		"Restart":           "on-success",
		"SuccessExitStatus": "1 2 8 SIGKILL",
	}

	dependencies := []string{
		"Requires=network.target",
		"After=network-online.target syslog.target",
	}

	serviceName := ServiceName()

	config := &service.Config{
		Name:             serviceName,
		DisplayName:      serviceName,
		Description:      serviceName,
		WorkingDirectory: ServiceDirectory(),
		Dependencies:     dependencies,
		Option:           options,
	}

	return config
}
