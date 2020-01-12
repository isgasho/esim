package new

func ProtoBufInit() {
	fc1 := &FileContent{
		FileName: "user.pb.go",
		Dir:      "internal/infra/third_party/protobuf/passport",
		Content: `// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passport/user.proto

package passport

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type GetUserByUserNameRequest struct {
	Username             string   {{!}}protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserByUserNameRequest) Reset()         { *m = GetUserByUserNameRequest{} }
func (m *GetUserByUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByUserNameRequest) ProtoMessage()    {}
func (*GetUserByUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{0}
}

func (m *GetUserByUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByUserNameRequest.Unmarshal(m, b)
}
func (m *GetUserByUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByUserNameRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByUserNameRequest.Merge(m, src)
}
func (m *GetUserByUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByUserNameRequest.Size(m)
}
func (m *GetUserByUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByUserNameRequest proto.InternalMessageInfo

func (m *GetUserByUserNameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserByIdRequest struct {
	Id                   int64    {{!}}protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserByIdRequest) Reset()         { *m = GetUserByIdRequest{} }
func (m *GetUserByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdRequest) ProtoMessage()    {}
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{1}
}

func (m *GetUserByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdRequest.Unmarshal(m, b)
}
func (m *GetUserByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdRequest.Merge(m, src)
}
func (m *GetUserByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdRequest.Size(m)
}
func (m *GetUserByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdRequest proto.InternalMessageInfo

func (m *GetUserByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserByIdsRequest struct {
	Ids                  string   {{!}}protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserByIdsRequest) Reset()         { *m = GetUserByIdsRequest{} }
func (m *GetUserByIdsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdsRequest) ProtoMessage()    {}
func (*GetUserByIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{2}
}

func (m *GetUserByIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdsRequest.Unmarshal(m, b)
}
func (m *GetUserByIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdsRequest.Merge(m, src)
}
func (m *GetUserByIdsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdsRequest.Size(m)
}
func (m *GetUserByIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdsRequest proto.InternalMessageInfo

func (m *GetUserByIdsRequest) GetIds() string {
	if m != nil {
		return m.Ids
	}
	return ""
}

type GetUserInfoRequest struct {
	Uid                  int64    {{!}}protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"{{!}}
	UserKeys             string   {{!}}protobuf:"bytes,2,opt,name=userKeys,proto3" json:"userKeys,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserInfoRequest) Reset()         { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()    {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{3}
}

func (m *GetUserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoRequest.Unmarshal(m, b)
}
func (m *GetUserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRequest.Merge(m, src)
}
func (m *GetUserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoRequest.Size(m)
}
func (m *GetUserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRequest proto.InternalMessageInfo

func (m *GetUserInfoRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GetUserInfoRequest) GetUserKeys() string {
	if m != nil {
		return m.UserKeys
	}
	return ""
}

type SetUserKeyRequest struct {
	Uid                  int64    {{!}}protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"{{!}}
	Key                  string   {{!}}protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"{{!}}
	Value                string   {{!}}protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *SetUserKeyRequest) Reset()         { *m = SetUserKeyRequest{} }
func (m *SetUserKeyRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserKeyRequest) ProtoMessage()    {}
func (*SetUserKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{4}
}

func (m *SetUserKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserKeyRequest.Unmarshal(m, b)
}
func (m *SetUserKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserKeyRequest.Marshal(b, m, deterministic)
}
func (m *SetUserKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserKeyRequest.Merge(m, src)
}
func (m *SetUserKeyRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserKeyRequest.Size(m)
}
func (m *SetUserKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserKeyRequest proto.InternalMessageInfo

func (m *SetUserKeyRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SetUserKeyRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetUserKeyRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetUserKeyArrRequest struct {
	Uid                  int64    {{!}}protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"{{!}}
	Info                 string   {{!}}protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *SetUserKeyArrRequest) Reset()         { *m = SetUserKeyArrRequest{} }
func (m *SetUserKeyArrRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserKeyArrRequest) ProtoMessage()    {}
func (*SetUserKeyArrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{5}
}

func (m *SetUserKeyArrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserKeyArrRequest.Unmarshal(m, b)
}
func (m *SetUserKeyArrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserKeyArrRequest.Marshal(b, m, deterministic)
}
func (m *SetUserKeyArrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserKeyArrRequest.Merge(m, src)
}
func (m *SetUserKeyArrRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserKeyArrRequest.Size(m)
}
func (m *SetUserKeyArrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserKeyArrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserKeyArrRequest proto.InternalMessageInfo

func (m *SetUserKeyArrRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SetUserKeyArrRequest) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type GetUserListRequest struct {
	Page                 int64    {{!}}protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"{{!}}
	PageSize             int64    {{!}}protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserListRequest) Reset()         { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()    {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{6}
}

func (m *GetUserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListRequest.Unmarshal(m, b)
}
func (m *GetUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListRequest.Marshal(b, m, deterministic)
}
func (m *GetUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListRequest.Merge(m, src)
}
func (m *GetUserListRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserListRequest.Size(m)
}
func (m *GetUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListRequest proto.InternalMessageInfo

func (m *GetUserListRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserListRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type GetUserIdsByUserNamesRequest struct {
	Names                string   {{!}}protobuf:"bytes,1,opt,name=names,proto3" json:"names,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserIdsByUserNamesRequest) Reset()         { *m = GetUserIdsByUserNamesRequest{} }
func (m *GetUserIdsByUserNamesRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserIdsByUserNamesRequest) ProtoMessage()    {}
func (*GetUserIdsByUserNamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{7}
}

func (m *GetUserIdsByUserNamesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserIdsByUserNamesRequest.Unmarshal(m, b)
}
func (m *GetUserIdsByUserNamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserIdsByUserNamesRequest.Marshal(b, m, deterministic)
}
func (m *GetUserIdsByUserNamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserIdsByUserNamesRequest.Merge(m, src)
}
func (m *GetUserIdsByUserNamesRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserIdsByUserNamesRequest.Size(m)
}
func (m *GetUserIdsByUserNamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserIdsByUserNamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserIdsByUserNamesRequest proto.InternalMessageInfo

func (m *GetUserIdsByUserNamesRequest) GetNames() string {
	if m != nil {
		return m.Names
	}
	return ""
}

// The response message containing the greetings
type GrpcReplyMap struct {
	Code                 int32             {{!}}protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"{{!}}
	Data                 map[string]string {{!}}protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"{{!}}
	Msg                  string            {{!}}protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{}          {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte            {{!}}json:"-"{{!}}
	XXX_sizecache        int32             {{!}}json:"-"{{!}}
}

func (m *GrpcReplyMap) Reset()         { *m = GrpcReplyMap{} }
func (m *GrpcReplyMap) String() string { return proto.CompactTextString(m) }
func (*GrpcReplyMap) ProtoMessage()    {}
func (*GrpcReplyMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{8}
}

func (m *GrpcReplyMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReplyMap.Unmarshal(m, b)
}
func (m *GrpcReplyMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReplyMap.Marshal(b, m, deterministic)
}
func (m *GrpcReplyMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReplyMap.Merge(m, src)
}
func (m *GrpcReplyMap) XXX_Size() int {
	return xxx_messageInfo_GrpcReplyMap.Size(m)
}
func (m *GrpcReplyMap) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReplyMap.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReplyMap proto.InternalMessageInfo

func (m *GrpcReplyMap) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcReplyMap) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GrpcReplyMap) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GrpcReplyMaps struct {
	Code                 int32      {{!}}protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"{{!}}
	Data                 []*DataMap {{!}}protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"{{!}}
	Msg                  string     {{!}}protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{}   {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte     {{!}}json:"-"{{!}}
	XXX_sizecache        int32      {{!}}json:"-"{{!}}
}

func (m *GrpcReplyMaps) Reset()         { *m = GrpcReplyMaps{} }
func (m *GrpcReplyMaps) String() string { return proto.CompactTextString(m) }
func (*GrpcReplyMaps) ProtoMessage()    {}
func (*GrpcReplyMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{9}
}

func (m *GrpcReplyMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReplyMaps.Unmarshal(m, b)
}
func (m *GrpcReplyMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReplyMaps.Marshal(b, m, deterministic)
}
func (m *GrpcReplyMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReplyMaps.Merge(m, src)
}
func (m *GrpcReplyMaps) XXX_Size() int {
	return xxx_messageInfo_GrpcReplyMaps.Size(m)
}
func (m *GrpcReplyMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReplyMaps.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReplyMaps proto.InternalMessageInfo

func (m *GrpcReplyMaps) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcReplyMaps) GetData() []*DataMap {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GrpcReplyMaps) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DataMap struct {
	Data                 map[string]string {{!}}protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"{{!}}
	XXX_NoUnkeyedLiteral struct{}          {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte            {{!}}json:"-"{{!}}
	XXX_sizecache        int32             {{!}}json:"-"{{!}}
}

func (m *DataMap) Reset()         { *m = DataMap{} }
func (m *DataMap) String() string { return proto.CompactTextString(m) }
func (*DataMap) ProtoMessage()    {}
func (*DataMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{10}
}

func (m *DataMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataMap.Unmarshal(m, b)
}
func (m *DataMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataMap.Marshal(b, m, deterministic)
}
func (m *DataMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataMap.Merge(m, src)
}
func (m *DataMap) XXX_Size() int {
	return xxx_messageInfo_DataMap.Size(m)
}
func (m *DataMap) XXX_DiscardUnknown() {
	xxx_messageInfo_DataMap.DiscardUnknown(m)
}

var xxx_messageInfo_DataMap proto.InternalMessageInfo

func (m *DataMap) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type GrpcReplyBool struct {
	Code                 int32    {{!}}protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"{{!}}
	Data                 bool     {{!}}protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"{{!}}
	Msg                  string   {{!}}protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GrpcReplyBool) Reset()         { *m = GrpcReplyBool{} }
func (m *GrpcReplyBool) String() string { return proto.CompactTextString(m) }
func (*GrpcReplyBool) ProtoMessage()    {}
func (*GrpcReplyBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{11}
}

func (m *GrpcReplyBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReplyBool.Unmarshal(m, b)
}
func (m *GrpcReplyBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReplyBool.Marshal(b, m, deterministic)
}
func (m *GrpcReplyBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReplyBool.Merge(m, src)
}
func (m *GrpcReplyBool) XXX_Size() int {
	return xxx_messageInfo_GrpcReplyBool.Size(m)
}
func (m *GrpcReplyBool) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReplyBool.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReplyBool proto.InternalMessageInfo

func (m *GrpcReplyBool) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcReplyBool) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

func (m *GrpcReplyBool) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GrpcReplyPage struct {
	Code                 int32    {{!}}protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"{{!}}
	Data                 *Page    {{!}}protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"{{!}}
	Msg                  string   {{!}}protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GrpcReplyPage) Reset()         { *m = GrpcReplyPage{} }
func (m *GrpcReplyPage) String() string { return proto.CompactTextString(m) }
func (*GrpcReplyPage) ProtoMessage()    {}
func (*GrpcReplyPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{12}
}

func (m *GrpcReplyPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReplyPage.Unmarshal(m, b)
}
func (m *GrpcReplyPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReplyPage.Marshal(b, m, deterministic)
}
func (m *GrpcReplyPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReplyPage.Merge(m, src)
}
func (m *GrpcReplyPage) XXX_Size() int {
	return xxx_messageInfo_GrpcReplyPage.Size(m)
}
func (m *GrpcReplyPage) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReplyPage.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReplyPage proto.InternalMessageInfo

func (m *GrpcReplyPage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcReplyPage) GetData() *Page {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GrpcReplyPage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Page struct {
	Total                int64      {{!}}protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"{{!}}
	List                 []*DataMap {{!}}protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"{{!}}
	PageTotal            int32      {{!}}protobuf:"varint,3,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{}   {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte     {{!}}json:"-"{{!}}
	XXX_sizecache        int32      {{!}}json:"-"{{!}}
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb943c226d71e9f, []int{13}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Page) GetList() []*DataMap {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Page) GetPageTotal() int32 {
	if m != nil {
		return m.PageTotal
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserByUserNameRequest)(nil), "passport.GetUserByUserNameRequest")
	proto.RegisterType((*GetUserByIdRequest)(nil), "passport.GetUserByIdRequest")
	proto.RegisterType((*GetUserByIdsRequest)(nil), "passport.GetUserByIdsRequest")
	proto.RegisterType((*GetUserInfoRequest)(nil), "passport.GetUserInfoRequest")
	proto.RegisterType((*SetUserKeyRequest)(nil), "passport.SetUserKeyRequest")
	proto.RegisterType((*SetUserKeyArrRequest)(nil), "passport.SetUserKeyArrRequest")
	proto.RegisterType((*GetUserListRequest)(nil), "passport.GetUserListRequest")
	proto.RegisterType((*GetUserIdsByUserNamesRequest)(nil), "passport.GetUserIdsByUserNamesRequest")
	proto.RegisterType((*GrpcReplyMap)(nil), "passport.GrpcReplyMap")
	proto.RegisterMapType((map[string]string)(nil), "passport.GrpcReplyMap.DataEntry")
	proto.RegisterType((*GrpcReplyMaps)(nil), "passport.GrpcReplyMaps")
	proto.RegisterType((*DataMap)(nil), "passport.DataMap")
	proto.RegisterMapType((map[string]string)(nil), "passport.DataMap.DataEntry")
	proto.RegisterType((*GrpcReplyBool)(nil), "passport.GrpcReplyBool")
	proto.RegisterType((*GrpcReplyPage)(nil), "passport.GrpcReplyPage")
	proto.RegisterType((*Page)(nil), "passport.Page")
}

func init() { proto.RegisterFile("passport/user.proto", fileDescriptor_9fb943c226d71e9f) }

var fileDescriptor_9fb943c226d71e9f = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x6b, 0x3b, 0x81, 0x74, 0xfa, 0xa1, 0x66, 0x1b, 0x20, 0x0a, 0x2d, 0x8a, 0x56, 0x7c,
	0xe4, 0x94, 0x4a, 0xa1, 0x02, 0x84, 0xb8, 0x10, 0xa5, 0x94, 0x50, 0x82, 0xc0, 0x85, 0x43, 0x25,
	0x24, 0xb4, 0xad, 0xb7, 0x95, 0x45, 0x12, 0x1b, 0xaf, 0x83, 0x64, 0x5e, 0x87, 0xb7, 0xe3, 0x29,
	0xd0, 0xac, 0xbd, 0xeb, 0x0d, 0x59, 0xe7, 0xc0, 0x25, 0x99, 0xd9, 0xcc, 0xfc, 0x76, 0x66, 0x3c,
	0xfe, 0x07, 0xf6, 0x63, 0x26, 0x44, 0x1c, 0x25, 0xe9, 0xd1, 0x42, 0xf0, 0xa4, 0x1f, 0x27, 0x51,
	0x1a, 0x91, 0x86, 0x3a, 0xa4, 0xcf, 0xa0, 0x7d, 0xca, 0xd3, 0x2f, 0x82, 0x27, 0xc3, 0x0c, 0x3f,
	0x3f, 0xb0, 0x19, 0xf7, 0xf9, 0x8f, 0x05, 0x17, 0x29, 0xe9, 0x40, 0x03, 0x73, 0xe6, 0x6c, 0xc6,
	0xdb, 0x4e, 0xd7, 0xe9, 0x6d, 0xfa, 0xda, 0xa7, 0x0f, 0x81, 0xe8, 0xbc, 0x71, 0xa0, 0x32, 0x76,
	0xc1, 0x0d, 0x03, 0x19, 0xeb, 0xf9, 0x6e, 0x18, 0xd0, 0x27, 0xb0, 0x6f, 0x44, 0x09, 0x15, 0xb6,
	0x07, 0x5e, 0x18, 0x88, 0x82, 0x89, 0x26, 0x1d, 0x6a, 0xdc, 0x78, 0x7e, 0x1d, 0x19, 0x71, 0x0b,
	0xcd, 0x43, 0x53, 0x95, 0x74, 0xc6, 0x33, 0xd1, 0x76, 0xcb, 0x92, 0xd0, 0xa7, 0x13, 0x68, 0x9e,
	0xe7, 0x8c, 0x33, 0x9e, 0x55, 0x23, 0xf6, 0xc0, 0xfb, 0xce, 0xb3, 0x22, 0x1b, 0x4d, 0xd2, 0x82,
	0xfa, 0x4f, 0x36, 0x5d, 0xf0, 0xb6, 0x27, 0xcf, 0x72, 0x87, 0xbe, 0x82, 0x56, 0x89, 0x7b, 0x9d,
	0x24, 0xd5, 0x44, 0x02, 0xb5, 0x70, 0x7e, 0x1d, 0x15, 0x48, 0x69, 0xd3, 0x91, 0x6e, 0xe8, 0x7d,
	0x28, 0x52, 0x95, 0x4b, 0xa0, 0x16, 0xb3, 0x1b, 0x5e, 0x24, 0x4b, 0x1b, 0x5b, 0xc2, 0xef, 0xf3,
	0xf0, 0x17, 0x97, 0x04, 0xcf, 0xd7, 0x3e, 0x3d, 0x86, 0x03, 0x35, 0x96, 0x40, 0x94, 0x0f, 0x48,
	0x0f, 0xb2, 0x05, 0x75, 0x7c, 0x1a, 0x6a, 0x94, 0xb9, 0x43, 0x7f, 0x3b, 0xb0, 0x7d, 0x9a, 0xc4,
	0x57, 0x3e, 0x8f, 0xa7, 0xd9, 0x84, 0xc5, 0x78, 0xed, 0x55, 0x14, 0xe4, 0xd7, 0xd6, 0x7d, 0x69,
	0x93, 0x63, 0xa8, 0x05, 0x2c, 0x65, 0x6d, 0xb7, 0xeb, 0xf5, 0xb6, 0x06, 0xdd, 0xbe, 0xda, 0x88,
	0xbe, 0x99, 0xd9, 0x1f, 0xb1, 0x94, 0x9d, 0xcc, 0xd3, 0x24, 0xf3, 0x65, 0x34, 0x36, 0x3f, 0x13,
	0x37, 0xc5, 0xa0, 0xd0, 0xec, 0x3c, 0x87, 0x4d, 0x1d, 0xa4, 0x66, 0xeb, 0x58, 0x66, 0xeb, 0x1a,
	0xb3, 0x7d, 0xe9, 0xbe, 0x70, 0xe8, 0x57, 0xd8, 0x31, 0xaf, 0x12, 0xd6, 0x2a, 0x1f, 0x2d, 0x55,
	0xd9, 0x2c, 0xab, 0xc4, 0x3b, 0x27, 0x2c, 0xae, 0x2a, 0x8b, 0x0a, 0xb8, 0x5d, 0x84, 0x90, 0xa3,
	0x82, 0xe1, 0x48, 0xc6, 0xfd, 0x15, 0xc6, 0xbf, 0x4d, 0xfe, 0x7f, 0x4b, 0x63, 0xa3, 0xa5, 0x61,
	0x14, 0x4d, 0xad, 0x2d, 0x11, 0xdd, 0x92, 0xd3, 0x6b, 0x54, 0xd6, 0x7f, 0x61, 0xa0, 0x3e, 0xe2,
	0x9a, 0xd8, 0x50, 0xd4, 0x40, 0x6d, 0x0d, 0x76, 0xcb, 0xce, 0x30, 0xa3, 0x12, 0x7d, 0x09, 0x35,
	0x49, 0x6c, 0x41, 0x3d, 0x8d, 0x52, 0x36, 0x2d, 0xb6, 0x31, 0x77, 0x70, 0xe2, 0xd3, 0x50, 0xa4,
	0x6b, 0x26, 0x8e, 0x3f, 0x93, 0x43, 0x00, 0xdc, 0xd2, 0x6f, 0x39, 0xc1, 0x93, 0x45, 0x6d, 0xe2,
	0xc9, 0x67, 0x3c, 0x18, 0xfc, 0xa9, 0x41, 0x43, 0xbd, 0xcd, 0xe4, 0x13, 0x34, 0x57, 0x34, 0x86,
	0x50, 0x63, 0xe3, 0x2a, 0x04, 0xa8, 0x73, 0xd7, 0xbe, 0x95, 0x74, 0x83, 0x9c, 0xc0, 0x96, 0x21,
	0x2c, 0xe4, 0xc0, 0x02, 0xd3, 0xaa, 0xb4, 0x06, 0xf3, 0x16, 0xb6, 0x4d, 0x7d, 0x22, 0x87, 0x56,
	0x8e, 0x7a, 0xdd, 0x3a, 0xf7, 0xec, 0x20, 0xb1, 0x54, 0x90, 0x6c, 0x79, 0xb5, 0x20, 0x43, 0xd7,
	0xd6, 0x14, 0x34, 0x02, 0x28, 0x45, 0x87, 0x18, 0xbb, 0xba, 0xa2, 0x6c, 0xd6, 0x62, 0x70, 0xe9,
	0xe8, 0x06, 0x79, 0x07, 0x3b, 0x4b, 0xd2, 0x45, 0x1e, 0xd8, 0x40, 0xa5, 0xa6, 0xad, 0x63, 0xbd,
	0xd1, 0x8d, 0xa1, 0x90, 0x59, 0x1a, 0x33, 0xf4, 0xcd, 0xca, 0xc1, 0x5d, 0xa3, 0x1b, 0xe4, 0x02,
	0xee, 0x58, 0xa5, 0x8c, 0x3c, 0x5e, 0x1d, 0x95, 0x4d, 0xeb, 0xaa, 0x87, 0x76, 0x79, 0x4b, 0xfe,
	0xa9, 0x3d, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x64, 0xd1, 0xb0, 0xeb, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInfoClient interface {
	GetUserByUserName(ctx context.Context, in *GetUserByUserNameRequest, opts ...grpc.CallOption) (*GrpcReplyMap, error)
}

type userInfoClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoClient(cc *grpc.ClientConn) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) GetUserByUserName(ctx context.Context, in *GetUserByUserNameRequest, opts ...grpc.CallOption) (*GrpcReplyMap, error) {
	out := new(GrpcReplyMap)
	err := c.cc.Invoke(ctx, "/passport.UserInfo/GetUserByUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


// UserInfoServer is the server API for UserInfo service.
type UserInfoServer interface {
	GetUserByUserName(context.Context, *GetUserByUserNameRequest) (*GrpcReplyMap, error)
}

// UnimplementedUserInfoServer can be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (*UnimplementedUserInfoServer) GetUserByUserName(ctx context.Context, req *GetUserByUserNameRequest) (*GrpcReplyMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUserName not implemented")
}


func RegisterUserInfoServer(s *grpc.Server, srv UserInfoServer) {
	s.RegisterService(&_UserInfo_serviceDesc, srv)
}

func _UserInfo_GetUserByUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserByUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passport.UserInfo/GetUserByUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserByUserName(ctx, req.(*GetUserByUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passport.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByUserName",
			Handler:    _UserInfo_GetUserByUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passport/user.proto",
}

`,
	}

	Files = append(Files, fc1)
}