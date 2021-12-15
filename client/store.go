package main

import (
	"path/filepath"
	"strings"
	"sync"
)

//需要同步的文件
var (
	//需要从服务器拉取的文件hash
	NeedPullFileStore = make(map[string]struct{})
	//需要上传到服务器的文件hash
	NeedPushFileStore = make(map[string]struct{})

	//文件hash；文件地址
	LocalStoreFileHash = make(map[string]string)

	//需要处理的从服务器拉取的event
	LocalEventStore = make(map[string]string)

	NeedCheckPathStoreMutex sync.Mutex
	//需要检查文件变动的路径
	NeedCheckPathStore = NeedCheckPath{}
)

type NeedCheckPath map[string]struct{}

//添加需要检查的路径,路径存在包含关系时，保留大的路径
func (p *NeedCheckPath) Push(path string) {
	path = filepath.Join("/", path)
	NeedCheckPathStoreMutex.Lock()
	for s := range *p {
		if strings.HasPrefix(path, s) {
			delete(*p, s)
			(*p)[path] = struct{}{}
			NeedCheckPathStoreMutex.Unlock()
			return
		}
		if strings.HasPrefix(s, path) {
			NeedCheckPathStoreMutex.Unlock()
			return
		}
	}
	(*p)[path] = struct{}{}
	NeedCheckPathStoreMutex.Unlock()
}
