package model

//用户记录文件现在没有，就被删除了
//记录上一次的文件列表，用于下次，找出删除的文件路径，和新增的文件夹路径
var (
	FileHasCheck = make(map[string]bool)
	DirHasCheck  = make(map[string]bool)
)

//文件和 文件夹不同，还有修改，时间；用文件路径和时间，来记录；用于比较，是否有更新

//用户找出创建、修改的文件

//文件的列表
//文件路径+文件最后修改时间 》hash值
//用户忽略，未变化的文件，避免hash计算
var FileTimeHasCheck = make(map[FileDto]string)

//-------------------file-------------
var FilePathInfos = make(map[string]FileDto)
