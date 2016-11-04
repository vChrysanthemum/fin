package main

import (
	"inn/ui"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	logFile, _ := os.OpenFile("./log/table.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	log.SetOutput(logFile)

	s := `
   <html>
    <head>
        <title>Test</title>
    </head>
    <body colorfg="blue">
		<table>
			<tr>
				<td>
					<par>
						asd;lfkasfkja;sdlf
					</par>
				</td>
			</tr>
			<tr>
				<td>
					<par border="true">
						asd;lfkasfkja;sdlf
					</par>
				</td>
			</tr>
			<tr>
				<td>
					<select>
						<option value="1">test1</option>
						<option value="1">test2</option>
					</select>
				</td>
				<td>
					<editor></editor>
				</td>
			</tr>
		</table>
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
