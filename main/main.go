package main

import (
	"fmt"
	"in/ui"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	if len(os.Args) < 1 {
		fmt.Println("Project name is needed.")
		return
	}

	projectName := os.Args[1]
	projectMainHtmlFilePath := filepath.Join(GlobalResBaseDir, "project/"+projectName+"/main.html")
	if _, err := os.Stat(projectMainHtmlFilePath); os.IsNotExist(err) {
		fmt.Println("Project is not existed.")
		return
	}

	ui.PrepareUI()
	content, _ := ioutil.ReadFile(projectMainHtmlFilePath)
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()
}
