package main

import (
	"github.com/lontten/lcloud/model"
	"github.com/lontten/lcloud/store"
	"github.com/lontten/lcloud/utils"
	"os"
	"path/filepath"
)

//遍历文件夹，获取文件信息，生成 file 变化event
func checkFileChangeEvent(path string) (model.SyncEvent, error) {
	e := model.NewSyncEvent()
	pathStore := store.FileHasPathStores[path]
	fileCheck := pathStore.FileCheck
	dirCheck := pathStore.DirCheck

	dir, err := os.ReadDir(path)
	if err != nil {
		return e, err
	}

	currentFileHashCheck := make(map[string]bool)
	currentDirHasCheck := make(map[string]bool)
	currentFilePathInfos := make(map[string]model.FileHashDto)
	currentFileTimeHasCheck := make(map[model.FileDto]string)

	for _, entry := range dir {
		filePath := filepath.Join(path, entry.Name())
		if entry.IsDir() { //文件夹

			currentDirHasCheck[filePath] = false //保存当前的文件夹列表
			_, ok := dirCheck[filePath]
			if ok {
				//文件夹已经存在,标记为true
				dirCheck[filePath] = true
				continue
			}

			//文件夹不存在，添加到add列表
			e.PushAddDirEvent(filePath)
			continue
		}

		//------------------文件-------------------
		info, err := entry.Info()
		if err != nil {
			return e, err
		}
		modTime := info.ModTime()

		currentFileHashCheck[filePath] = false //保存当前的文件列表

		_, ok := fileCheck[filePath]
		if ok {
			//文件夹已经存在,标记为true
			fileCheck[filePath] = true
		}

		fileDto := model.FileDto{
			Path:    filePath,
			ModTime: modTime,
		}
		hash, ok := store.GlobalFileTimeHashStore[fileDto]
		if ok {
			//文件已经存在，标记为true
			fileCheck[filePath] = true

			currentFilePathInfos[filePath] = model.FileHashDto{
				Path:    filePath,
				Hash:    hash,
				ModTime: modTime,
			}
			currentFileTimeHasCheck[fileDto] = hash

			continue
		}

		hash = utils.GetFileSHA256HashCode(filePath)
		//文件不存在，添加到add列表
		e.PushAddFileEvent(model.HashDto{
			Path: filePath,
			Hash: hash,
		})

		currentFilePathInfos[filePath] = model.FileHashDto{
			Path:    filePath,
			Hash:    hash,
			ModTime: modTime,
		}
		currentFileTimeHasCheck[fileDto] = hash
	}

	for s, b := range fileCheck {
		if !b {
			e.PushDelFileEvent(model.HashDto{
				Path: s,
				Hash: store.GlobalFileInfoStore[s].Hash,
			})
		}
	}

	for s, b := range fileCheck {
		if !b {
			e.PushDelDirEvent(s)
		}
	}

	store.FileHasPathStores[path] = store.PathStore{
		DirCheck:  currentDirHasCheck,
		FileCheck: currentFileHashCheck,
	}

	store.GlobalFileInfoStore = currentFilePathInfos
	store.GlobalFileTimeHashStore = currentFileTimeHasCheck

	return e, nil
}
