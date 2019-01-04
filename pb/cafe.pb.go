// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cafe.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CafeChallenge struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeChallenge) Reset()         { *m = CafeChallenge{} }
func (m *CafeChallenge) String() string { return proto.CompactTextString(m) }
func (*CafeChallenge) ProtoMessage()    {}
func (*CafeChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{0}
}
func (m *CafeChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeChallenge.Unmarshal(m, b)
}
func (m *CafeChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeChallenge.Marshal(b, m, deterministic)
}
func (dst *CafeChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeChallenge.Merge(dst, src)
}
func (m *CafeChallenge) XXX_Size() int {
	return xxx_messageInfo_CafeChallenge.Size(m)
}
func (m *CafeChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_CafeChallenge proto.InternalMessageInfo

func (m *CafeChallenge) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CafeNonce struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeNonce) Reset()         { *m = CafeNonce{} }
func (m *CafeNonce) String() string { return proto.CompactTextString(m) }
func (*CafeNonce) ProtoMessage()    {}
func (*CafeNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{1}
}
func (m *CafeNonce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeNonce.Unmarshal(m, b)
}
func (m *CafeNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeNonce.Marshal(b, m, deterministic)
}
func (dst *CafeNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeNonce.Merge(dst, src)
}
func (m *CafeNonce) XXX_Size() int {
	return xxx_messageInfo_CafeNonce.Size(m)
}
func (m *CafeNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeNonce.DiscardUnknown(m)
}

var xxx_messageInfo_CafeNonce proto.InternalMessageInfo

func (m *CafeNonce) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CafeRegistration struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Sig                  []byte   `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRegistration) Reset()         { *m = CafeRegistration{} }
func (m *CafeRegistration) String() string { return proto.CompactTextString(m) }
func (*CafeRegistration) ProtoMessage()    {}
func (*CafeRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{2}
}
func (m *CafeRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRegistration.Unmarshal(m, b)
}
func (m *CafeRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRegistration.Marshal(b, m, deterministic)
}
func (dst *CafeRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRegistration.Merge(dst, src)
}
func (m *CafeRegistration) XXX_Size() int {
	return xxx_messageInfo_CafeRegistration.Size(m)
}
func (m *CafeRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRegistration proto.InternalMessageInfo

func (m *CafeRegistration) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CafeRegistration) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CafeRegistration) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *CafeRegistration) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type CafeSession struct {
	Cafe                 string               `protobuf:"bytes,1,opt,name=cafe,proto3" json:"cafe,omitempty"`
	Access               string               `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Exp                  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Refresh              string               `protobuf:"bytes,4,opt,name=refresh,proto3" json:"refresh,omitempty"`
	Rexp                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=rexp,proto3" json:"rexp,omitempty"`
	Subject              string               `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	Type                 string               `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	SwarmAddrs           []string             `protobuf:"bytes,8,rep,name=swarmAddrs,proto3" json:"swarmAddrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeSession) Reset()         { *m = CafeSession{} }
func (m *CafeSession) String() string { return proto.CompactTextString(m) }
func (*CafeSession) ProtoMessage()    {}
func (*CafeSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{3}
}
func (m *CafeSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeSession.Unmarshal(m, b)
}
func (m *CafeSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeSession.Marshal(b, m, deterministic)
}
func (dst *CafeSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeSession.Merge(dst, src)
}
func (m *CafeSession) XXX_Size() int {
	return xxx_messageInfo_CafeSession.Size(m)
}
func (m *CafeSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeSession proto.InternalMessageInfo

func (m *CafeSession) GetCafe() string {
	if m != nil {
		return m.Cafe
	}
	return ""
}

func (m *CafeSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeSession) GetExp() *timestamp.Timestamp {
	if m != nil {
		return m.Exp
	}
	return nil
}

func (m *CafeSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

func (m *CafeSession) GetRexp() *timestamp.Timestamp {
	if m != nil {
		return m.Rexp
	}
	return nil
}

func (m *CafeSession) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CafeSession) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CafeSession) GetSwarmAddrs() []string {
	if m != nil {
		return m.SwarmAddrs
	}
	return nil
}

type CafeRefreshSession struct {
	Access               string   `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              string   `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRefreshSession) Reset()         { *m = CafeRefreshSession{} }
func (m *CafeRefreshSession) String() string { return proto.CompactTextString(m) }
func (*CafeRefreshSession) ProtoMessage()    {}
func (*CafeRefreshSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{4}
}
func (m *CafeRefreshSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRefreshSession.Unmarshal(m, b)
}
func (m *CafeRefreshSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRefreshSession.Marshal(b, m, deterministic)
}
func (dst *CafeRefreshSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRefreshSession.Merge(dst, src)
}
func (m *CafeRefreshSession) XXX_Size() int {
	return xxx_messageInfo_CafeRefreshSession.Size(m)
}
func (m *CafeRefreshSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRefreshSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRefreshSession proto.InternalMessageInfo

func (m *CafeRefreshSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeRefreshSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

type CafeStore struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cids                 []string `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStore) Reset()         { *m = CafeStore{} }
func (m *CafeStore) String() string { return proto.CompactTextString(m) }
func (*CafeStore) ProtoMessage()    {}
func (*CafeStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{5}
}
func (m *CafeStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStore.Unmarshal(m, b)
}
func (m *CafeStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStore.Marshal(b, m, deterministic)
}
func (dst *CafeStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStore.Merge(dst, src)
}
func (m *CafeStore) XXX_Size() int {
	return xxx_messageInfo_CafeStore.Size(m)
}
func (m *CafeStore) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStore.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStore proto.InternalMessageInfo

