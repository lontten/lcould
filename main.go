package main

import (
	"fmt"
	"github.com/lontten/lcloud/utils"
	"os"
	"path/filepath"
)

//处理del
func sFolder(arr []FileModDto) {

}

func main() {
	add, del, err := mapLocalDir("/Users/lontten/")
	sFolder(add)

	fmt.Println(err)

}

//文件同步对比
type FileSync struct {
	Hash256 string
	Fi      Finfo
}

//遍历文件夹，获取文件信息，生成 HashStore，FolderStore
func mapLocalDir(path string) (add, del []FileModDto, err error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	currentFileFlags := make(map[string]bool)
	currentFolerFlags := make(map[string]bool)

	for _, entry := range dir {
		filePath := filepath.Join(path, entry.Name())
		if entry.IsDir() { //文件夹
			currentFolerFlags[filePath] = false //保存当前的文件夹列表

			_, ok := FolderStores[filePath]
			if ok {
				//文件夹已经存在,标记为true
				FolderFlags[filePath] = true
			} else {
				//文件夹不存在，添加到add列表
				add = append(add, FileModDto{
					Path: filePath,
				})
			}
		} else { //文件
			currentFileFlags[filePath] = false //保存当前的文件列表

			info, err := entry.Info()
			if err != nil {
				return nil, nil, err
			}
			modTime := info.ModTime()
			_, ok := FileStorss[Finfo{
				Path:    filePath,
				ModTime: modTime,
			}]
			if ok {
				//文件已经存在，标记为true
				FileFlags[filePath] = true
			} else {
				//文件不存在，添加到add列表
				add = append(add, FileModDto{
					Path: filePath,
					Hash: utils.GetFileSHA256HashCode(filePath),
				})

			}
		}
	}

	for s, b := range FolderFlags {
		if !b {
			del = append(del, FileModDto{
				Path: s,
				Hash: "folder",
			})
		}
	}

	for s, b := range FolderFlags {
		if !b {
			del = append(del, FileModDto{
				Path: s,
				Hash: "",
			})
		}
	}

	return
}
