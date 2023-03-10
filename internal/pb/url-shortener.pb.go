// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: url-shortener.proto

package pb

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

type GetLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortenedString string `protobuf:"bytes,1,opt,name=shortened_string,json=shortenedString,proto3" json:"shortened_string,omitempty"`
}

func (x *GetLinkRequest) Reset() {
	*x = GetLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkRequest) ProtoMessage() {}

func (x *GetLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkRequest.ProtoReflect.Descriptor instead.
func (*GetLinkRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *GetLinkRequest) GetShortenedString() string {
	if x != nil {
		return x.ShortenedString
	}
	return ""
}

type GetLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetLinkResponse) Reset() {
	*x = GetLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkResponse) ProtoMessage() {}

func (x *GetLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkResponse.ProtoReflect.Descriptor instead.
func (*GetLinkResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetLinkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url             string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ShortenedString string `protobuf:"bytes,2,opt,name=shortened_string,json=shortenedString,proto3" json:"shortened_string,omitempty"`
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLinkRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateLinkRequest) GetShortenedString() string {
	if x != nil {
		return x.ShortenedString
	}
	return ""
}

type CreateLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortenedString string `protobuf:"bytes,1,opt,name=shortened_string,json=shortenedString,proto3" json:"shortened_string,omitempty"`
}

func (x *CreateLinkResponse) Reset() {
	*x = CreateLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkResponse) ProtoMessage() {}

func (x *CreateLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateLinkResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLinkResponse) GetShortenedString() string {
	if x != nil {
		return x.ShortenedString
	}
	return ""
}

type GetLinkUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortenedString string `protobuf:"bytes,1,opt,name=shortened_string,json=shortenedString,proto3" json:"shortened_string,omitempty"`
}

func (x *GetLinkUserRequest) Reset() {
	*x = GetLinkUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkUserRequest) ProtoMessage() {}

func (x *GetLinkUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkUserRequest.ProtoReflect.Descriptor instead.
func (*GetLinkUserRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *GetLinkUserRequest) GetShortenedString() string {
	if x != nil {
		return x.ShortenedString
	}
	return ""
}

type GetLinkUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetLinkUserResponse) Reset() {
	*x = GetLinkUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkUserResponse) ProtoMessage() {}

func (x *GetLinkUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkUserResponse.ProtoReflect.Descriptor instead.
func (*GetLinkUserResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *GetLinkUserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_proto_rawDescGZIP(), []int{7}
}

var File_url_shortener_proto protoreflect.FileDescriptor

var file_url_shortener_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x50, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xdc, 0x02, 0x0a, 0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x2e, 0x75,
	0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x72, 0x6c,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72,
	0x69, 0x61, 0x33, 0x70, 0x70, 0x70, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_shortener_proto_rawDescOnce sync.Once
	file_url_shortener_proto_rawDescData = file_url_shortener_proto_rawDesc
)

func file_url_shortener_proto_rawDescGZIP() []byte {
	file_url_shortener_proto_rawDescOnce.Do(func() {
		file_url_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_shortener_proto_rawDescData)
	})
	return file_url_shortener_proto_rawDescData
}

var file_url_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_url_shortener_proto_goTypes = []interface{}{
	(*GetLinkRequest)(nil),      // 0: url_shortener.GetLinkRequest
	(*GetLinkResponse)(nil),     // 1: url_shortener.GetLinkResponse
	(*CreateLinkRequest)(nil),   // 2: url_shortener.CreateLinkRequest
	(*CreateLinkResponse)(nil),  // 3: url_shortener.CreateLinkResponse
	(*GetLinkUserRequest)(nil),  // 4: url_shortener.GetLinkUserRequest
	(*GetLinkUserResponse)(nil), // 5: url_shortener.GetLinkUserResponse
	(*CreateUserRequest)(nil),   // 6: url_shortener.CreateUserRequest
	(*CreateUserResponse)(nil),  // 7: url_shortener.CreateUserResponse
}
var file_url_shortener_proto_depIdxs = []int32{
	0, // 0: url_shortener.UrlShortener.GetLink:input_type -> url_shortener.GetLinkRequest
	2, // 1: url_shortener.UrlShortener.CreateLink:input_type -> url_shortener.CreateLinkRequest
	4, // 2: url_shortener.UrlShortener.GetLinkUser:input_type -> url_shortener.GetLinkUserRequest
	6, // 3: url_shortener.UrlShortener.CreateUser:input_type -> url_shortener.CreateUserRequest
	1, // 4: url_shortener.UrlShortener.GetLink:output_type -> url_shortener.GetLinkResponse
	3, // 5: url_shortener.UrlShortener.CreateLink:output_type -> url_shortener.CreateLinkResponse
	5, // 6: url_shortener.UrlShortener.GetLinkUser:output_type -> url_shortener.GetLinkUserResponse
	7, // 7: url_shortener.UrlShortener.CreateUser:output_type -> url_shortener.CreateUserResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_shortener_proto_init() }
func file_url_shortener_proto_init() {
	if File_url_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkRequest); i {
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
		file_url_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkResponse); i {
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
		file_url_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkRequest); i {
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
		file_url_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkResponse); i {
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
		file_url_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkUserRequest); i {
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
		file_url_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkUserResponse); i {
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
		file_url_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_url_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_url_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_shortener_proto_goTypes,
		DependencyIndexes: file_url_shortener_proto_depIdxs,
		MessageInfos:      file_url_shortener_proto_msgTypes,
	}.Build()
	File_url_shortener_proto = out.File
	file_url_shortener_proto_rawDesc = nil
	file_url_shortener_proto_goTypes = nil
	file_url_shortener_proto_depIdxs = nil
}