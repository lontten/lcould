package main

import "time"

type DirDto struct {
	Path    string    `json:"path"`
	ModTime time.Time `json:"mod_time"`
}
