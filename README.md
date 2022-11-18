# tiny-c-go
Example project for using packages that uses C code in Go.

## How to use

Import your package in your Go code(using an allias could help you to avoid name conflicts):

```go
import alias "github.com/yourname/yourpackage"
```

Then you can use the package as you would normally do.

```go
return alias.YourFunction(args)
```


