package main

func main() {
	//1.检查本地变化，同步到服务器
	//2.拉取需要同步的变化，合并到本地的eventStore
	//3.拉取需要上传的文件hash，合并到本地的 needPullHash
	go doCheck()

	//4.依据本地eventStore，修改本地文件；本地没有的文件hash 添加对 needPushHash
	go doSyncFromLocalEventStore()
	//5.上传需要同步的文件
	go doPullFileFormServer()
	//6.下载需要同步的文件
	go doPushFileToServer()

}
