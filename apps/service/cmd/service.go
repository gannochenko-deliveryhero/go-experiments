package service

import (
	"fmt"
)

func main() {
	fmt.Println(envString("SERVICE_MESSAGE", "Hello (default)"))
}
