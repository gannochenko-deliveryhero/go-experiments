package main

import (
	"fmt"
	util "github.com/gannochenko-deliveryhero/go-experiments/apps/service/internal/util"
)

func main() {
	fmt.Println(util.getEnv("SERVICE_MESSAGE", "Hello (default)"))
}
