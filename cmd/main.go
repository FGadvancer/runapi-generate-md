package main

import (
	"flag"
	"fmt"
	"github.com/ruapi-generate-md/internal"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("generate start at ", t)
	projectName := flag.String("p", "ServerAPI", "api project name")
	flag.Parse()
	internal.GeneratePageByItemID("..", *projectName)
	fmt.Println("generate end cost time ", time.Since(t))
}
