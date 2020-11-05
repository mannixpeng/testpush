package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var appName string
	var goos string
	var cgoEnabled string
	var goArch string
	//var path string

	//flag.StringVar(&path, "path", "", "文件路径")
	flag.StringVar(&appName, "appName", "huasheng", "app名称")
	flag.StringVar(&goos, "os", "linux", "运行平台")
	flag.StringVar(&cgoEnabled, "cgoEnabled", "0", "是否使能cgo")
	flag.StringVar(&goArch, "goArch", "amd64", "架构")
	flag.Parse()

	os.Setenv("CGO_ENABLED", cgoEnabled)
	os.Setenv("GOOS", goos)
	os.Setenv("GOARCH", goArch)

	name := appName+"_"+goos+"_"+goArch
	if strings.Contains(goos, "window") {
		name += ".exe"
	}
	cmd := exec.Command("cmd", "/C", fmt.Sprintf("go build -o %s", name))

	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}

	defer stdout.Close() // 保证关闭输出流

	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}
