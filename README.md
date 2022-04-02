# servicego

The servicego package wraps [github.com/kardianos/service](https://github.com/kardianos/service) to remove a 
lot of the boilerplate setup and also to automate the installation of the service on a host.

### Summary

* Build an application that implements 
[service.Interface](https://github.com/kardianos/service/blob/5c08916379a92cb1806764e911af33c55762a753/service.go#L331)
* call goservice.Run(yourimpl service.Interface) in a main method
* build your application binaries for your target platform(s)
* deploy built binaries with `yourbinary -action deploy`


