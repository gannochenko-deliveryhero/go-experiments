package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
