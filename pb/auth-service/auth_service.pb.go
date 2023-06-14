// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: auth_service.proto

package auth_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName  string `protobuf:"bytes,1,opt,name=FullName,proto3" json:"FullName,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	PhoneNum  string `protobuf:"bytes,4,opt,name=PhoneNum,proto3" json:"PhoneNum,omitempty"`
	BirthDate string `protobuf:"bytes,5,opt,name=BirthDate,proto3" json:"BirthDate,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpRequest) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

func (x *SignUpRequest) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SignUpResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	User_ID int32  `protobuf:"varint,4,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SignInResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SignInResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignInResponse) GetUser_ID() int32 {
	if x != nil {
		return x.User_ID
	}
	return 0
}

type AUC_SRVC_DATA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId string `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	LotId     int64  `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
}

func (x *AUC_SRVC_DATA) Reset() {
	*x = AUC_SRVC_DATA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AUC_SRVC_DATA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AUC_SRVC_DATA) ProtoMessage() {}

func (x *AUC_SRVC_DATA) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AUC_SRVC_DATA.ProtoReflect.Descriptor instead.
func (*AUC_SRVC_DATA) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *AUC_SRVC_DATA) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

func (x *AUC_SRVC_DATA) GetLotId() int64 {
	if x != nil {
		return x.LotId
	}
	return 0
}

type MNG_SRVC_DATA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId string `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
}

func (x *MNG_SRVC_DATA) Reset() {
	*x = MNG_SRVC_DATA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MNG_SRVC_DATA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MNG_SRVC_DATA) ProtoMessage() {}

func (x *MNG_SRVC_DATA) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MNG_SRVC_DATA.ProtoReflect.Descriptor instead.
func (*MNG_SRVC_DATA) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *MNG_SRVC_DATA) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

type RUN_SRVC_DATA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId string `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	LotId     int64  `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
}

func (x *RUN_SRVC_DATA) Reset() {
	*x = RUN_SRVC_DATA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RUN_SRVC_DATA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RUN_SRVC_DATA) ProtoMessage() {}

func (x *RUN_SRVC_DATA) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RUN_SRVC_DATA.ProtoReflect.Descriptor instead.
func (*RUN_SRVC_DATA) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{6}
}

func (x *RUN_SRVC_DATA) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

