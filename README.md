# tiny-c-go
Example project for using packages that uses C code in Go.

## How to use

If you are willing to use your own C code, go to [my simple example](https://github.com/SSC92/simple-go-c) and follow the instructions for creating a package that uses C code.


Import your package in your Go code(using an allias could help you to avoid name conflicts):

```go
import alias "github.com/yourname/yourpackage"
```

Then you can use the package as you would normally do.

```go
return alias.YourFunction(args)
```


