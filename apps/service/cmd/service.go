package main

import (
	"fmt"
)

func main() {
	fmt.Println(getEnv("SERVICE_MESSAGE", "Hello (default)"))
}