func (x *RUN_SRVC_DATA) GetLotId() int64 {
	if x != nil {
		return x.LotId
	}
	return 0
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Types that are assignable to SRVC_DATA:
	//
	//	*ValidateRequest_AUC_SRVC_DATA
	//	*ValidateRequest_MNG_SRVC_DATA
	//	*ValidateRequest_RUN_SRVC_DATA
	SRVC_DATA isValidateRequest_SRVC_DATA `protobuf_oneof:"SRVC_DATA"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ValidateRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (m *ValidateRequest) GetSRVC_DATA() isValidateRequest_SRVC_DATA {
	if m != nil {
		return m.SRVC_DATA
	}
	return nil
}

func (x *ValidateRequest) GetAUC_SRVC_DATA() *AUC_SRVC_DATA {
	if x, ok := x.GetSRVC_DATA().(*ValidateRequest_AUC_SRVC_DATA); ok {
		return x.AUC_SRVC_DATA
	}
	return nil
}

func (x *ValidateRequest) GetMNG_SRVC_DATA() *MNG_SRVC_DATA {
	if x, ok := x.GetSRVC_DATA().(*ValidateRequest_MNG_SRVC_DATA); ok {
		return x.MNG_SRVC_DATA
	}
	return nil
}

func (x *ValidateRequest) GetRUN_SRVC_DATA() *RUN_SRVC_DATA {
	if x, ok := x.GetSRVC_DATA().(*ValidateRequest_RUN_SRVC_DATA); ok {
		return x.RUN_SRVC_DATA
	}
	return nil
}

type isValidateRequest_SRVC_DATA interface {
	isValidateRequest_SRVC_DATA()
}

type ValidateRequest_AUC_SRVC_DATA struct {
	AUC_SRVC_DATA *AUC_SRVC_DATA `protobuf:"bytes,3,opt,name=AUC_SRVC_DATA,json=AUCSRVCDATA,proto3,oneof"`
}

type ValidateRequest_MNG_SRVC_DATA struct {
	MNG_SRVC_DATA *MNG_SRVC_DATA `protobuf:"bytes,4,opt,name=MNG_SRVC_DATA,json=MNGSRVCDATA,proto3,oneof"`
}

type ValidateRequest_RUN_SRVC_DATA struct {
	RUN_SRVC_DATA *RUN_SRVC_DATA `protobuf:"bytes,5,opt,name=RUN_SRVC_DATA,json=RUNSRVCDATA,proto3,oneof"`
}

func (*ValidateRequest_AUC_SRVC_DATA) isValidateRequest_SRVC_DATA() {}

func (*ValidateRequest_MNG_SRVC_DATA) isValidateRequest_SRVC_DATA() {}

func (*ValidateRequest_RUN_SRVC_DATA) isValidateRequest_SRVC_DATA() {}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	UserID   int64  `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FullName string `protobuf:"bytes,4,opt,name=FullName,proto3" json:"FullName,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ValidateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ValidateResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ValidateResponse) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *ValidateResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x97,
	0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6d, 0x0a, 0x0e, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x0d, 0x41, 0x55,
	0x43, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f, 0x74, 0x49,
	0x64, 0x22, 0x2e, 0x0a, 0x0d, 0x4d, 0x4e, 0x47, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x45, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x8b, 0x02, 0x0a, 0x0f, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0d,
	0x41, 0x55, 0x43, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x55, 0x43, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x48, 0x00, 0x52, 0x0b,
	0x41, 0x55, 0x43, 0x53, 0x52, 0x56, 0x43, 0x44, 0x41, 0x54, 0x41, 0x12, 0x3d, 0x0a, 0x0d, 0x4d,
	0x4e, 0x47, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x4e,
	0x47, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x48, 0x00, 0x52, 0x0b, 0x4d,
	0x4e, 0x47, 0x53, 0x52, 0x56, 0x43, 0x44, 0x41, 0x54, 0x41, 0x12, 0x3d, 0x0a, 0x0d, 0x52, 0x55,
	0x4e, 0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x55, 0x4e,
	0x5f, 0x53, 0x52, 0x56, 0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x48, 0x00, 0x52, 0x0b, 0x52, 0x55,
	0x4e, 0x53, 0x52, 0x56, 0x43, 0x44, 0x41, 0x54, 0x41, 0x42, 0x0b, 0x0a, 0x09, 0x53, 0x52, 0x56,
	0x43, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_service_proto_rawDescOnce sync.Once
	file_auth_service_proto_rawDescData = file_auth_service_proto_rawDesc
)

func file_auth_service_proto_rawDescGZIP() []byte {
	file_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_service_proto_rawDescData)
	})
	return file_auth_service_proto_rawDescData
}

var file_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auth_service_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),    // 0: protobuf.SignUpRequest
	(*SignUpResponse)(nil),   // 1: protobuf.SignUpResponse
	(*SignInRequest)(nil),    // 2: protobuf.SignInRequest
	(*SignInResponse)(nil),   // 3: protobuf.SignInResponse
	(*AUC_SRVC_DATA)(nil),    // 4: protobuf.AUC_SRVC_DATA
	(*MNG_SRVC_DATA)(nil),    // 5: protobuf.MNG_SRVC_DATA
	(*RUN_SRVC_DATA)(nil),    // 6: protobuf.RUN_SRVC_DATA
	(*ValidateRequest)(nil),  // 7: protobuf.ValidateRequest
	(*ValidateResponse)(nil), // 8: protobuf.ValidateResponse
}
var file_auth_service_proto_depIdxs = []int32{
	4, // 0: protobuf.ValidateRequest.AUC_SRVC_DATA:type_name -> protobuf.AUC_SRVC_DATA
	5, // 1: protobuf.ValidateRequest.MNG_SRVC_DATA:type_name -> protobuf.MNG_SRVC_DATA
	6, // 2: protobuf.ValidateRequest.RUN_SRVC_DATA:type_name -> protobuf.RUN_SRVC_DATA
	0, // 3: protobuf.AuthService.SignUp:input_type -> protobuf.SignUpRequest
	2, // 4: protobuf.AuthService.SignIn:input_type -> protobuf.SignInRequest
	7, // 5: protobuf.AuthService.Validate:input_type -> protobuf.ValidateRequest
	1, // 6: protobuf.AuthService.SignUp:output_type -> protobuf.SignUpResponse
	3, // 7: protobuf.AuthService.SignIn:output_type -> protobuf.SignInResponse
	8, // 8: protobuf.AuthService.Validate:output_type -> protobuf.ValidateResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AUC_SRVC_DATA); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MNG_SRVC_DATA); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RUN_SRVC_DATA); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_auth_service_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ValidateRequest_AUC_SRVC_DATA)(nil),
		(*ValidateRequest_MNG_SRVC_DATA)(nil),
		(*ValidateRequest_RUN_SRVC_DATA)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
		MessageInfos:      file_auth_service_proto_msgTypes,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}
