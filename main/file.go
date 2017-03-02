package main

import (
	"fin/script"
	"fin/ui"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/net/html"

	"github.com/gizak/termui"
)

// 编辑器渲染好后加载文件
func hookersAfterFirstUIRender(arg interface{}) {
	var err error
	filePath := arg.(string)
	defaultEditor, ok := ui.GCurrentRenderPage.IDToNodeMap["EditorDefault"]
	if false == ok {
		log.Println(filePath)
		return
	}

	defaultEditor.UIBlock.Height = termui.TermHeight() - 1

	defaultTab, ok := ui.GCurrentRenderPage.IDToNodeMap["TabDefault"]
	defaultTab.ParseAttribute([]html.Attribute{
		html.Attribute{"", "label", filePath},
		html.Attribute{"", "name", "default-" + filePath},
	})

	err = defaultEditor.Data.(*ui.NodeEditor).Editor.LoadFile(filePath)
	if nil != err {
		log.Println(err)
	}

	ui.GCurrentRenderPage.ReRender()
}

func loadFile(filePath string) {
	ui.GlobalOption.ProjectPath = GlobalResBaseDir
	script.GlobalOption.ProjectPath = GlobalResBaseDir

	projectMainHtmlFilePath := filepath.Join(GlobalEditorBaseDir, "main.html")
	if _, err := os.Stat(projectMainHtmlFilePath); os.IsNotExist(err) {
		log.Println("Project is not existed.")
		return
	}

	logFile, err := os.OpenFile(filepath.Join(GlobalEditorBaseDir, "main.log"),
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	AssertErrIsNil(err)
	log.SetOutput(logFile)

	ui.PrepareUI()
	content, _ := ioutil.ReadFile(projectMainHtmlFilePath)
	page, err := ui.Parse(string(content))
	if nil != err {
		panic(err)
	}

	page.HookersAfterFirstUIRender = append(page.HookersAfterFirstUIRender, ui.Hooker{
		Arg: filePath,
		Do:  hookersAfterFirstUIRender,
	})
	page.Render()
	page.Serve()
}
