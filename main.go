package main

import (
	"github.com/lontten/lcloud/utils"
	"os"
	"path/filepath"
)

func main() {
}

//文件同步对比
type FileSync struct {
	Hash256 string
	Fi      Finfo
}

//遍历文件夹，获取文件信息，生成 HashStore，PathStore
func mapLocalDir(path string, updPath []string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	dirs := make([]string, 0)
	files := make([]string, 0)

	for _, entry := range dir {
		fileName := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			dirs = append(dirs, fileName)
		} else {
			files = append(files, fileName)

			info, err := entry.Info()
			if err != nil {
				return err
			}
			modTime := info.ModTime()
			code := utils.GetFileSHA256HashCode(fileName)
			fInfo := Finfo{
				Path:    fileName,
				ModTime: modTime,
			}
			addFile(code, fInfo)
		}
	}

	for i, pathStore := range PathStores {
		if pathStore.Path == path {
			PathStores[i].DirPath = dirs
			PathStores[i].FilePath = files
			return nil
		}
	}
	//没有匹配到路径，添加
	store := PathStore{
		Path:     path,
		FilePath: files,
		DirPath:  dirs,
	}
	PathStores = append(PathStores, store)
	return nil
}

//有相同的直接返回fals，没有添加后，返回true
func addFile(code string, fInfo Finfo) bool {
	for i, store := range HashStores {
		//已存在相同的指纹，添加进数组
		if store.Hash256 == code {
			for _, finfo := range store.FiArr {
				if finfo.Path == fInfo.Path {
					return false
				}
			}
			HashStores[i].FiArr = append(store.FiArr, fInfo)
			return true
		}
	}

	//未存在相同的指纹，直接创建
	fis := make([]Finfo, 1)
	fis[0] = fInfo
	hf := HashStore{Hash256: code, FiArr: fis}
	HashStores = append(HashStores, hf)
	return true
}
