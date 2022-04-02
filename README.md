# servicego

The servicego package wraps [github.com/kardianos/service](https://github.com/kardianos/service)
to remove boilerplate and also to automate the installation of the service on a host.

The installation is very opinionated with respect to where service binaries live, the default user the service runs as,
and naming of the service based on the go executable name.

### Summary

* Build an application that implements [Service](api.go)
* Embed the [Defaults](defaults.go) struct to handle logging and reasonable defaults
    * Or choose [DefaultLogger](logging.go) or [DefaultConfig](config.go) individually based on your needs
* Your main method should call goservice.Run(yourimpl Service)
* Build your application binaries for your target platform(s)
* Deploy built binaries on these platforms with `yourbinary -action deploy`
* Stop, start, uninstall services using typical [github.com/kardianos/service](https://github.com/kardianos/service)
  actions

### Example

See [example.go](example/example.go) for an example service.