func (m *CafeStore) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStore) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeObjectList struct {
	Cids                 []string `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeObjectList) Reset()         { *m = CafeObjectList{} }
func (m *CafeObjectList) String() string { return proto.CompactTextString(m) }
func (*CafeObjectList) ProtoMessage()    {}
func (*CafeObjectList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{6}
}
func (m *CafeObjectList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeObjectList.Unmarshal(m, b)
}
func (m *CafeObjectList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeObjectList.Marshal(b, m, deterministic)
}
func (dst *CafeObjectList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeObjectList.Merge(dst, src)
}
func (m *CafeObjectList) XXX_Size() int {
	return xxx_messageInfo_CafeObjectList.Size(m)
}
func (m *CafeObjectList) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeObjectList.DiscardUnknown(m)
}

var xxx_messageInfo_CafeObjectList proto.InternalMessageInfo

func (m *CafeObjectList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeObject struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cid                  string   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Node                 []byte   `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeObject) Reset()         { *m = CafeObject{} }
func (m *CafeObject) String() string { return proto.CompactTextString(m) }
func (*CafeObject) ProtoMessage()    {}
func (*CafeObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{7}
}
func (m *CafeObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeObject.Unmarshal(m, b)
}
func (m *CafeObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeObject.Marshal(b, m, deterministic)
}
func (dst *CafeObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeObject.Merge(dst, src)
}
func (m *CafeObject) XXX_Size() int {
	return xxx_messageInfo_CafeObject.Size(m)
}
func (m *CafeObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeObject.DiscardUnknown(m)
}

var xxx_messageInfo_CafeObject proto.InternalMessageInfo

func (m *CafeObject) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeObject) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *CafeObject) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CafeObject) GetNode() []byte {
	if m != nil {
		return m.Node
	}
	return nil
}

type CafeStoreThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStoreThread) Reset()         { *m = CafeStoreThread{} }
func (m *CafeStoreThread) String() string { return proto.CompactTextString(m) }
func (*CafeStoreThread) ProtoMessage()    {}
func (*CafeStoreThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{8}
}
func (m *CafeStoreThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStoreThread.Unmarshal(m, b)
}
func (m *CafeStoreThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStoreThread.Marshal(b, m, deterministic)
}
func (dst *CafeStoreThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStoreThread.Merge(dst, src)
}
func (m *CafeStoreThread) XXX_Size() int {
	return xxx_messageInfo_CafeStoreThread.Size(m)
}
func (m *CafeStoreThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStoreThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStoreThread proto.InternalMessageInfo

func (m *CafeStoreThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStoreThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeStoreThread) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type CafeThread struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Sk                   []byte   `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Schema               string   `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Initiator            string   `protobuf:"bytes,5,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Type                 int32    `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	State                int32    `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	Head                 string   `protobuf:"bytes,8,opt,name=head,proto3" json:"head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeThread) Reset()         { *m = CafeThread{} }
func (m *CafeThread) String() string { return proto.CompactTextString(m) }
func (*CafeThread) ProtoMessage()    {}
func (*CafeThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{9}
}
func (m *CafeThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeThread.Unmarshal(m, b)
}
func (m *CafeThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeThread.Marshal(b, m, deterministic)
}
func (dst *CafeThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeThread.Merge(dst, src)
}
func (m *CafeThread) XXX_Size() int {
	return xxx_messageInfo_CafeThread.Size(m)
}
func (m *CafeThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeThread proto.InternalMessageInfo

func (m *CafeThread) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CafeThread) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *CafeThread) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CafeThread) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *CafeThread) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *CafeThread) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CafeThread) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *CafeThread) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

type CafeStored struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStored) Reset()         { *m = CafeStored{} }
func (m *CafeStored) String() string { return proto.CompactTextString(m) }
func (*CafeStored) ProtoMessage()    {}
func (*CafeStored) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{10}
}
func (m *CafeStored) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStored.Unmarshal(m, b)
}
func (m *CafeStored) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStored.Marshal(b, m, deterministic)
}
func (dst *CafeStored) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStored.Merge(dst, src)
}
func (m *CafeStored) XXX_Size() int {
	return xxx_messageInfo_CafeStored.Size(m)
}
func (m *CafeStored) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStored.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStored proto.InternalMessageInfo

