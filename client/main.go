package main

import "github.com/lontten/lcloud/model"

func main() {

}

//用户保存一个文件夹下的所有文件，和文件夹
type FileStore struct {
	Path string
	//上一次文件夹的列表，bool用来下一次标记未发现（删除）的文件夹
	Dirs  map[string]bool
	Files map[model.FileDto]bool
}
