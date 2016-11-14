package main

import (
	"fmt"
	"in/ui"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	ui.PrepareUI()
	target := os.Args[1]

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	logFile, _ := os.OpenFile(fmt.Sprintf("./log/%s.log", target), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	log.SetOutput(logFile)

	content, _ := ioutil.ReadFile(fmt.Sprintf("./test/html/%s.html", target))
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()
}
