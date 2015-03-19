package rproto

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

type AttributeJson struct {
	Key   string
	Value string
}

type EventJson struct {
	Time          int64
	State         string
	Service       string
	Host          string
	Description   string
	Tags          []string
	TTL           float32
	Attributes    AttributeJson
	Metric_sint64 int64
	Metric_d      float64
	Metric_f      float32
}

func JsonToEvent(jsonEventText []byte) *Event {

	eJ := new(EventJson)

	err := json.Unmarshal(jsonEventText, eJ)

	if err != nil {
		panic(err)
	}

	eP := &Event{
		Time:        proto.Int64(eJ.Time),
		State:       proto.String(eJ.State),
		Service:     proto.String(eJ.Service),
		Host:        proto.String(eJ.Host),
		Description: proto.String(eJ.Description),
		//tags??
		Ttl: proto.Float32(eJ.TTL),
		//attributes?
		MetricSint64: proto.Int64(eJ.Metric_sint64),
		MetricD:      proto.Float64(eJ.Metric_d),
		MetricF:      proto.Float32(eJ.Metric_f),
	}

	return eP
}
