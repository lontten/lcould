package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

//一个文件夹 下的 文件和文件夹的列表
type PathStore struct {
	Path     string
	FilePath []string
	DirPath  []string
}

// hash256,文件指纹，
//在不同的文件路径下，会有内容相同的文件，用来保证，相同文件只会被存储一次
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
