package model

import "time"

//文件信息
type FileDto struct {
	//文件的路径
	Path string
	//文件的创建时间，路径加时间，空间坐标+时间坐标，确定一个文件的唯一性
	ModTime time.Time
}

//文件信息
type FileHashDto struct {
	//文件的路径
	Path string
	//文件的创建时间，路径加时间，空间坐标+时间坐标，确定一个文件的唯一性
	ModTime time.Time
	Hash    string
}
