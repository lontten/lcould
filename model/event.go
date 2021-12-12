package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//用户发送本地文件变化，接收服务器发来的文件变化

type HashDto struct {
	Hash string `json:"hash"`
	Path string `json:"path"`
}

type Event struct {
	// 删除的文件路径
	Del map[string]struct{}
	// 新增的文件路径
	Add map[string]struct{}
}

//同步事件
type SyncEvent struct {
	FileEvent map[string]Event

	DirEvent Event
}

func NewSyncEvent() SyncEvent {
	event := Event{
		Del: map[string]struct{}{},
		Add: map[string]struct{}{},
	}
	return SyncEvent{
		FileEvent: make(map[string]Event),
		DirEvent:  event,
	}
}

func (e *SyncEvent) PushAddFileEvent(arr ...HashDto) {
	for _, d := range arr {
		event, ok := e.FileEvent[d.Hash]
		if !ok {
			event = Event{
				Del: map[string]struct{}{},
				Add: map[string]struct{}{},
			}
			event.Add[d.Path] = struct{}{}
			e.FileEvent[d.Hash] = event
		} else {
			event.Add[d.Path] = struct{}{}
			e.FileEvent[d.Hash] = event
		}
	}
}
func (e *SyncEvent) PushDelFileEvent(arr ...HashDto) {
	for _, d := range arr {
		event, ok := e.FileEvent[d.Hash]
		if !ok {
			event = Event{
				Del: map[string]struct{}{},
				Add: map[string]struct{}{},
			}
			event.Del[d.Path] = struct{}{}
			e.FileEvent[d.Hash] = event
		} else {
			event.Del[d.Path] = struct{}{}
			e.FileEvent[d.Hash] = event
		}
	}
}

func (e *SyncEvent) PushAddDirEvent(arr ...string) {
	for _, s := range arr {
		e.DirEvent.Add[s] = struct{}{}
	}
}
func (e *SyncEvent) PushDelDirEvent(arr ...string) {
	for _, s := range arr {
		e.DirEvent.Del[s] = struct{}{}
	}
}

// go toString
func (conf *SyncEvent) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}
