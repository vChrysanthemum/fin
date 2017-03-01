package main

import (
	"fin/script"
	"fin/ui"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

var (
	GlobalEditorBaseDir string
	GlobalResBaseDir    string
)

func AssertErrIsNil(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Project filepath is needed")
		return
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 8192)
		for _ = range sigChan {
			length := runtime.Stack(stacktrace, true)
			log.Println(string(stacktrace[:length]))
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)

	GlobalEditorBaseDir = filepath.Join(os.Getenv("HOME"), ".fin", "res", "editor")
	GlobalResBaseDir = filepath.Join(os.Getenv("HOME"), ".fin", "res")
	ui.GlobalOption.ResBaseDir = GlobalResBaseDir
	script.GlobalOption.ResBaseDir = GlobalResBaseDir

	/*
		projectPath, err := filepath.Abs(os.Args[1])
		if nil != err {
			log.Println(err)
			return
		}
	*/
	projectPath := os.Args[1]

	projectFileStat, err := os.Stat(projectPath)
	if nil == err && true == projectFileStat.IsDir() {
		loadProject(projectPath)
	} else {
		loadFile(projectPath)
	}
}
