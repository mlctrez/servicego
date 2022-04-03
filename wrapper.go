package servicego

import (
	"flag"
	"os"

	"github.com/kardianos/service"
)

type Defaults struct {
	DefaultConfig
	DefaultLogger
}

// Run is intended to be only called from a main method
func Run(s Service) {
	s.Logger(service.ConsoleLogger)

	svc, err := service.New(s, s.Config())
	fatal(s, err)

	// todo: implement error channel?
	logger, err := svc.Logger(nil)
	fatal(s, err)

	s.Logger(logger)

	action := parseAction()

	switch action {
	case "deploy":
		fatal(s, deploy(s))
	case "run":
		fatal(s, svc.Run())
	default:
		fatal(s, service.Control(svc, action))
	}

}

func fatal(impl Service, err error) {
	if err == nil {
		return
	}
	impl.Log().Error(err)
	os.Exit(1)
}

func parseAction() string {
	var action string
	flag.StringVar(&action, "action", "run", actionUsage())
	flag.Parse()
	if !isValidAction(action) {
		flag.Usage()
		os.Exit(-1)
	}
	return action
}

func isValidAction(action string) bool {
	for _, s := range service.ControlAction {
		if action == s {
			return true
		}
	}
	return action == "deploy" || action == "run"
}

func actionUsage() string {
	result := "service and deployment control actions : "
	for _, s := range service.ControlAction {
		result += s + ", "
	}
	result += "deploy, run"
	return result
}
