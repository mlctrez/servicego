# servicego

### Background

The servicego package provides a wrapper for the functionality
in [github.com/kardianos/service](https://github.com/kardianos/service)

This was written to simplify usage when writing lots of services, eliminating much of the boiler place code to set up a new service.

### Defaults

Service name, display name, and description all default to the executable name.

If this is not the same itch that you have then this module is not for you.

### A Brief How-To

Build your service and embed *the defaults*. Find those
elusive [Defaults](https://github.com/mlctrez/servicego/search?q=Defaults) via github.

```go
type exampleService struct {
servicego.Defaults
}
```

If you like the defaults then nothing needs to be done. Go away now or keep reading.

### If you dare

So the defaults are basic but you can implement your own 
logging and configuration implementations. Use the source luke. 
