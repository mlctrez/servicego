package servicego

import (
	"fmt"
	"github.com/kardianos/service"
	"strings"
)

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

// RequiresService allows adding additional items to the requires section of the service config.
//
//  type svc struct {
//	  servicego.Defaults
//  }
//
//	func (s *svc) Config() *service.Config {
//	  return servicego.RequiresService(s, "other")
//	}
func RequiresService(provider ConfigProvider, name string) *service.Config {
	config := provider.Config()
	deps := config.Dependencies
	for i, dependency := range deps {
		if strings.HasPrefix(dependency, "Requires") {
			dependency += fmt.Sprintf(" %s.service", name)
			deps[i] = dependency
			break
		}
	}
	return config
}
