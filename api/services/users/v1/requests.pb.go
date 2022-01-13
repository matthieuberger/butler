// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: services/users/v1/requests.proto

package users

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[0]
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
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{0}
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

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastName    string `protobuf:"bytes,1,opt,name=lastName,proto3" json:"lastName,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CompanyName string `protobuf:"bytes,5,opt,name=companyName,proto3" json:"companyName,omitempty"`
	CompanyRole string `protobuf:"bytes,6,opt,name=companyRole,proto3" json:"companyRole,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[1]
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
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignUpRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
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

func (x *SignUpRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *SignUpRequest) GetCompanyRole() string {
	if x != nil {
		return x.CompanyRole
	}
	return ""
}

type SignUpWithInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvitationId string                                  `protobuf:"bytes,1,opt,name=invitationId,proto3" json:"invitationId,omitempty"`
	Infos        *SignUpWithInvitationRequest_SignupInfo `protobuf:"bytes,2,opt,name=infos,proto3" json:"infos,omitempty"`
}

func (x *SignUpWithInvitationRequest) Reset() {
	*x = SignUpWithInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpWithInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpWithInvitationRequest) ProtoMessage() {}

func (x *SignUpWithInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpWithInvitationRequest.ProtoReflect.Descriptor instead.
func (*SignUpWithInvitationRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpWithInvitationRequest) GetInvitationId() string {
	if x != nil {
		return x.InvitationId
	}
	return ""
}

func (x *SignUpWithInvitationRequest) GetInfos() *SignUpWithInvitationRequest_SignupInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type SignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{3}
}

func (x *SignOutRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type CreateOrganizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserRole string `protobuf:"bytes,2,opt,name=userRole,proto3" json:"userRole,omitempty"`
}

func (x *CreateOrganizationRequest) Reset() {
	*x = CreateOrganizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrganizationRequest) ProtoMessage() {}

func (x *CreateOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrganizationRequest.ProtoReflect.Descriptor instead.
func (*CreateOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOrganizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrganizationRequest) GetUserRole() string {
	if x != nil {
		return x.UserRole
	}
	return ""
}

type CreateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId string                                `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Workspace      *CreateWorkspaceRequest_WorkspaceInfo `protobuf:"bytes,2,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *CreateWorkspaceRequest) Reset() {
	*x = CreateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceRequest) ProtoMessage() {}

func (x *CreateWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{6}
}

func (x *CreateWorkspaceRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *CreateWorkspaceRequest) GetWorkspace() *CreateWorkspaceRequest_WorkspaceInfo {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type InviteInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role      string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *InviteInfos) Reset() {
	*x = InviteInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteInfos) ProtoMessage() {}

func (x *InviteInfos) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteInfos.ProtoReflect.Descriptor instead.
func (*InviteInfos) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{7}
}

func (x *InviteInfos) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *InviteInfos) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *InviteInfos) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InviteInfos) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type InviteOrganizationMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invitation     *InviteInfos `protobuf:"bytes,1,opt,name=invitation,proto3" json:"invitation,omitempty"`
	OrganizationId string       `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *InviteOrganizationMemberRequest) Reset() {
	*x = InviteOrganizationMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteOrganizationMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteOrganizationMemberRequest) ProtoMessage() {}

func (x *InviteOrganizationMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteOrganizationMemberRequest.ProtoReflect.Descriptor instead.
func (*InviteOrganizationMemberRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{8}
}

func (x *InviteOrganizationMemberRequest) GetInvitation() *InviteInfos {
	if x != nil {
		return x.Invitation
	}
	return nil
}

func (x *InviteOrganizationMemberRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type InviteWorkspaceMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invitation  *InviteInfos `protobuf:"bytes,1,opt,name=invitation,proto3" json:"invitation,omitempty"`
	WorkspaceId string       `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *InviteWorkspaceMemberRequest) Reset() {
	*x = InviteWorkspaceMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteWorkspaceMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteWorkspaceMemberRequest) ProtoMessage() {}

