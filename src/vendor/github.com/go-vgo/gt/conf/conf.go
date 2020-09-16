// Copyright 2017 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/gt/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package conf

import (
	"log"
	"sync"

	// "time"

	"github.com/fsnotify/fsnotify"
)

var (
	// config     Config
	confLock = new(sync.RWMutex)
)

// NewWatcher new fsnotify watcher
func NewWatcher(paths string, config interface{}) {
	Watch(paths, config)
}

// Watch new fsnotify watcher
func Watch(paths string, config interface{}) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("fsnotify.NewWatcher() error: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("fsnotify watcher events: ", event)

				// if event.Op&fsnotify.Chmod == fsnotify.Chmod {
				// 	log.Println("fsnotify watcher.Events: ignore CHMOD event: ", event)
				// 	continue
				// }

				if event.Op&fsnotify.Write == fsnotify.Write {
					// log.Println("modified file: ", event.Name)
					err := Init(paths, config)
					if err == nil {
						log.Println("fsnotify watch config: ", config)
					}
				}
			case err := <-watcher.Errors:
				log.Println("fsnotify watcher.Errors error: ", err)
			}
		}
	}()

	err = watcher.Add(paths)
	if err != nil {
		log.Fatal("fsnotify watcher.Add: ", err)
	}
	<-done
}
