package main

import (
	"fmt"
	"github.com/lontten/lcloud/model"
	"github.com/lontten/lcloud/utils"
	"os"
	"path/filepath"
)

func main() {
	e, err := checkFileModEvent("/Users/lontten/kk/")

	fmt.Println(e)
	fmt.Println(err)

}

//遍历文件夹，获取文件信息，生成 file 变化event
func checkFileModEvent(path string) (model.SyncEvent, error) {
	e := model.NewSyncEvent()

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
			_, ok := model.DirHasCheck[filePath]
			if ok {
				//文件夹已经存在,标记为true
				model.DirHasCheck[filePath] = true
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

		_, ok := model.FileHasCheck[filePath]
		if ok {
			//文件夹已经存在,标记为true
			model.FileHasCheck[filePath] = true
		}

		fileDto := model.FileDto{
			Path:    filePath,
			ModTime: modTime,
		}
		hash, ok := model.FileTimeHasCheck[fileDto]
		if ok {
			//文件已经存在，标记为true
			model.FileHasCheck[filePath] = true

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

	for s, b := range model.FileHasCheck {
		if !b {
			e.PushDelFileEvent(model.HashDto{
				Path: s,
				Hash: model.FilePathInfos[s].Hash,
			})
		}
	}

	for s, b := range model.DirHasCheck {
		if !b {
			e.PushDelDirEvent(s)
		}
	}

	model.FileHasCheck = currentFileHashCheck
	model.DirHasCheck = currentDirHasCheck
	model.FilePathInfos = currentFilePathInfos
	model.FileTimeHasCheck = currentFileTimeHasCheck

	return e, nil
}
