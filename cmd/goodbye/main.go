package main

import (
	"fmt"

	"github.com/cttttt/multi-go-go/cmd/goodbye/pkg/goodbye"
	"github.com/cttttt/multi-go-go/pkg/hello"
)

func main() {
	fmt.Printf("%s and %s, World\n", hello.Hello(), goodbye.Goodbye())
}
