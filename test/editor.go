package main

import (
	"inn/ui"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	logFile, _ := os.OpenFile("./log/editor.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	log.SetOutput(logFile)

	s := `
   <html>
    <head>
        <title>Test</title>
    </head>
    <body colorfg="blue">
        <editor></editor>
    </body>
    </html>
    `
	page, err := ui.Parse(s)
	if nil != err {
		panic(err)
	}
	page.Render()
	page.Serve()
}
