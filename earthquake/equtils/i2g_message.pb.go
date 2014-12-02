// Code generated by protoc-gen-go.
// source: i2g_message.proto
// DO NOT EDIT!

/*
Package equtils is a generated protocol buffer package.

It is generated from these files:
	i2g_message.proto

It has these top-level messages:
	I2GMsgReq_Event_FuncCall
	I2GMsgReq_Event
	I2GMsgReq_Initiation
	I2GMsgReq
	I2GMsgRsp
*/
package equtils

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type I2GMsgReq_Event_Type int32

const (
	I2GMsgReq_Event_FUNC_CALL I2GMsgReq_Event_Type = 1
)

var I2GMsgReq_Event_Type_name = map[int32]string{
	1: "FUNC_CALL",
}
var I2GMsgReq_Event_Type_value = map[string]int32{
	"FUNC_CALL": 1,
}

func (x I2GMsgReq_Event_Type) Enum() *I2GMsgReq_Event_Type {
	p := new(I2GMsgReq_Event_Type)
	*p = x
	return p
}
func (x I2GMsgReq_Event_Type) String() string {
	return proto.EnumName(I2GMsgReq_Event_Type_name, int32(x))
}
func (x *I2GMsgReq_Event_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(I2GMsgReq_Event_Type_value, data, "I2GMsgReq_Event_Type")
	if err != nil {
		return err
	}
	*x = I2GMsgReq_Event_Type(value)
	return nil
}

type I2GMsgReq_Type int32

const (
	I2GMsgReq_EVENT      I2GMsgReq_Type = 1
	I2GMsgReq_INITIATION I2GMsgReq_Type = 2
)

var I2GMsgReq_Type_name = map[int32]string{
	1: "EVENT",
	2: "INITIATION",
}
var I2GMsgReq_Type_value = map[string]int32{
	"EVENT":      1,
	"INITIATION": 2,
}

func (x I2GMsgReq_Type) Enum() *I2GMsgReq_Type {
	p := new(I2GMsgReq_Type)
	*p = x
	return p
}
func (x I2GMsgReq_Type) String() string {
	return proto.EnumName(I2GMsgReq_Type_name, int32(x))
}
func (x *I2GMsgReq_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(I2GMsgReq_Type_value, data, "I2GMsgReq_Type")
	if err != nil {
		return err
	}
	*x = I2GMsgReq_Type(value)
	return nil
}

type I2GMsgRsp_Result int32

const (
	I2GMsgRsp_ACK   I2GMsgRsp_Result = 1
	I2GMsgRsp_ERROR I2GMsgRsp_Result = 2
	I2GMsgRsp_END   I2GMsgRsp_Result = 3
)

var I2GMsgRsp_Result_name = map[int32]string{
	1: "ACK",
	2: "ERROR",
	3: "END",
}
var I2GMsgRsp_Result_value = map[string]int32{
	"ACK":   1,
	"ERROR": 2,
	"END":   3,
}

func (x I2GMsgRsp_Result) Enum() *I2GMsgRsp_Result {
	p := new(I2GMsgRsp_Result)
	*p = x
	return p
}
func (x I2GMsgRsp_Result) String() string {
	return proto.EnumName(I2GMsgRsp_Result_name, int32(x))
}
func (x *I2GMsgRsp_Result) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(I2GMsgRsp_Result_value, data, "I2GMsgRsp_Result")
	if err != nil {
		return err
	}
	*x = I2GMsgRsp_Result(value)
	return nil
}

type I2GMsgReq_Event_FuncCall struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *I2GMsgReq_Event_FuncCall) Reset()         { *m = I2GMsgReq_Event_FuncCall{} }
func (m *I2GMsgReq_Event_FuncCall) String() string { return proto.CompactTextString(m) }
func (*I2GMsgReq_Event_FuncCall) ProtoMessage()    {}

