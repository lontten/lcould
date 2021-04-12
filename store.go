package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type PathStore struct {
	Path     string
	FilePath []string
	DirPath  []string
}

type HashStore struct {
	Hash256 string
	FiArr   []Finfo
}

var (
	PathStores = make([]PathStore, 0)
	HashStores = make([]HashStore, 0)
)

//文件信息
type Finfo struct {
	Path    string
	ModTime time.Time
}

func (conf *HashStore) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}
