package main

import (
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

	content, _ := ioutil.ReadFile(filepath.Join(GlobalResBaseDir, "project/travller/main.html"))
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()

}
