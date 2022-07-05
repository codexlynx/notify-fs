package watcher

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

type Fs struct {
	OnCreateHook func(filePath string)
	watcher      *fsnotify.Watcher
}

func (watcherFs *Fs) Watch() {
	for {
		select {
		case event, _ := <-watcherFs.watcher.Events:
			if event.Op == fsnotify.Create {
				go watcherFs.OnCreateHook(event.Name)
			}
		case err, _ := <-watcherFs.watcher.Errors:
			log.Println(err)
		}
	}
}

func (watcherFs *Fs) Close() {
	err := watcherFs.watcher.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func NewFs(directory string, onCreateHook func(filePath string)) (*Fs, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	err = watcher.Add(directory)
	if err != nil {
		return nil, err
	}
	fsWatcher := &Fs{
		OnCreateHook: onCreateHook,
		watcher:      watcher,
	}
	return fsWatcher, nil
}
