package obsfg

import (
	"fmt"
)

// OBSObject 定义OBS对象的属性
type OBSObject struct {
	Key       string `json:"key"`
	Size      int64  `json:"size"`
	ETag      string `json:"eTag"`
	VersionId string `json:"versionId"`
	Sequencer string `json:"sequencer"`
}

// OBSBucket 定义OBS桶的属性
type OBSBucket struct {
	Bucket        string `json:"bucket"`
	Name          string `json:"name"`
	OwnerIdentity struct {
		ID string `json:"ID"`
	} `json:"ownerIdentity"`
}

// OBSData 定义OBS事件的数据部分
type OBSData struct {
	OBS struct {
		Bucket          OBSBucket `json:"bucket"`
		Object          OBSObject `json:"object"`
		Version         string    `json:"Version"`
		ConfigurationID string    `json:"configurationId"`
	} `json:"obs"`
	EventVersion      string `json:"eventVersion"`
	EventSource       string `json:"eventSource"`
	EventTime         string `json:"eventTime"`
	EventName         string `json:"eventName"`
	EventRegion       string `json:"eventRegion"`
	RequestParameters struct {
		SourceIPAddress string `json:"sourceIPAddress"`
	} `json:"requestParameters"`
	ResponseElements struct {
		XObsRequestID string `json:"x-obs-request-id"`
		XObsID2       string `json:"x-obs-id-2"`
	} `json:"responseElements"`
	UserIdentity struct {
		ID string `json:"ID"`
	} `json:"userIdentity"`
}

// OBSTriggerEvent 定义OBS触发器事件的完整结构
type OBSTriggerEvent struct {
	Data            OBSData `json:"data"`
	DataContentType string  `json:"datacontenttype"`
	Subject         string  `json:"subject"`
	SpecVersion     string  `json:"specversion"`
	ID              string  `json:"id"`
	Source          string  `json:"source"`
	Time            string  `json:"time"`
	Type            string  `json:"type"`
	TTL             string  `json:"ttl"`
	DataSchema      string  `json:"dataschema"`

	// 为了保持兼容性，保留原始字段
	Bucket string `json:"bucket"`
	Object string `json:"object"`
}

func (e *OBSTriggerEvent) String() string {
	// 获取真正的bucket和key值，优先使用新结构
	var bucket, objectKey string

	if e.Data.OBS.Bucket.Bucket != "" {
		bucket = e.Data.OBS.Bucket.Bucket
	} else {
		bucket = e.Bucket
	}

	if e.Data.OBS.Object.Key != "" {
		objectKey = e.Data.OBS.Object.Key
	} else {
		objectKey = e.Object
	}

	return fmt.Sprintf(`{ "bucket":"%v","object":"%v", "subject":"%v", "id":"%v" }`,
		bucket, objectKey, e.Subject, e.ID)
}
