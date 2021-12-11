package main

import (
	"fmt"
	"github.com/lontten/lcloud/model"
	"github.com/lontten/lcloud/utils"
	"os"
	"path/filepath"
)

func main() {
	e, err := checkFileModEvent("/Users/lontten/")

	fmt.Println(e)
	fmt.Println(err)

}

//遍历文件夹，获取文件信息，生成 file 变化event
func checkFileModEvent(path string) (model.SyncEvent, error) {
	e := model.SyncEvent{}

	dir, err := os.ReadDir(path)
	if err != nil {
		return e, err
	}

	currentFileFlags := make(map[string]bool)
	currentDirFlags := make(map[string]bool)
	currentPathFileStores := make(map[string]model.FileDto)

	for _, entry := range dir {
		filePath := filepath.Join(path, entry.Name())
		if entry.IsDir() { //文件夹

			currentDirFlags[filePath] = false //保存当前的文件夹列表
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

		currentFileFlags[filePath] = false //保存当前的文件列表
		currentPathFileStores[filePath] = model.FileDto{
			Path:    filePath,
			ModTime: modTime,
		}

		_, ok := model.FileHasCheck[filePath]
		if ok {
			//文件夹已经存在,标记为true
			model.FileHasCheck[filePath] = true
			continue
		}

		_, ok = model.FileTimeHasCheck[model.FileDto{
			Path:    filePath,
			ModTime: modTime,
		}]
		if ok {
			//文件已经存在，标记为true
			model.FileHasCheck[filePath] = true
			continue
		}
		//文件不存在，添加到add列表
		e.PushAddFileEvent(model.FileHashDto{
			Path:    filePath,
			ModTime: modTime,
			Hash:    utils.GetFileSHA256HashCode(filePath),
		})
	}

	for s, b := range model.FileHasCheck {
		if !b {
			e.PushDelFileEvent(model.FilePathInfos[s])
		}
	}

	for s, b := range model.DirHasCheck {
		if !b {
			e.PushDelDirEvent(s)
		}
	}

	model.FileHasCheck = currentFileFlags
	model.DirHasCheck = currentDirFlags
	model.FilePathInfos = currentPathFileStores

	return e, nil
}
