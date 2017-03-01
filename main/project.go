package main

import (
	"fin/script"
	"fin/ui"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func loadProject(projectPath string) {
	ui.GlobalOption.ProjectPath = projectPath
	script.GlobalOption.ProjectPath = projectPath

	projectMainHtmlFilePath := filepath.Join(projectPath, "main.html")
	if _, err := os.Stat(projectMainHtmlFilePath); os.IsNotExist(err) {
		log.Println("Project is not existed.")
		return
	}

	logFile, err := os.OpenFile(filepath.Join(projectPath, "main.log"),
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	AssertErrIsNil(err)
	log.SetOutput(logFile)

	ui.PrepareUI()
	content, _ := ioutil.ReadFile(projectMainHtmlFilePath)
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()
}
