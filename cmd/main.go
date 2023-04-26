package main

import (
	"flag"
	"fmt"
	"github.com/ruapi-generate-md/internal"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("generate start at ", t)
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(path)
	_ = filepath.Dir(dir)

	//fmt.Printf("当前代码执行目录的上一级目录为：%s\n", parentDir)
	projectName := flag.String("p", "ServerAPI", "api project name")
	flag.Parse()
	internal.GeneratePageByItemID("..", *projectName)
	fmt.Println("generate end cost time ", time.Since(t))
}