func (m *I2GMsgReq_Event_FuncCall) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type I2GMsgReq_Event struct {
	Type             *I2GMsgReq_Event_Type     `protobuf:"varint,1,req,name=type,enum=equtils.I2GMsgReq_Event_Type" json:"type,omitempty"`
	FuncCall         *I2GMsgReq_Event_FuncCall `protobuf:"bytes,2,opt" json:"FuncCall,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *I2GMsgReq_Event) Reset()         { *m = I2GMsgReq_Event{} }
func (m *I2GMsgReq_Event) String() string { return proto.CompactTextString(m) }
func (*I2GMsgReq_Event) ProtoMessage()    {}

func (m *I2GMsgReq_Event) GetType() I2GMsgReq_Event_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return I2GMsgReq_Event_FUNC_CALL
}

func (m *I2GMsgReq_Event) GetFuncCall() *I2GMsgReq_Event_FuncCall {
	if m != nil {
		return m.FuncCall
	}
	return nil
}

type I2GMsgReq_Initiation struct {
	NodeId           *string `protobuf:"bytes,1,req,name=nodeId" json:"nodeId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *I2GMsgReq_Initiation) Reset()         { *m = I2GMsgReq_Initiation{} }
func (m *I2GMsgReq_Initiation) String() string { return proto.CompactTextString(m) }
func (*I2GMsgReq_Initiation) ProtoMessage()    {}

func (m *I2GMsgReq_Initiation) GetNodeId() string {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return ""
}

type I2GMsgReq struct {
	NodeId           *string               `protobuf:"bytes,1,req,name=node_id" json:"node_id,omitempty"`
	Type             *I2GMsgReq_Type       `protobuf:"varint,2,req,name=type,enum=equtils.I2GMsgReq_Type" json:"type,omitempty"`
	Pid              *int32                `protobuf:"varint,3,req,name=pid" json:"pid,omitempty"`
	Tid              *int32                `protobuf:"varint,4,req,name=tid" json:"tid,omitempty"`
	MsgId            *int32                `protobuf:"varint,5,req,name=msg_id" json:"msg_id,omitempty"`
	GaMsgId          *int32                `protobuf:"varint,6,opt,name=ga_msg_id" json:"ga_msg_id,omitempty"`
	Event            *I2GMsgReq_Event      `protobuf:"bytes,7,opt" json:"Event,omitempty"`
	Initiation       *I2GMsgReq_Initiation `protobuf:"bytes,8,opt" json:"Initiation,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *I2GMsgReq) Reset()         { *m = I2GMsgReq{} }
func (m *I2GMsgReq) String() string { return proto.CompactTextString(m) }
func (*I2GMsgReq) ProtoMessage()    {}

func (m *I2GMsgReq) GetNodeId() string {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return ""
}

func (m *I2GMsgReq) GetType() I2GMsgReq_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return I2GMsgReq_EVENT
}

func (m *I2GMsgReq) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *I2GMsgReq) GetTid() int32 {
	if m != nil && m.Tid != nil {
		return *m.Tid
	}
	return 0
}

func (m *I2GMsgReq) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *I2GMsgReq) GetGaMsgId() int32 {
	if m != nil && m.GaMsgId != nil {
		return *m.GaMsgId
	}
	return 0
}

func (m *I2GMsgReq) GetEvent() *I2GMsgReq_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *I2GMsgReq) GetInitiation() *I2GMsgReq_Initiation {
	if m != nil {
		return m.Initiation
	}
	return nil
}

type I2GMsgRsp struct {
	Res              *I2GMsgRsp_Result `protobuf:"varint,1,req,name=res,enum=equtils.I2GMsgRsp_Result" json:"res,omitempty"`
	MsgId            *int32            `protobuf:"varint,2,opt,name=msg_id" json:"msg_id,omitempty"`
	GaMsgId          *int32            `protobuf:"varint,3,opt,name=ga_msg_id" json:"ga_msg_id,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *I2GMsgRsp) Reset()         { *m = I2GMsgRsp{} }
func (m *I2GMsgRsp) String() string { return proto.CompactTextString(m) }
func (*I2GMsgRsp) ProtoMessage()    {}

func (m *I2GMsgRsp) GetRes() I2GMsgRsp_Result {
	if m != nil && m.Res != nil {
		return *m.Res
	}
	return I2GMsgRsp_ACK
}

func (m *I2GMsgRsp) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *I2GMsgRsp) GetGaMsgId() int32 {
	if m != nil && m.GaMsgId != nil {
		return *m.GaMsgId
	}
	return 0
}

func init() {
	proto.RegisterEnum("equtils.I2GMsgReq_Event_Type", I2GMsgReq_Event_Type_name, I2GMsgReq_Event_Type_value)
	proto.RegisterEnum("equtils.I2GMsgReq_Type", I2GMsgReq_Type_name, I2GMsgReq_Type_value)
	proto.RegisterEnum("equtils.I2GMsgRsp_Result", I2GMsgRsp_Result_name, I2GMsgRsp_Result_value)
}
