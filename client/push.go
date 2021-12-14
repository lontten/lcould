package main

import "github.com/lontten/lcloud/model"

//向服务器上传本地文件变化event
func pushLocalEvent(event model.SyncEvent) (PushEventResp, error) {

	return PushEventResp{}, nil
}

//向服务器上传，没有的文件；根据hash
func doPushFileToServer() {

}

type PushEventResp struct {
	//想要推送的文件的hash
	NeedPushHash []string
	//需要接收的文件变动event
	Event model.SyncEvent
}