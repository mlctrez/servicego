# servicego

### Background

The servicego module provides a wrapper for the functionality
in [github.com/kardianos/service](https://github.com/kardianos/service)

This was written to simplify my usage of [github.com/kardianos/service](https://github.com/kardianos/service) and
therefore it is
very [itch-to-scratch](https://opensource.com/article/17/4/itch-to-scratch-model-user-problems#:~:text=Basically%2C%20the%20idea%20is%20that,to%20describe%20what%20actually%20happens.)

If this is not the same itch that you have then this module is not for you.

### Your still here

Yes this is a litmus test. If you are still here after we mangled the english language then continue. No funny business
after this I promise. [don't read into it](https://www.google.com/search?q=your+vs+you%27re)

### O hey where were we

Build your service and embed *the defaults*. Find those
elusive [Defaults](https://github.com/mlctrez/servicego/search?q=Defaults) via github.

```go
type exampleService struct {
servicego.Defaults
}
```

If you like the defaults then nothing needs to be done. Go away now or keep reading.

### If you dare

So the defaults are basic but you can implement your own logging and configuration implementations. Use the soure luke. 
