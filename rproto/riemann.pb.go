// Code generated by protoc-gen-go.
// source: riemann.proto
// DO NOT EDIT!

/*
Package rproto is a generated protocol buffer package.

It is generated from these files:
	riemann.proto

It has these top-level messages:
	State
	Event
	Query
	Msg
	Attribute
*/
package rproto

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Deprecated; state was used by early versions of the protocol, but not any
// more.
type State struct {
	Time             *int64   `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	State            *string  `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Service          *string  `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Host             *string  `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	Description      *string  `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Once             *bool    `protobuf:"varint,6,opt,name=once" json:"once,omitempty"`
	Tags             []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	Ttl              *float32 `protobuf:"fixed32,8,opt,name=ttl" json:"ttl,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}

func (m *State) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *State) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *State) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

func (m *State) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *State) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *State) GetOnce() bool {
	if m != nil && m.Once != nil {
		return *m.Once
	}
	return false
}

func (m *State) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *State) GetTtl() float32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

type Event struct {
	Time             *int64       `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	State            *string      `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Service          *string      `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Host             *string      `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	Description      *string      `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Tags             []string     `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	Ttl              *float32     `protobuf:"fixed32,8,opt,name=ttl" json:"ttl,omitempty"`
	Attributes       []*Attribute `protobuf:"bytes,9,rep,name=attributes" json:"attributes,omitempty"`
	MetricSint64     *int64       `protobuf:"zigzag64,13,opt,name=metric_sint64" json:"metric_sint64,omitempty"`
	MetricD          *float64     `protobuf:"fixed64,14,opt,name=metric_d" json:"metric_d,omitempty"`
	MetricF          *float32     `protobuf:"fixed32,15,opt,name=metric_f" json:"metric_f,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}

func (m *Event) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Event) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *Event) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

func (m *Event) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Event) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Event) GetTtl() float32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *Event) GetAttributes() []*Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Event) GetMetricSint64() int64 {
	if m != nil && m.MetricSint64 != nil {
		return *m.MetricSint64
	}
	return 0
}

func (m *Event) GetMetricD() float64 {
	if m != nil && m.MetricD != nil {
		return *m.MetricD
	}
	return 0
}

func (m *Event) GetMetricF() float32 {
	if m != nil && m.MetricF != nil {
		return *m.MetricF
	}
	return 0
}

type Query struct {
	String_          *string `protobuf:"bytes,1,opt,name=string" json:"string,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}

func (m *Query) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

type Msg struct {
	Ok               *bool    `protobuf:"varint,2,opt,name=ok" json:"ok,omitempty"`
	Error            *string  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	States           []*State `protobuf:"bytes,4,rep,name=states" json:"states,omitempty"`
	Query            *Query   `protobuf:"bytes,5,opt,name=query" json:"query,omitempty"`
	Events           []*Event `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}

func (m *Msg) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *Msg) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *Msg) GetStates() []*State {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *Msg) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Msg) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Attribute struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}

func (m *Attribute) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Attribute) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
}