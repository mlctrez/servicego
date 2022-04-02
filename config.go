package servicego

import "github.com/kardianos/service"

type DefaultConfig struct{}

func (d *DefaultConfig) Config() *service.Config {
	return defaultConfig()
}

func defaultConfig() *service.Config {

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	dependencies := []string{
		"Requires=network.target",
		"After=network-online.target syslog.target",
	}

	config := &service.Config{
		Name:         ServiceName(),
		DisplayName:  ServiceName(),
		Description:  ServiceName(),
		Dependencies: dependencies,
		Option:       options,
	}

	return config
}
