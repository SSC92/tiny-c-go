package tinycgo

import (
	simplecgo "github.com/SSC92/simple-go-c"
)

func Tiny(x int) int {
	return simplecgo.MyCFunction(x)
}