func (m *CafeStored) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeDeliverMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeliverMessage) Reset()         { *m = CafeDeliverMessage{} }
func (m *CafeDeliverMessage) String() string { return proto.CompactTextString(m) }
func (*CafeDeliverMessage) ProtoMessage()    {}
func (*CafeDeliverMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{11}
}
func (m *CafeDeliverMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeliverMessage.Unmarshal(m, b)
}
func (m *CafeDeliverMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeliverMessage.Marshal(b, m, deterministic)
}
func (dst *CafeDeliverMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeliverMessage.Merge(dst, src)
}
func (m *CafeDeliverMessage) XXX_Size() int {
	return xxx_messageInfo_CafeDeliverMessage.Size(m)
}
func (m *CafeDeliverMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeliverMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeliverMessage proto.InternalMessageInfo

func (m *CafeDeliverMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeDeliverMessage) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type CafeCheckMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeCheckMessages) Reset()         { *m = CafeCheckMessages{} }
func (m *CafeCheckMessages) String() string { return proto.CompactTextString(m) }
func (*CafeCheckMessages) ProtoMessage()    {}
func (*CafeCheckMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{12}
}
func (m *CafeCheckMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeCheckMessages.Unmarshal(m, b)
}
func (m *CafeCheckMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeCheckMessages.Marshal(b, m, deterministic)
}
func (dst *CafeCheckMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeCheckMessages.Merge(dst, src)
}
func (m *CafeCheckMessages) XXX_Size() int {
	return xxx_messageInfo_CafeCheckMessages.Size(m)
}
func (m *CafeCheckMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeCheckMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeCheckMessages proto.InternalMessageInfo

func (m *CafeCheckMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeMessage struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId               string               `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CafeMessage) Reset()         { *m = CafeMessage{} }
func (m *CafeMessage) String() string { return proto.CompactTextString(m) }
func (*CafeMessage) ProtoMessage()    {}
func (*CafeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{13}
}
func (m *CafeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessage.Unmarshal(m, b)
}
func (m *CafeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessage.Marshal(b, m, deterministic)
}
func (dst *CafeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessage.Merge(dst, src)
}
func (m *CafeMessage) XXX_Size() int {
	return xxx_messageInfo_CafeMessage.Size(m)
}
func (m *CafeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessage proto.InternalMessageInfo

func (m *CafeMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeMessage) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *CafeMessage) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type CafeMessages struct {
	Messages             []*CafeMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CafeMessages) Reset()         { *m = CafeMessages{} }
func (m *CafeMessages) String() string { return proto.CompactTextString(m) }
func (*CafeMessages) ProtoMessage()    {}
func (*CafeMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{14}
}
func (m *CafeMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessages.Unmarshal(m, b)
}
func (m *CafeMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessages.Marshal(b, m, deterministic)
}
func (dst *CafeMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessages.Merge(dst, src)
}
func (m *CafeMessages) XXX_Size() int {
	return xxx_messageInfo_CafeMessages.Size(m)
}
func (m *CafeMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessages proto.InternalMessageInfo

func (m *CafeMessages) GetMessages() []*CafeMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CafeDeleteMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessages) Reset()         { *m = CafeDeleteMessages{} }
func (m *CafeDeleteMessages) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessages) ProtoMessage()    {}
func (*CafeDeleteMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{15}
}
func (m *CafeDeleteMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessages.Unmarshal(m, b)
}
func (m *CafeDeleteMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessages.Marshal(b, m, deterministic)
}
func (dst *CafeDeleteMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessages.Merge(dst, src)
}
func (m *CafeDeleteMessages) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessages.Size(m)
}
func (m *CafeDeleteMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessages proto.InternalMessageInfo

func (m *CafeDeleteMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeDeleteMessagesAck struct {
	More                 bool     `protobuf:"varint,1,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessagesAck) Reset()         { *m = CafeDeleteMessagesAck{} }
func (m *CafeDeleteMessagesAck) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessagesAck) ProtoMessage()    {}
func (*CafeDeleteMessagesAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_5da1b2563c594fd4, []int{16}
}
func (m *CafeDeleteMessagesAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessagesAck.Unmarshal(m, b)
}
func (m *CafeDeleteMessagesAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessagesAck.Marshal(b, m, deterministic)
}
func (dst *CafeDeleteMessagesAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessagesAck.Merge(dst, src)
}
func (m *CafeDeleteMessagesAck) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessagesAck.Size(m)
}
func (m *CafeDeleteMessagesAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessagesAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessagesAck proto.InternalMessageInfo

func (m *CafeDeleteMessagesAck) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

func init() {
	proto.RegisterType((*CafeChallenge)(nil), "CafeChallenge")
	proto.RegisterType((*CafeNonce)(nil), "CafeNonce")
	proto.RegisterType((*CafeRegistration)(nil), "CafeRegistration")
	proto.RegisterType((*CafeSession)(nil), "CafeSession")
	proto.RegisterType((*CafeRefreshSession)(nil), "CafeRefreshSession")
	proto.RegisterType((*CafeStore)(nil), "CafeStore")
	proto.RegisterType((*CafeObjectList)(nil), "CafeObjectList")
	proto.RegisterType((*CafeObject)(nil), "CafeObject")
	proto.RegisterType((*CafeStoreThread)(nil), "CafeStoreThread")
	proto.RegisterType((*CafeThread)(nil), "CafeThread")
	proto.RegisterType((*CafeStored)(nil), "CafeStored")
	proto.RegisterType((*CafeDeliverMessage)(nil), "CafeDeliverMessage")
	proto.RegisterType((*CafeCheckMessages)(nil), "CafeCheckMessages")
	proto.RegisterType((*CafeMessage)(nil), "CafeMessage")
	proto.RegisterType((*CafeMessages)(nil), "CafeMessages")
	proto.RegisterType((*CafeDeleteMessages)(nil), "CafeDeleteMessages")
	proto.RegisterType((*CafeDeleteMessagesAck)(nil), "CafeDeleteMessagesAck")
}

func init() { proto.RegisterFile("cafe.proto", fileDescriptor_cafe_5da1b2563c594fd4) }

var fileDescriptor_cafe_5da1b2563c594fd4 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6b, 0x13, 0x4f,
	0x14, 0x65, 0xb3, 0x49, 0x9a, 0xdc, 0xe6, 0xd7, 0x5f, 0x5c, 0x6a, 0x59, 0x4a, 0xd1, 0x3a, 0xf8,
	0x90, 0xaa, 0xa4, 0x50, 0x11, 0x7c, 0xb4, 0x56, 0x04, 0xc1, 0x3f, 0xb0, 0x2d, 0x08, 0xe2, 0xcb,
	0x64, 0xe7, 0x26, 0x99, 0x66, 0xff, 0x31, 0x33, 0xad, 0xed, 0x17, 0xf3, 0xeb, 0x29, 0x77, 0x66,
	0x76, 0xb3, 0x55, 0x43, 0xdf, 0xee, 0x99, 0x9c, 0x9c, 0x73, 0xef, 0x99, 0x3b, 0x0b, 0x90, 0xf2,
	0x39, 0x4e, 0x2b, 0x55, 0x9a, 0x72, 0xff, 0xf1, 0xa2, 0x2c, 0x17, 0x19, 0x1e, 0x5b, 0x34, 0xbb,
	0x9a, 0x1f, 0x1b, 0x99, 0xa3, 0x36, 0x3c, 0xaf, 0x1c, 0x81, 0x1d, 0xc1, 0x7f, 0x67, 0x7c, 0x8e,
	0x67, 0x4b, 0x9e, 0x65, 0x58, 0x2c, 0x30, 0x8a, 0x61, 0x8b, 0x0b, 0xa1, 0x50, 0xeb, 0x38, 0x38,
	0x0c, 0x26, 0xc3, 0xa4, 0x86, 0xec, 0x09, 0x0c, 0x89, 0xfa, 0xb9, 0x2c, 0x52, 0x8c, 0x76, 0xa1,
	0x77, 0xcd, 0xb3, 0x2b, 0xf4, 0x24, 0x07, 0xd8, 0x25, 0x8c, 0x89, 0x92, 0xe0, 0x42, 0x6a, 0xa3,
	0xb8, 0x91, 0x65, 0xb1, 0x59, 0x70, 0xad, 0xd1, 0x69, 0x69, 0xd0, 0x69, 0x41, 0x16, 0x71, 0xe8,
	0x4e, 0x2d, 0x88, 0xc6, 0x10, 0x6a, 0xb9, 0x88, 0xbb, 0x87, 0xc1, 0x64, 0x94, 0x50, 0xc9, 0x7e,
	0x05, 0xb0, 0x4d, 0x66, 0xe7, 0xa8, 0x35, 0xf9, 0x44, 0xd0, 0xa5, 0xc1, 0xbd, 0x89, 0xad, 0xa3,
	0x3d, 0xe8, 0xf3, 0x34, 0x25, 0x6b, 0x67, 0xe1, 0x51, 0xf4, 0x02, 0x42, 0xbc, 0xa9, 0xac, 0xc3,
	0xf6, 0xc9, 0xfe, 0xd4, 0x85, 0x34, 0xad, 0x43, 0x9a, 0x5e, 0xd4, 0x21, 0x25, 0x44, 0xa3, 0x09,
	0x14, 0xce, 0x15, 0xea, 0xa5, 0xf5, 0x1f, 0x26, 0x35, 0x8c, 0xa6, 0xd0, 0x55, 0x24, 0xd4, 0xbb,
	0x57, 0xc8, 0xf2, 0x48, 0x49, 0x5f, 0xcd, 0x2e, 0x31, 0x35, 0x71, 0xdf, 0x29, 0x79, 0x48, 0xdd,
	0x9b, 0xdb, 0x0a, 0xe3, 0x2d, 0xd7, 0x3d, 0xd5, 0xd1, 0x23, 0x00, 0xfd, 0x83, 0xab, 0xfc, 0x54,
	0x08, 0xa5, 0xe3, 0xc1, 0x61, 0x38, 0x19, 0x26, 0xad, 0x13, 0xf6, 0x1e, 0x22, 0x97, 0xb6, 0x6d,
	0xa6, 0xce, 0x61, 0x3d, 0x73, 0x70, 0x67, 0xe6, 0xd6, 0x14, 0x9d, 0x3b, 0x53, 0xb0, 0x57, 0xee,
	0x62, 0xcf, 0x4d, 0xa9, 0x6c, 0xfc, 0xa6, 0x5c, 0x61, 0x51, 0x5f, 0xac, 0x05, 0x36, 0x5c, 0x29,
	0x28, 0xc6, 0xd0, 0x86, 0x2b, 0x85, 0x66, 0x4f, 0x61, 0x87, 0xfe, 0xf6, 0xc5, 0x0e, 0xf0, 0x51,
	0x6a, 0xd3, 0xb0, 0x82, 0x16, 0xeb, 0x3b, 0xc0, 0x9a, 0xb5, 0x41, 0x7d, 0x0c, 0x61, 0x2a, 0x85,
	0x6f, 0x8b, 0x4a, 0x52, 0x12, 0xdc, 0x70, 0x7b, 0x43, 0xa3, 0xc4, 0xd6, 0x74, 0x56, 0x94, 0x02,
	0xfd, 0x0e, 0xd8, 0x9a, 0x7d, 0x85, 0xff, 0x9b, 0xd6, 0x2f, 0x96, 0x0a, 0xb9, 0xd8, 0x60, 0xb1,
	0x03, 0x9d, 0xc6, 0xa1, 0x23, 0x05, 0x65, 0x9b, 0xca, 0x6a, 0x89, 0xca, 0xe0, 0x8d, 0xf1, 0x36,
	0xad, 0x13, 0xf6, 0x33, 0x70, 0x7d, 0x7b, 0xd1, 0x31, 0x84, 0x2b, 0xbc, 0xf5, 0x92, 0x54, 0x92,
	0xa0, 0x5e, 0x59, 0xc1, 0x51, 0xd2, 0xd1, 0x2b, 0xdb, 0x1d, 0xcf, 0xeb, 0xad, 0xb5, 0x35, 0x5d,
	0x85, 0x4e, 0x97, 0x98, 0x73, 0xbf, 0x37, 0x1e, 0x45, 0x07, 0x30, 0x94, 0x85, 0x34, 0x92, 0x9b,
	0x52, 0xd9, 0xdd, 0x19, 0x26, 0xeb, 0x83, 0x66, 0x15, 0x68, 0x43, 0x7a, 0x7e, 0x15, 0x76, 0xa1,
	0xa7, 0x0d, 0x37, 0x6e, 0x3f, 0x7a, 0x89, 0x03, 0xc4, 0x5c, 0x22, 0x17, 0xf1, 0xc0, 0x79, 0x52,
	0xcd, 0x0e, 0x5c, 0xdf, 0x36, 0x11, 0xe1, 0xc7, 0x0e, 0xea, 0xb1, 0xd9, 0x1b, 0xb7, 0x32, 0xef,
	0x30, 0x93, 0xd7, 0xa8, 0x3e, 0xa1, 0xd6, 0x7c, 0x81, 0x7f, 0xb2, 0xa2, 0x7d, 0x18, 0xa4, 0x99,
	0xc4, 0xc2, 0x7c, 0xa8, 0x23, 0x6b, 0x30, 0x3b, 0x82, 0x07, 0xee, 0x83, 0x81, 0xe9, 0xca, 0xff,
	0x5f, 0xff, 0x3b, 0x73, 0x86, 0xee, 0x81, 0x6e, 0x72, 0xd9, 0x83, 0x7e, 0x85, 0xa8, 0x1a, 0x0f,
	0x8f, 0xe8, 0x51, 0x09, 0x1a, 0xf5, 0xfe, 0xd7, 0x69, 0x79, 0xec, 0x35, 0x8c, 0x5a, 0x36, 0x3a,
	0x9a, 0xc0, 0x20, 0xf7, 0xb5, 0xdd, 0xc4, 0xed, 0x93, 0xd1, 0xb4, 0x45, 0x48, 0x9a, 0x5f, 0xd9,
	0xb3, 0x26, 0x0d, 0x34, 0x78, 0xcf, 0x30, 0xcf, 0xe1, 0xe1, 0xdf, 0xdc, 0xd3, 0xd4, 0x5e, 0x7c,
	0x5e, 0x2a, 0xf7, 0xdd, 0x19, 0x24, 0xb6, 0x7e, 0xdb, 0xfd, 0xd6, 0xa9, 0x66, 0xb3, 0xbe, 0x6d,
	0xf9, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xdb, 0xdf, 0x73, 0x91, 0x05, 0x00, 0x00,
}
