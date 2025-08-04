package main

import (
	"github.com/katatrina/go12-service/cmd"
)

func main() {
	cmd.Execute() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
