package model

//同步事件
type SyncEvent struct {
	//要创建的文件列表
	AddFileEvent []FileHashDto
	//要删除的文件列表
	DelFileEvent []FileDto
	//要创建的目录列表
	AddDirEvent []string
	//要删除的目录列表
	DelDirEvent []string
}

func (e SyncEvent) PushAddFileEvent(d ...FileHashDto) {

}
func (e SyncEvent) PushDelFileEvent(d ...FileDto) {

}
func (e SyncEvent) PushAddDirEvent(d ...string) {

}
func (e SyncEvent) PushDelDirEvent(d ...string) {

}
