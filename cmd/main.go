package main

import (
	"fmt"
	"github.com/ruapi-generate-md/internal"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("generate start at ", t)
	internal.GeneratePageByItemID("/data/ruapi-generate-md", "OpenIM服务器API")
	fmt.Println("generate end cost time ", time.Since(t))
}
