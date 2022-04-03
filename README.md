# servicego

### Background

The servicego package provides a wrapper for the functionality
in [github.com/kardianos/service](https://github.com/kardianos/service)

This was written to simplify usage when writing lots of services, eliminating 
much of the boiler place code to set up a new service.

### Defaults

Service name, display name, and description all default to the executable name.

If this is not the same itch that you have then this module is not for you.

### A Brief How-To

Embed [Defaults](wrapper.go) in a struct for your [Service](api.go).

```go
type exampleService struct {
    servicego.Defaults
}
```
Implement Start and Stop as you would in 
[github.com/kardianos/service](https://github.com/kardianos/service)

```go
func (e *exampleService) Start(s service.Service) error {
    // Start should not block so do non-setup work in goroutines.
    // Return non nil error if the service cannot be started.
}

func (e *exampleService) Stop(s service.Service) error {
    // Cancel goroutines if you require a graceful shutdown.
    // Stop should not block and should return with a few seconds.
}
```
That's it.
