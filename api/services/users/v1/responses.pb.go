// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: services/users/v1/responses.proto

package users

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

// The response message containing the User.
type AuthenticatedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *AuthenticatedUser) Reset() {
	*x = AuthenticatedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedUser) ProtoMessage() {}

func (x *AuthenticatedUser) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedUser.ProtoReflect.Descriptor instead.
func (*AuthenticatedUser) Descriptor() ([]byte, []int) {
	return file_services_users_v1_responses_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticatedUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AuthenticatedUser) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthenticatedUser) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type OrganizationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organizations []*Organization `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty"`
}

func (x *OrganizationListResponse) Reset() {
	*x = OrganizationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationListResponse) ProtoMessage() {}

func (x *OrganizationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationListResponse.ProtoReflect.Descriptor instead.
func (*OrganizationListResponse) Descriptor() ([]byte, []int) {
	return file_services_users_v1_responses_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationListResponse) GetOrganizations() []*Organization {
	if x != nil {
		return x.Organizations
	}
	return nil
}

type OrganizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization *Organization `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *OrganizationResponse) Reset() {
	*x = OrganizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationResponse) ProtoMessage() {}

func (x *OrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_responses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationResponse.ProtoReflect.Descriptor instead.
func (*OrganizationResponse) Descriptor() ([]byte, []int) {
	return file_services_users_v1_responses_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationResponse) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

type WorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *WorkspaceResponse) Reset() {
	*x = WorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_responses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceResponse) ProtoMessage() {}

func (x *WorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_responses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceResponse.ProtoReflect.Descriptor instead.
func (*WorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_services_users_v1_responses_proto_rawDescGZIP(), []int{3}
}

func (x *WorkspaceResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type InvitationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invitations []*Invitation `protobuf:"bytes,1,rep,name=invitations,proto3" json:"invitations,omitempty"`
}

func (x *InvitationListResponse) Reset() {
	*x = InvitationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_responses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationListResponse) ProtoMessage() {}

func (x *InvitationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_responses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationListResponse.ProtoReflect.Descriptor instead.
func (*InvitationListResponse) Descriptor() ([]byte, []int) {
	return file_services_users_v1_responses_proto_rawDescGZIP(), []int{4}
}

func (x *InvitationListResponse) GetInvitations() []*Invitation {
	if x != nil {
		return x.Invitations
	}
	return nil
}

var File_services_users_v1_responses_proto protoreflect.FileDescriptor

var file_services_users_v1_responses_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x11, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x52, 0x0a, 0x18, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4c, 0x0a, 0x14, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x75, 0x74, 0x6c, 0x65, 0x72, 0x68, 0x71, 0x2f, 0x62, 0x75, 0x74, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_users_v1_responses_proto_rawDescOnce sync.Once
	file_services_users_v1_responses_proto_rawDescData = file_services_users_v1_responses_proto_rawDesc
)

func file_services_users_v1_responses_proto_rawDescGZIP() []byte {
	file_services_users_v1_responses_proto_rawDescOnce.Do(func() {
		file_services_users_v1_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_users_v1_responses_proto_rawDescData)
	})
	return file_services_users_v1_responses_proto_rawDescData
}

var file_services_users_v1_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_users_v1_responses_proto_goTypes = []interface{}{
	(*AuthenticatedUser)(nil),        // 0: v1.AuthenticatedUser
	(*OrganizationListResponse)(nil), // 1: v1.OrganizationListResponse
	(*OrganizationResponse)(nil),     // 2: v1.OrganizationResponse
	(*WorkspaceResponse)(nil),        // 3: v1.WorkspaceResponse
	(*InvitationListResponse)(nil),   // 4: v1.InvitationListResponse
	(*User)(nil),                     // 5: v1.User
	(*Organization)(nil),             // 6: v1.Organization
	(*Workspace)(nil),                // 7: v1.Workspace
	(*Invitation)(nil),               // 8: v1.Invitation
}
var file_services_users_v1_responses_proto_depIdxs = []int32{
	5, // 0: v1.AuthenticatedUser.user:type_name -> v1.User
	6, // 1: v1.OrganizationListResponse.organizations:type_name -> v1.Organization
	6, // 2: v1.OrganizationResponse.organization:type_name -> v1.Organization
	7, // 3: v1.WorkspaceResponse.workspace:type_name -> v1.Workspace
	8, // 4: v1.InvitationListResponse.invitations:type_name -> v1.Invitation
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_users_v1_responses_proto_init() }
func file_services_users_v1_responses_proto_init() {
	if File_services_users_v1_responses_proto != nil {
		return
	}
	file_services_users_v1_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_users_v1_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedUser); i {
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
		file_services_users_v1_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationListResponse); i {
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
		file_services_users_v1_responses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationResponse); i {
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
		file_services_users_v1_responses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceResponse); i {
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
		file_services_users_v1_responses_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationListResponse); i {
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
			RawDescriptor: file_services_users_v1_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_users_v1_responses_proto_goTypes,
		DependencyIndexes: file_services_users_v1_responses_proto_depIdxs,
		MessageInfos:      file_services_users_v1_responses_proto_msgTypes,
	}.Build()
	File_services_users_v1_responses_proto = out.File
	file_services_users_v1_responses_proto_rawDesc = nil
	file_services_users_v1_responses_proto_goTypes = nil
	file_services_users_v1_responses_proto_depIdxs = nil
}