func (x *InviteWorkspaceMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteWorkspaceMemberRequest.ProtoReflect.Descriptor instead.
func (*InviteWorkspaceMemberRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{9}
}

func (x *InviteWorkspaceMemberRequest) GetInvitation() *InviteInfos {
	if x != nil {
		return x.Invitation
	}
	return nil
}

func (x *InviteWorkspaceMemberRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type CompleteOnboardingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *CompleteOnboardingRequest) Reset() {
	*x = CompleteOnboardingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOnboardingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOnboardingRequest) ProtoMessage() {}

func (x *CompleteOnboardingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOnboardingRequest.ProtoReflect.Descriptor instead.
func (*CompleteOnboardingRequest) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{10}
}

func (x *CompleteOnboardingRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type SignUpWithInvitationRequest_SignupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Binding:
	//	*SignUpWithInvitationRequest_SignupInfo_OrganizationId
	//	*SignUpWithInvitationRequest_SignupInfo_WorkspaceId
	Binding  isSignUpWithInvitationRequest_SignupInfo_Binding `protobuf_oneof:"binding"`
	Token    string                                           `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Password string                                           `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpWithInvitationRequest_SignupInfo) Reset() {
	*x = SignUpWithInvitationRequest_SignupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpWithInvitationRequest_SignupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpWithInvitationRequest_SignupInfo) ProtoMessage() {}

func (x *SignUpWithInvitationRequest_SignupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpWithInvitationRequest_SignupInfo.ProtoReflect.Descriptor instead.
func (*SignUpWithInvitationRequest_SignupInfo) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{2, 0}
}

func (m *SignUpWithInvitationRequest_SignupInfo) GetBinding() isSignUpWithInvitationRequest_SignupInfo_Binding {
	if m != nil {
		return m.Binding
	}
	return nil
}

func (x *SignUpWithInvitationRequest_SignupInfo) GetOrganizationId() string {
	if x, ok := x.GetBinding().(*SignUpWithInvitationRequest_SignupInfo_OrganizationId); ok {
		return x.OrganizationId
	}
	return ""
}

func (x *SignUpWithInvitationRequest_SignupInfo) GetWorkspaceId() string {
	if x, ok := x.GetBinding().(*SignUpWithInvitationRequest_SignupInfo_WorkspaceId); ok {
		return x.WorkspaceId
	}
	return ""
}

func (x *SignUpWithInvitationRequest_SignupInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignUpWithInvitationRequest_SignupInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type isSignUpWithInvitationRequest_SignupInfo_Binding interface {
	isSignUpWithInvitationRequest_SignupInfo_Binding()
}

type SignUpWithInvitationRequest_SignupInfo_OrganizationId struct {
	OrganizationId string `protobuf:"bytes,2,opt,name=organizationId,proto3,oneof"`
}

type SignUpWithInvitationRequest_SignupInfo_WorkspaceId struct {
	WorkspaceId string `protobuf:"bytes,3,opt,name=workspaceId,proto3,oneof"`
}

func (*SignUpWithInvitationRequest_SignupInfo_OrganizationId) isSignUpWithInvitationRequest_SignupInfo_Binding() {
}

func (*SignUpWithInvitationRequest_SignupInfo_WorkspaceId) isSignUpWithInvitationRequest_SignupInfo_Binding() {
}

type CreateWorkspaceRequest_WorkspaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateWorkspaceRequest_WorkspaceInfo) Reset() {
	*x = CreateWorkspaceRequest_WorkspaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_users_v1_requests_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceRequest_WorkspaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceRequest_WorkspaceInfo) ProtoMessage() {}

func (x *CreateWorkspaceRequest_WorkspaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_users_v1_requests_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceRequest_WorkspaceInfo.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceRequest_WorkspaceInfo) Descriptor() ([]byte, []int) {
	return file_services_users_v1_requests_proto_rawDescGZIP(), []int{6, 0}
}

func (x *CreateWorkspaceRequest_WorkspaceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWorkspaceRequest_WorkspaceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_services_users_v1_requests_proto protoreflect.FileDescriptor

var file_services_users_v1_requests_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x08, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xf5, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x08,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x6f, 0x6c, 0x65,
	0x22, 0x9d, 0x02, 0x0a, 0x1b, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x57,
	0x69, 0x74, 0x68, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x97, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x22, 0x32, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4b, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x1a, 0x45, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x0b, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x7b, 0x0a, 0x1f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x1c,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x44, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x74, 0x6c, 0x65, 0x72, 0x68, 0x71, 0x2f, 0x62, 0x75,
	0x74, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_users_v1_requests_proto_rawDescOnce sync.Once
	file_services_users_v1_requests_proto_rawDescData = file_services_users_v1_requests_proto_rawDesc
)

func file_services_users_v1_requests_proto_rawDescGZIP() []byte {
	file_services_users_v1_requests_proto_rawDescOnce.Do(func() {
		file_services_users_v1_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_users_v1_requests_proto_rawDescData)
	})
	return file_services_users_v1_requests_proto_rawDescData
}

