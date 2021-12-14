package main

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

	//需要检查文件变动的路径
	NeedCheckPathStore = make(map[string]struct{})
)
