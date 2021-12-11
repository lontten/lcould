package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

var (
	FolderStores FolderStore
	FileStorss   FileStore

	FileFlags   FileFlag
	FolderFlags FileFlag
)

//文件夹的列表
type FolderStore map[string]bool

// hash256,文件指纹，
//在不同的文件路径下，会有内容相同的文件，用来保证，相同文件只会被存储一次
//string > 文件hash值，
//fiarr > 文件信息
type HashStore map[string]FiArr

//文件路径+文件最后修改时间 》hash值
//用户忽略，未变化的文件，避免hash计算
type FileStore map[Finfo]string

//记录上一次的文件列表，用于下次，找出删除的文件路径
type FileFlag map[string]bool

func (f *FileFlag) Reset() {
	for s := range *f {
		(*f)[s] = false
	}
}

type FiArr map[Finfo]struct{}

func (a *FiArr) Add(finfo Finfo) bool {
	_, ok := (*a)[finfo]
	if !ok {
		(*a)[finfo] = struct{}{}
		return true
	}
	return false
}

//文件信息
type Finfo struct {
	Path    string
	ModTime time.Time
}

//文件夹直接上传，修改
//遍历hash，有hash，只修改文件路径列表；没有hash，上传hash对应的文件后，再修改文件路径列表
type FileModDto struct {
	//文件路径
	Path string
	//文件hash，文件夹，此项为 ""
	Hash string
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

func (conf *FolderStore) String() string {
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
