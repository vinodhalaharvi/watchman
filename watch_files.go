package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
)

func NewFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		var wg sync.WaitGroup
		for {
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						fmt.Printf("%s\n", "Error")
						//return
					}
					log.Println("event:", event)
					if event.Op&fsnotify.Write == fsnotify.Write {
						log.Println("modified file:", event.Name)
					}
					if event.Op&fsnotify.Create == fsnotify.Create {
						fileName := event.Name
						classify(&fileName)
						log.Println("creeated file:", event.Name)
						return
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						fmt.Printf("%s\n", "Error")
						//return
					}
					log.Println("error:", err)
				}
			}()
		}
		wg.Wait()
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
