package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/go-fsnotify/fsnotify"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Usage: " + filepath.Base(os.Args[0]) + " [directory]")
		return
	}
	var dir string
	if len(os.Args) == 2 {
		dir = os.Args[1]
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	type Event struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}
	enc := json.NewEncoder(os.Stdout)
	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			switch {
			case event.Op&fsnotify.Write == fsnotify.Write:
				enc.Encode(&Event{Type: "write", Name: event.Name})
			case event.Op&fsnotify.Create == fsnotify.Create:
				enc.Encode(&Event{Type: "create", Name: event.Name})
			case event.Op&fsnotify.Remove == fsnotify.Remove:
				enc.Encode(&Event{Type: "remove", Name: event.Name})
			case event.Op&fsnotify.Rename == fsnotify.Rename:
				enc.Encode(&Event{Type: "rename", Name: event.Name})
			case event.Op&fsnotify.Chmod == fsnotify.Chmod:
				enc.Encode(&Event{Type: "rename", Name: event.Name})
			}
		case err := <-watcher.Errors:
			log.Fatal("error: ", err)
		}
	}
}