var file_services_users_v1_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_services_users_v1_requests_proto_goTypes = []interface{}{
	(*SignInRequest)(nil),                          // 0: v1.SignInRequest
	(*SignUpRequest)(nil),                          // 1: v1.SignUpRequest
	(*SignUpWithInvitationRequest)(nil),            // 2: v1.SignUpWithInvitationRequest
	(*SignOutRequest)(nil),                         // 3: v1.SignOutRequest
	(*RefreshRequest)(nil),                         // 4: v1.RefreshRequest
	(*CreateOrganizationRequest)(nil),              // 5: v1.CreateOrganizationRequest
	(*CreateWorkspaceRequest)(nil),                 // 6: v1.CreateWorkspaceRequest
	(*InviteInfos)(nil),                            // 7: v1.InviteInfos
	(*InviteOrganizationMemberRequest)(nil),        // 8: v1.InviteOrganizationMemberRequest
	(*InviteWorkspaceMemberRequest)(nil),           // 9: v1.InviteWorkspaceMemberRequest
	(*CompleteOnboardingRequest)(nil),              // 10: v1.CompleteOnboardingRequest
	(*SignUpWithInvitationRequest_SignupInfo)(nil), // 11: v1.SignUpWithInvitationRequest.SignupInfo
	(*CreateWorkspaceRequest_WorkspaceInfo)(nil),   // 12: v1.CreateWorkspaceRequest.WorkspaceInfo
}
var file_services_users_v1_requests_proto_depIdxs = []int32{
	11, // 0: v1.SignUpWithInvitationRequest.infos:type_name -> v1.SignUpWithInvitationRequest.SignupInfo
	12, // 1: v1.CreateWorkspaceRequest.workspace:type_name -> v1.CreateWorkspaceRequest.WorkspaceInfo
	7,  // 2: v1.InviteOrganizationMemberRequest.invitation:type_name -> v1.InviteInfos
	7,  // 3: v1.InviteWorkspaceMemberRequest.invitation:type_name -> v1.InviteInfos
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_services_users_v1_requests_proto_init() }
func file_services_users_v1_requests_proto_init() {
	if File_services_users_v1_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_users_v1_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_users_v1_requests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_users_v1_requests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpWithInvitationRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrganizationRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteInfos); i {
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
		file_services_users_v1_requests_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteOrganizationMemberRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteWorkspaceMemberRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOnboardingRequest); i {
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
		file_services_users_v1_requests_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpWithInvitationRequest_SignupInfo); i {
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
		file_services_users_v1_requests_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceRequest_WorkspaceInfo); i {
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
	file_services_users_v1_requests_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*SignUpWithInvitationRequest_SignupInfo_OrganizationId)(nil),
		(*SignUpWithInvitationRequest_SignupInfo_WorkspaceId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_users_v1_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_users_v1_requests_proto_goTypes,
		DependencyIndexes: file_services_users_v1_requests_proto_depIdxs,
		MessageInfos:      file_services_users_v1_requests_proto_msgTypes,
	}.Build()
	File_services_users_v1_requests_proto = out.File
	file_services_users_v1_requests_proto_rawDesc = nil
	file_services_users_v1_requests_proto_goTypes = nil
	file_services_users_v1_requests_proto_depIdxs = nil
}
