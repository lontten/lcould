package main

import "github.com/lontten/lcloud/model"

//把从服务器同步来的event合并到本地eventStore中
func AddSyncEvent2LocalEvent(event model.SyncEvent)  {

}

//从服务器拉取本地没有的文件
func doPullFileFormServer() {

}

//以eventStore为依据，对本地文件进行修改 for true
func doSyncFromLocalEventStore() {
}
