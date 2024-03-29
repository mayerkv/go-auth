// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: grpc-service/auth-service.proto

package grpc_service

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AccountRole int32

const (
	AccountRole_USER  AccountRole = 0
	AccountRole_ADMIN AccountRole = 1
)

// Enum value maps for AccountRole.
var (
	AccountRole_name = map[int32]string{
		0: "USER",
		1: "ADMIN",
	}
	AccountRole_value = map[string]int32{
		"USER":  0,
		"ADMIN": 1,
	}
)

func (x AccountRole) Enum() *AccountRole {
	p := new(AccountRole)
	*p = x
	return p
}

func (x AccountRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountRole) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_service_auth_service_proto_enumTypes[0].Descriptor()
}

func (AccountRole) Type() protoreflect.EnumType {
	return &file_grpc_service_auth_service_proto_enumTypes[0]
}

func (x AccountRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountRole.Descriptor instead.
func (AccountRole) EnumDescriptor() ([]byte, []int) {
	return file_grpc_service_auth_service_proto_rawDescGZIP(), []int{0}
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string      `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string      `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	UserId string      `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Role   AccountRole `protobuf:"varint,4,opt,name=Role,proto3,enum=AccountRole" json:"Role,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_grpc_service_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAccountRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateAccountRequest) GetRole() AccountRole {
	if x != nil {
		return x.Role
	}
	return AccountRole_USER
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_service_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_service_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_grpc_service_auth_service_proto_rawDescGZIP(), []int{1}
}

var File_grpc_service_auth_service_proto protoreflect.FileDescriptor

var file_grpc_service_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a,
	0x22, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x01, 0x32, 0x4f, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x79, 0x65, 0x72, 0x6b, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_service_auth_service_proto_rawDescOnce sync.Once
	file_grpc_service_auth_service_proto_rawDescData = file_grpc_service_auth_service_proto_rawDesc
)

func file_grpc_service_auth_service_proto_rawDescGZIP() []byte {
	file_grpc_service_auth_service_proto_rawDescOnce.Do(func() {
		file_grpc_service_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_service_auth_service_proto_rawDescData)
	})
	return file_grpc_service_auth_service_proto_rawDescData
}

var file_grpc_service_auth_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_service_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_service_auth_service_proto_goTypes = []interface{}{
	(AccountRole)(0),              // 0: AccountRole
	(*CreateAccountRequest)(nil),  // 1: CreateAccountRequest
	(*CreateAccountResponse)(nil), // 2: CreateAccountResponse
}
var file_grpc_service_auth_service_proto_depIdxs = []int32{
	0, // 0: CreateAccountRequest.Role:type_name -> AccountRole
	1, // 1: AuthService.CreateAccount:input_type -> CreateAccountRequest
	2, // 2: AuthService.CreateAccount:output_type -> CreateAccountResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_service_auth_service_proto_init() }
func file_grpc_service_auth_service_proto_init() {
	if File_grpc_service_auth_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_service_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_grpc_service_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
			RawDescriptor: file_grpc_service_auth_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_service_auth_service_proto_goTypes,
		DependencyIndexes: file_grpc_service_auth_service_proto_depIdxs,
		EnumInfos:         file_grpc_service_auth_service_proto_enumTypes,
		MessageInfos:      file_grpc_service_auth_service_proto_msgTypes,
	}.Build()
	File_grpc_service_auth_service_proto = out.File
	file_grpc_service_auth_service_proto_rawDesc = nil
	file_grpc_service_auth_service_proto_goTypes = nil
	file_grpc_service_auth_service_proto_depIdxs = nil
}
