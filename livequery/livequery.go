package livequery

import (
	"strings"

	"github.com/lfq7413/tomato/config"
)

// M ...
type M map[string]interface{}

// LiveQuery 接收指定类的对象保存与对象删除的通知，发送对象数据到发布者，由发布者通知订阅者，订阅者实时接收数据
type LiveQuery struct {
	classNames         []string
	liveQueryPublisher *cloudCodePublisher
}

// TLiveQuery ...
var TLiveQuery *LiveQuery

func init() {
	classeNames := strings.Split(config.TConfig.LiveQueryClasses, "|")
	pubType := config.TConfig.PublisherType
	pubURL := config.TConfig.PublisherURL
	TLiveQuery = NewLiveQuery(classeNames, pubType, pubURL)
}

// NewLiveQuery 初始化 LiveQuery
// classNames 支持的类列表
// pubType 发布者类型
// pubURL 发布者的 URL
func NewLiveQuery(classNames []string, pubType, pubURL string) *LiveQuery {
	liveQuery := &LiveQuery{}
	if classNames == nil && len(classNames) == 0 {
		liveQuery.classNames = []string{}
	} else {
		liveQuery.classNames = classNames
	}
	liveQuery.liveQueryPublisher = newCloudCodePublisher(pubType, pubURL)

	return liveQuery
}

// OnAfterSave 保存对象之后调用
func (l *LiveQuery) OnAfterSave(className string, currentObject, originalObject map[string]interface{}) {
	if l.HasLiveQuery(className) == false {
		return
	}
	req := l.makePublisherRequest(currentObject, originalObject)
	l.liveQueryPublisher.onCloudCodeAfterSave(req)
}

// OnAfterDelete 删除对象之后调用
func (l *LiveQuery) OnAfterDelete(className string, currentObject, originalObject map[string]interface{}) {
	if l.HasLiveQuery(className) == false {
		return
	}
	req := l.makePublisherRequest(currentObject, originalObject)
	l.liveQueryPublisher.onCloudCodeAfterDelete(req)
}

// HasLiveQuery 是否有对应的 className
func (l *LiveQuery) HasLiveQuery(className string) bool {
	for _, n := range l.classNames {
		if n == className {
			return true
		}
	}
	return false
}

// makePublisherRequest 组装待发布的消息，格式如下
// {
// 	"object": {...},
// 	"original": {...}
// }
func (l *LiveQuery) makePublisherRequest(currentObject, originalObject M) M {
	req := M{
		"object": currentObject,
	}
	if currentObject != nil {
		req["original"] = originalObject
	}
	return req
}
