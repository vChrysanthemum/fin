package main

import (
	"inn/ui"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	logFile, _ := os.OpenFile("./log/select.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	log.SetOutput(logFile)

	s := `
    <html>
    <head>
        <title>Test</title>
    </head>
    <body colorfg="blue">
        <!--
		<select BorderLabel="测试">
		-->
        <select>
            <option value="shit"> 你好 </option>
            <option value="shit"> hello </option>
            <option> hi </option>
        </select>
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
