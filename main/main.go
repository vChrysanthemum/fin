package main

import (
	"fin/script"
	"fin/ui"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

var (
	GlobalResBaseDir string
)

func AssertErrIsNil(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Project filepath is needed.")
		return
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	projectPath := os.Args[1]

	GlobalResBaseDir = filepath.Join(os.Getenv("HOME"), ".fin")

	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 8192)
		for _ = range sigChan {
			length := runtime.Stack(stacktrace, true)
			log.Println(string(stacktrace[:length]))
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)

	projectMainHtmlFilePath := filepath.Join(projectPath, "main.html")
	if _, err := os.Stat(projectMainHtmlFilePath); os.IsNotExist(err) {
		fmt.Println("Project is not existed.")
		return
	}

	logFile, err := os.OpenFile(filepath.Join(projectPath, "main.log"),
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	AssertErrIsNil(err)
	log.SetOutput(logFile)

	ui.GlobalOption.ResBaseDir = GlobalResBaseDir
	ui.GlobalOption.ProjectPath = projectPath

	script.GlobalOption.ResBaseDir = GlobalResBaseDir
	script.GlobalOption.ProjectPath = projectPath

	ui.PrepareUI()
	content, _ := ioutil.ReadFile(projectMainHtmlFilePath)
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()
}
