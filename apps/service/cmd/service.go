package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(envString("SERVICE_MESSAGE", "Hello (default)"))
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
