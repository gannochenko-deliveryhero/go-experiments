package main

import (
	"fmt"
	"service/internal/util"
)

func main() {
	fmt.Println(util.GetEnv("SERVICE_MESSAGE", "Hello (default)"))
}
