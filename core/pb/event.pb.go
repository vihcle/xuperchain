// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// BlockStatus block status
type BlockStatus int32

const (
	BlockStatus_ERROR    BlockStatus = 0
	BlockStatus_TRUNK    BlockStatus = 1
	BlockStatus_BRANCH   BlockStatus = 2
	BlockStatus_NONEXIST BlockStatus = 3
)

var BlockStatus_name = map[int32]string{
	0: "ERROR",
	1: "TRUNK",
	2: "BRANCH",
	3: "NONEXIST",
}

var BlockStatus_value = map[string]int32{
	"ERROR":    0,
	"TRUNK":    1,
	"BRANCH":   2,
	"NONEXIST": 3,
}

func (x BlockStatus) String() string {
	return proto.EnumName(BlockStatus_name, int32(x))
}

func (BlockStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

///// EventType 事件类型
type EventType int32

const (
	EventType_UNDEFINED          EventType = 0
	EventType_BLOCK              EventType = 1
	EventType_TRANSACTION        EventType = 2
	EventType_ACCOUNT            EventType = 3
	EventType_SUBSCRIBE_RESPONSE EventType = 4
)

var EventType_name = map[int32]string{
	0: "UNDEFINED",
	1: "BLOCK",
	2: "TRANSACTION",
	3: "ACCOUNT",
	4: "SUBSCRIBE_RESPONSE",
}

var EventType_value = map[string]int32{
	"UNDEFINED":          0,
	"BLOCK":              1,
	"TRANSACTION":        2,
	"ACCOUNT":            3,
	"SUBSCRIBE_RESPONSE": 4,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

type UnsubscribeStatusInfo int32

const (
	UnsubscribeStatusInfo_UNSUBSCRIBE_UNDEFINED UnsubscribeStatusInfo = 0
	UnsubscribeStatusInfo_UNSUBSCRIBE_SUCCESS   UnsubscribeStatusInfo = 1
	UnsubscribeStatusInfo_UNSUBSCRIBE_FAILED    UnsubscribeStatusInfo = 2
)

var UnsubscribeStatusInfo_name = map[int32]string{
	0: "UNSUBSCRIBE_UNDEFINED",
	1: "UNSUBSCRIBE_SUCCESS",
	2: "UNSUBSCRIBE_FAILED",
}

var UnsubscribeStatusInfo_value = map[string]int32{
	"UNSUBSCRIBE_UNDEFINED": 0,
	"UNSUBSCRIBE_SUCCESS":   1,
	"UNSUBSCRIBE_FAILED":    2,
}

func (x UnsubscribeStatusInfo) String() string {
	return proto.EnumName(UnsubscribeStatusInfo_name, int32(x))
}

func (UnsubscribeStatusInfo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

// BlockStatusInfo 区块元数据
type BlockStatusInfo struct {
	Bcname               string      `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Proposer             string      `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Height               int64       `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Status               BlockStatus `protobuf:"varint,4,opt,name=status,proto3,enum=pb.BlockStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlockStatusInfo) Reset()         { *m = BlockStatusInfo{} }
func (m *BlockStatusInfo) String() string { return proto.CompactTextString(m) }
func (*BlockStatusInfo) ProtoMessage()    {}
func (*BlockStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *BlockStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockStatusInfo.Unmarshal(m, b)
}
func (m *BlockStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockStatusInfo.Marshal(b, m, deterministic)
}
func (m *BlockStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStatusInfo.Merge(m, src)
}
func (m *BlockStatusInfo) XXX_Size() int {
	return xxx_messageInfo_BlockStatusInfo.Size(m)
}
func (m *BlockStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStatusInfo proto.InternalMessageInfo

func (m *BlockStatusInfo) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *BlockStatusInfo) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *BlockStatusInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockStatusInfo) GetStatus() BlockStatus {
	if m != nil {
		return m.Status
	}
	return BlockStatus_ERROR
}

type TransactionStatusInfo struct {
	Bcname               string            `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Initiator            string            `protobuf:"bytes,2,opt,name=initiator,proto3" json:"initiator,omitempty"`
	AuthRequire          []string          `protobuf:"bytes,3,rep,name=auth_require,json=authRequire,proto3" json:"auth_require,omitempty"`
	Status               TransactionStatus `protobuf:"varint,4,opt,name=status,proto3,enum=pb.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionStatusInfo) Reset()         { *m = TransactionStatusInfo{} }
func (m *TransactionStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionStatusInfo) ProtoMessage()    {}
func (*TransactionStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *TransactionStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStatusInfo.Unmarshal(m, b)
}
func (m *TransactionStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStatusInfo.Marshal(b, m, deterministic)
}
func (m *TransactionStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStatusInfo.Merge(m, src)
}
func (m *TransactionStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionStatusInfo.Size(m)
}
func (m *TransactionStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStatusInfo proto.InternalMessageInfo

func (m *TransactionStatusInfo) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *TransactionStatusInfo) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *TransactionStatusInfo) GetAuthRequire() []string {
	if m != nil {
		return m.AuthRequire
	}
	return nil
}

func (m *TransactionStatusInfo) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_UNDEFINE
}

type AccountStatusInfo struct {
	Bcname               string            `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	FromAddr             []string          `protobuf:"bytes,2,rep,name=from_addr,json=fromAddr,proto3" json:"from_addr,omitempty"`
	ToAddr               []string          `protobuf:"bytes,3,rep,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	Status               TransactionStatus `protobuf:"varint,4,opt,name=status,proto3,enum=pb.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AccountStatusInfo) Reset()         { *m = AccountStatusInfo{} }
func (m *AccountStatusInfo) String() string { return proto.CompactTextString(m) }
func (*AccountStatusInfo) ProtoMessage()    {}
func (*AccountStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

func (m *AccountStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountStatusInfo.Unmarshal(m, b)
}
func (m *AccountStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountStatusInfo.Marshal(b, m, deterministic)
}
func (m *AccountStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountStatusInfo.Merge(m, src)
}
func (m *AccountStatusInfo) XXX_Size() int {
	return xxx_messageInfo_AccountStatusInfo.Size(m)
}
func (m *AccountStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountStatusInfo proto.InternalMessageInfo

func (m *AccountStatusInfo) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *AccountStatusInfo) GetFromAddr() []string {
	if m != nil {
		return m.FromAddr
	}
	return nil
}

func (m *AccountStatusInfo) GetToAddr() []string {
	if m != nil {
		return m.ToAddr
	}
	return nil
}

func (m *AccountStatusInfo) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_UNDEFINE
}

/////// 事件订阅请求数据结构
// EventRequest 将事件订阅统一归为EventRequest
type EventRequest struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.EventType" json:"type,omitempty"`
	Payload              []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNDEFINED
}

func (m *EventRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// BlockEventRequest 订阅区块请求
type BlockEventRequest struct {
	Bcname               string   `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Proposer             string   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	StartHeight          int64    `protobuf:"varint,3,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight            int64    `protobuf:"varint,4,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	NeedContent          bool     `protobuf:"varint,5,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockEventRequest) Reset()         { *m = BlockEventRequest{} }
func (m *BlockEventRequest) String() string { return proto.CompactTextString(m) }
func (*BlockEventRequest) ProtoMessage()    {}
func (*BlockEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{4}
}

func (m *BlockEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEventRequest.Unmarshal(m, b)
}
func (m *BlockEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEventRequest.Marshal(b, m, deterministic)
}
func (m *BlockEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEventRequest.Merge(m, src)
}
func (m *BlockEventRequest) XXX_Size() int {
	return xxx_messageInfo_BlockEventRequest.Size(m)
}
func (m *BlockEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEventRequest proto.InternalMessageInfo

func (m *BlockEventRequest) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *BlockEventRequest) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *BlockEventRequest) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *BlockEventRequest) GetEndHeight() int64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

func (m *BlockEventRequest) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

// TransactionEventRequest 订阅交易请求
type TransactionEventRequest struct {
	Bcname               string   `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Initiator            string   `protobuf:"bytes,2,opt,name=initiator,proto3" json:"initiator,omitempty"`
	AuthRequire          string   `protobuf:"bytes,3,opt,name=auth_require,json=authRequire,proto3" json:"auth_require,omitempty"`
	NeedContent          bool     `protobuf:"varint,4,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionEventRequest) Reset()         { *m = TransactionEventRequest{} }
func (m *TransactionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionEventRequest) ProtoMessage()    {}
func (*TransactionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{5}
}

func (m *TransactionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEventRequest.Unmarshal(m, b)
}
func (m *TransactionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEventRequest.Marshal(b, m, deterministic)
}
func (m *TransactionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEventRequest.Merge(m, src)
}
func (m *TransactionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionEventRequest.Size(m)
}
func (m *TransactionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEventRequest proto.InternalMessageInfo

func (m *TransactionEventRequest) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *TransactionEventRequest) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *TransactionEventRequest) GetAuthRequire() string {
	if m != nil {
		return m.AuthRequire
	}
	return ""
}

func (m *TransactionEventRequest) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

// AccountEventRequest 订阅账户请求
type AccountEventRequest struct {
	Bcname               string   `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	FromAddr             string   `protobuf:"bytes,2,opt,name=from_addr,json=fromAddr,proto3" json:"from_addr,omitempty"`
	ToAddr               string   `protobuf:"bytes,3,opt,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	NeedContent          bool     `protobuf:"varint,4,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEventRequest) Reset()         { *m = AccountEventRequest{} }
func (m *AccountEventRequest) String() string { return proto.CompactTextString(m) }
func (*AccountEventRequest) ProtoMessage()    {}
func (*AccountEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{6}
}

func (m *AccountEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEventRequest.Unmarshal(m, b)
}
func (m *AccountEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEventRequest.Marshal(b, m, deterministic)
}
func (m *AccountEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEventRequest.Merge(m, src)
}
func (m *AccountEventRequest) XXX_Size() int {
	return xxx_messageInfo_AccountEventRequest.Size(m)
}
func (m *AccountEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEventRequest proto.InternalMessageInfo

func (m *AccountEventRequest) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *AccountEventRequest) GetFromAddr() string {
	if m != nil {
		return m.FromAddr
	}
	return ""
}

func (m *AccountEventRequest) GetToAddr() string {
	if m != nil {
		return m.ToAddr
	}
	return ""
}

func (m *AccountEventRequest) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

// UnsubscribeRequest 取消事件订阅请求
type UnsubscribeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{7}
}

func (m *UnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeRequest.Unmarshal(m, b)
}
func (m *UnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeRequest.Marshal(b, m, deterministic)
}
func (m *UnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeRequest.Merge(m, src)
}
func (m *UnsubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeRequest.Size(m)
}
func (m *UnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeRequest proto.InternalMessageInfo

func (m *UnsubscribeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//////// 事件订阅返回数据结构
/////// UnsubscribeResponse 取消订阅pb接口
type UnsubscribeResponse struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               UnsubscribeStatusInfo `protobuf:"varint,3,opt,name=status,proto3,enum=pb.UnsubscribeStatusInfo" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UnsubscribeResponse) Reset()         { *m = UnsubscribeResponse{} }
func (m *UnsubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeResponse) ProtoMessage()    {}
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{8}
}

func (m *UnsubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeResponse.Unmarshal(m, b)
}
func (m *UnsubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeResponse.Marshal(b, m, deterministic)
}
func (m *UnsubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeResponse.Merge(m, src)
}
func (m *UnsubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeResponse.Size(m)
}
func (m *UnsubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeResponse proto.InternalMessageInfo

func (m *UnsubscribeResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UnsubscribeResponse) GetStatus() UnsubscribeStatusInfo {
	if m != nil {
		return m.Status
	}
	return UnsubscribeStatusInfo_UNSUBSCRIBE_UNDEFINED
}

///// Event 将多种事件订阅回复统一为Event, 易于扩展
type Event struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 EventType              `protobuf:"varint,2,opt,name=type,proto3,enum=pb.EventType" json:"type,omitempty"`
	Payload              []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	BlockStatus          *BlockStatusInfo       `protobuf:"bytes,4,opt,name=block_status,json=blockStatus,proto3" json:"block_status,omitempty"`
	TxStatus             *TransactionStatusInfo `protobuf:"bytes,5,opt,name=tx_status,json=txStatus,proto3" json:"tx_status,omitempty"`
	AccountStatus        *AccountStatusInfo     `protobuf:"bytes,6,opt,name=account_status,json=accountStatus,proto3" json:"account_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{9}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNDEFINED
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetBlockStatus() *BlockStatusInfo {
	if m != nil {
		return m.BlockStatus
	}
	return nil
}

func (m *Event) GetTxStatus() *TransactionStatusInfo {
	if m != nil {
		return m.TxStatus
	}
	return nil
}

func (m *Event) GetAccountStatus() *AccountStatusInfo {
	if m != nil {
		return m.AccountStatus
	}
	return nil
}

// BlockEvent 订阅区块返回
type BlockEvent struct {
	Block                *InternalBlock `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BlockEvent) Reset()         { *m = BlockEvent{} }
func (m *BlockEvent) String() string { return proto.CompactTextString(m) }
func (*BlockEvent) ProtoMessage()    {}
func (*BlockEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{10}
}

func (m *BlockEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEvent.Unmarshal(m, b)
}
func (m *BlockEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEvent.Marshal(b, m, deterministic)
}
func (m *BlockEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEvent.Merge(m, src)
}
func (m *BlockEvent) XXX_Size() int {
	return xxx_messageInfo_BlockEvent.Size(m)
}
func (m *BlockEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEvent proto.InternalMessageInfo

func (m *BlockEvent) GetBlock() *InternalBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

// TransactionEvent 订阅交易返回
type TransactionEvent struct {
	Tx                   *Transaction `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransactionEvent) Reset()         { *m = TransactionEvent{} }
func (m *TransactionEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionEvent) ProtoMessage()    {}
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{11}
}

func (m *TransactionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEvent.Unmarshal(m, b)
}
func (m *TransactionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEvent.Marshal(b, m, deterministic)
}
func (m *TransactionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEvent.Merge(m, src)
}
func (m *TransactionEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionEvent.Size(m)
}
func (m *TransactionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEvent proto.InternalMessageInfo

func (m *TransactionEvent) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.BlockStatus", BlockStatus_name, BlockStatus_value)
	proto.RegisterEnum("pb.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("pb.UnsubscribeStatusInfo", UnsubscribeStatusInfo_name, UnsubscribeStatusInfo_value)
	proto.RegisterType((*BlockStatusInfo)(nil), "pb.BlockStatusInfo")
	proto.RegisterType((*TransactionStatusInfo)(nil), "pb.TransactionStatusInfo")
	proto.RegisterType((*AccountStatusInfo)(nil), "pb.AccountStatusInfo")
	proto.RegisterType((*EventRequest)(nil), "pb.EventRequest")
	proto.RegisterType((*BlockEventRequest)(nil), "pb.BlockEventRequest")
	proto.RegisterType((*TransactionEventRequest)(nil), "pb.TransactionEventRequest")
	proto.RegisterType((*AccountEventRequest)(nil), "pb.AccountEventRequest")
	proto.RegisterType((*UnsubscribeRequest)(nil), "pb.UnsubscribeRequest")
	proto.RegisterType((*UnsubscribeResponse)(nil), "pb.UnsubscribeResponse")
	proto.RegisterType((*Event)(nil), "pb.Event")
	proto.RegisterType((*BlockEvent)(nil), "pb.BlockEvent")
	proto.RegisterType((*TransactionEvent)(nil), "pb.TransactionEvent")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6f, 0xdb, 0x36,
	0x18, 0xc6, 0x4b, 0xc9, 0x76, 0xac, 0x57, 0x4a, 0xa2, 0x30, 0x48, 0xe2, 0x64, 0x1b, 0xe6, 0x08,
	0x03, 0x6a, 0x04, 0x58, 0xb0, 0xb9, 0x58, 0x4f, 0x1b, 0x30, 0x5b, 0x51, 0x51, 0xa3, 0x85, 0x5c,
	0x50, 0x36, 0x50, 0x60, 0x07, 0x43, 0x7f, 0xd8, 0x45, 0x58, 0x4a, 0x69, 0x12, 0x5d, 0x38, 0x87,
	0x5d, 0x76, 0xe8, 0x79, 0x97, 0x9d, 0xf7, 0x05, 0xf6, 0x21, 0x07, 0x52, 0x92, 0xad, 0xc8, 0x1d,
	0xe0, 0xf4, 0x16, 0x3e, 0x2f, 0x9f, 0x97, 0xbf, 0x90, 0xcf, 0x2b, 0x83, 0x4e, 0x3f, 0x50, 0xc6,
	0xaf, 0xd3, 0x2c, 0xe1, 0x09, 0x56, 0xd2, 0xe0, 0xc2, 0x58, 0x85, 0xb7, 0x7e, 0xcc, 0x0a, 0xc5,
	0xfa, 0x88, 0xe0, 0x70, 0x7c, 0x97, 0x84, 0xbf, 0x79, 0xdc, 0xe7, 0xcb, 0x7c, 0xc2, 0xde, 0x25,
	0xf8, 0x14, 0x3a, 0x41, 0xc8, 0xfc, 0xf7, 0xb4, 0x87, 0xfa, 0x68, 0xa0, 0x91, 0x72, 0x85, 0x2f,
	0xa0, 0x9b, 0x66, 0x49, 0x9a, 0xe4, 0x34, 0xeb, 0x29, 0xb2, 0xb2, 0x5e, 0x0b, 0xcf, 0x2d, 0x8d,
	0x7f, 0xbd, 0xe5, 0x3d, 0xb5, 0x8f, 0x06, 0x2a, 0x29, 0x57, 0xf8, 0x29, 0x74, 0x72, 0xd9, 0xb9,
	0xd7, 0xea, 0xa3, 0xc1, 0xc1, 0xf0, 0xf0, 0x3a, 0x0d, 0xae, 0x6b, 0x07, 0x92, 0xb2, 0x6c, 0xfd,
	0x83, 0xe0, 0x64, 0x96, 0xf9, 0x2c, 0xf7, 0x43, 0x1e, 0x27, 0x6c, 0x07, 0x9c, 0x2f, 0x41, 0x8b,
	0x59, 0xcc, 0x63, 0x9f, 0x27, 0x15, 0xcf, 0x46, 0xc0, 0x97, 0x60, 0xf8, 0x4b, 0x7e, 0xbb, 0xc8,
	0xe8, 0xef, 0xcb, 0x38, 0xa3, 0x3d, 0xb5, 0xaf, 0x0e, 0x34, 0xa2, 0x0b, 0x8d, 0x14, 0x12, 0xfe,
	0xb6, 0xc1, 0x76, 0x22, 0xd8, 0xb6, 0x18, 0xd6, 0x84, 0x7f, 0x21, 0x38, 0x1a, 0x85, 0x61, 0xb2,
	0x64, 0x7c, 0x07, 0xba, 0x2f, 0x40, 0x7b, 0x97, 0x25, 0xef, 0x17, 0x7e, 0x14, 0x09, 0x3a, 0x71,
	0x78, 0x57, 0x08, 0xa3, 0x28, 0xca, 0xf0, 0x19, 0xec, 0xf1, 0xa4, 0x28, 0x15, 0x5c, 0x1d, 0x9e,
	0xc8, 0xc2, 0x23, 0x91, 0x5e, 0x81, 0xe1, 0x88, 0xe7, 0x15, 0xff, 0x11, 0xcd, 0x39, 0xbe, 0x84,
	0x16, 0xbf, 0x4f, 0x0b, 0x94, 0x83, 0xe1, 0xbe, 0x30, 0xcb, 0xfa, 0xec, 0x3e, 0xa5, 0x44, 0x96,
	0x70, 0x0f, 0xf6, 0x52, 0xff, 0xfe, 0x2e, 0xf1, 0x23, 0x79, 0x67, 0x06, 0xa9, 0x96, 0xd6, 0xbf,
	0x08, 0x8e, 0xe4, 0xcb, 0x3c, 0x68, 0xf9, 0x39, 0x61, 0xb8, 0x04, 0x23, 0xe7, 0x7e, 0xc6, 0x17,
	0x0f, 0x22, 0xa1, 0x4b, 0xed, 0x65, 0x91, 0x8b, 0xaf, 0x00, 0x28, 0x8b, 0xaa, 0x0d, 0x2d, 0xb9,
	0x41, 0xa3, 0x2c, 0x2a, 0xcb, 0x97, 0x60, 0x30, 0x4a, 0xa3, 0x45, 0x98, 0x30, 0x4e, 0x19, 0xef,
	0xb5, 0xfb, 0x68, 0xd0, 0x25, 0xba, 0xd0, 0xec, 0x42, 0xb2, 0xfe, 0x46, 0x70, 0x56, 0xbb, 0x99,
	0x9d, 0xa0, 0x1f, 0x1b, 0x19, 0xd4, 0x8c, 0x4c, 0x93, 0xab, 0xb5, 0xcd, 0xf5, 0x11, 0xc1, 0x71,
	0x19, 0x93, 0x9d, 0x98, 0x1a, 0x41, 0x41, 0xff, 0x1f, 0x14, 0x54, 0x0b, 0xca, 0x0e, 0x20, 0xdf,
	0x00, 0x9e, 0xb3, 0x7c, 0x19, 0xe4, 0x61, 0x16, 0x07, 0xb4, 0xc2, 0x38, 0x00, 0x25, 0x8e, 0x4a,
	0x04, 0x25, 0x8e, 0xac, 0xb7, 0x70, 0xfc, 0x60, 0x57, 0x9e, 0x26, 0x2c, 0xa7, 0xcd, 0x6d, 0xf8,
	0xfb, 0x75, 0x30, 0x55, 0x99, 0xad, 0x73, 0x91, 0xad, 0x9a, 0x71, 0x33, 0x11, 0xeb, 0x70, 0xfe,
	0xa9, 0x40, 0x5b, 0xde, 0xc0, 0x56, 0xb3, 0x2a, 0xa6, 0xca, 0x4e, 0x31, 0x55, 0x1f, 0xc4, 0x14,
	0x3f, 0x07, 0x23, 0x10, 0x29, 0x5d, 0xd4, 0x06, 0x45, 0x1f, 0x1e, 0x37, 0xbe, 0x2b, 0x92, 0x44,
	0x0f, 0x36, 0x02, 0x7e, 0x0e, 0x1a, 0x5f, 0x55, 0xa6, 0xb6, 0x34, 0x9d, 0x7f, 0x72, 0xba, 0xa4,
	0xb5, 0xcb, 0x57, 0xa5, 0xef, 0x47, 0x38, 0xf0, 0x8b, 0xe7, 0xac, 0xcc, 0x1d, 0x69, 0x96, 0xa3,
	0xb9, 0xf5, 0x3d, 0x20, 0xfb, 0x7e, 0x5d, 0xb2, 0x7e, 0x00, 0xd8, 0xcc, 0x14, 0x7e, 0x0a, 0x6d,
	0x89, 0x24, 0xef, 0x42, 0x1f, 0x1e, 0x89, 0x16, 0x13, 0xc6, 0x69, 0xc6, 0xfc, 0x3b, 0xb9, 0x8d,
	0x14, 0x75, 0xeb, 0x19, 0x98, 0xcd, 0x6c, 0xe3, 0xaf, 0x41, 0xe1, 0xab, 0xd2, 0x79, 0xd8, 0x20,
	0x27, 0x0a, 0x5f, 0x5d, 0xfd, 0x04, 0x7a, 0xed, 0x06, 0xb0, 0x06, 0x6d, 0x87, 0x90, 0x29, 0x31,
	0x9f, 0x88, 0x3f, 0x67, 0x64, 0xee, 0xbe, 0x32, 0x11, 0x06, 0xe8, 0x8c, 0xc9, 0xc8, 0xb5, 0x5f,
	0x9a, 0x0a, 0x36, 0xa0, 0xeb, 0x4e, 0x5d, 0xe7, 0xed, 0xc4, 0x9b, 0x99, 0xea, 0xd5, 0x2f, 0xa0,
	0xad, 0x5f, 0x01, 0xef, 0x83, 0x36, 0x77, 0x6f, 0x9c, 0x17, 0x13, 0xd7, 0xb9, 0x29, 0x1a, 0x8c,
	0x5f, 0x4f, 0x6d, 0xd1, 0xe0, 0x10, 0xf4, 0x19, 0x19, 0xb9, 0xde, 0xc8, 0x9e, 0x4d, 0xa6, 0xae,
	0xa9, 0x60, 0x1d, 0xf6, 0x46, 0xb6, 0x3d, 0x9d, 0xbb, 0x33, 0x53, 0xc5, 0xa7, 0x80, 0xbd, 0xf9,
	0xd8, 0xb3, 0xc9, 0x64, 0xec, 0x2c, 0x88, 0xe3, 0xbd, 0x99, 0xba, 0x9e, 0x63, 0xb6, 0xae, 0x42,
	0x38, 0xf9, 0x64, 0x5a, 0xf0, 0x39, 0x9c, 0xcc, 0xdd, 0x8d, 0xa5, 0x7e, 0xe8, 0x19, 0x1c, 0xd7,
	0x4b, 0xde, 0xdc, 0xb6, 0x1d, 0xcf, 0x33, 0x91, 0x38, 0xa4, 0x5e, 0x78, 0x31, 0x9a, 0xbc, 0x76,
	0x6e, 0x4c, 0x65, 0xf8, 0x07, 0xec, 0xbf, 0x59, 0x06, 0xf9, 0x32, 0xf0, 0x68, 0xf6, 0x21, 0x0e,
	0x29, 0xbe, 0x02, 0xcd, 0xab, 0xce, 0xc4, 0xe6, 0x3a, 0x67, 0xe5, 0x2c, 0x5c, 0x68, 0x6b, 0xe5,
	0x3b, 0x84, 0x7f, 0x06, 0xbd, 0x46, 0x88, 0x4f, 0x1b, 0x01, 0xaf, 0x3c, 0x67, 0x5b, 0x7a, 0x31,
	0x31, 0xd6, 0x93, 0xa0, 0x23, 0x7f, 0x52, 0x9f, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x41, 0x0b,
	0x91, 0x99, 0x73, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubsubServiceClient is the client API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Subscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
}

type pubsubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubsubServiceClient(cc *grpc.ClientConn) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubsubService_serviceDesc.Streams[0], "/pb.PubsubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubsubServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/pb.PubsubService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubsubServiceServer is the server API for PubsubService service.
type PubsubServiceServer interface {
	Subscribe(*EventRequest, PubsubService_SubscribeServer) error
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
}

func RegisterPubsubServiceServer(s *grpc.Server, srv PubsubServiceServer) {
	s.RegisterService(&_PubsubService_serviceDesc, srv)
}

func _PubsubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subscribe(m, &pubsubServiceSubscribeServer{stream})
}

type PubsubService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _PubsubService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PubsubService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubsubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _PubsubService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}
