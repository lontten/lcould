package store

import (
	"github.com/lontten/lcloud/model"
)

//一个文件夹夹所有的fielstore
var FileHasPathStores = make(map[string]PathStore)

type PathStore struct {
	//文件夹路径
	DirCheck map[string]bool

	//文件路径
	FileCheck map[string]bool
}

type FileTimeHashDto map[model.FileDto]string

//文件路径+文件最后修改时间
var GlobalFileTimeHashStore = FileTimeHashDto{}

func (s FileTimeHashDto) Add(d model.FileDto, hash string) {
	for dto := range GlobalFileTimeHashStore {
		if dto.Path == d.Path {
			delete(GlobalFileTimeHashStore, d)
			break
		}
	}
	GlobalFileTimeHashStore[d] = hash
}

func (s FileTimeHashDto) Del(path string) {
	for dto := range GlobalFileTimeHashStore {
		if dto.Path == path {
			delete(GlobalFileTimeHashStore, dto)
			return
		}
	}
}

var GlobalFileInfoStore = make(map[string]model.FileHashDto)
