package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"

	"goforbroke1006/gfb-live-reload/pkg/runner"
)

func init() {
	flag.Parse()
}

func main() {

	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	if err := filepath.Walk("./", watchDir(watcher)); err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		done <- true
	}()

	fsEvents := make(chan bool, 1024)

	cmdRunner := runner.New(flag.Args())

	go func() {
		ticker := time.NewTicker(time.Second)
		reload := false
		for {
			select {
			case <-fsEvents:
				reload = true

			case <-ticker.C:
				if reload {
					cmdRunner.Reload()
					reload = false
				}
			}
		}
	}()

	go cmdRunner.Run()
	defer cmdRunner.Terminate()

	go func() {
		for {
			select {
			case _ = <-watcher.Events:
				fsEvents <- true
			}
		}
	}()

	<-done
}

func watchDir(watcher *fsnotify.Watcher) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {

		if fi.Mode().IsDir() {
			return watcher.Add(path)
		}

		return nil
	}
}
