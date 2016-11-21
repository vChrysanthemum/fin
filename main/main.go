package main

import (
	"fmt"
	"in/script"
	"in/ui"
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

func main() {
	GlobalResBaseDir = filepath.Join(os.Getenv("HOME"), ".in")

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	logFile, _ := os.OpenFile(filepath.Join(GlobalResBaseDir, "in.log"),
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	log.SetOutput(logFile)

	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 8192)
		for _ = range sigChan {
			length := runtime.Stack(stacktrace, true)
			log.Println(string(stacktrace[:length]))
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)

	if len(os.Args) < 1 {
		fmt.Println("Project name is needed.")
		return
	}

	projectName := os.Args[1]
	projectMainHtmlFilePath := filepath.Join(GlobalResBaseDir, "project", projectName, "main.html")
	if _, err := os.Stat(projectMainHtmlFilePath); os.IsNotExist(err) {
		fmt.Println("Project is not existed.")
		return
	}

	ui.GlobalOption.ResBaseDir = GlobalResBaseDir
	ui.GlobalOption.ProjectName = projectName

	script.GlobalOption.ResBaseDir = GlobalResBaseDir
	script.GlobalOption.ProjectName = projectName

	ui.PrepareUI()
	content, _ := ioutil.ReadFile(projectMainHtmlFilePath)
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()
}
