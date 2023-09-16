// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: simoompb/v1/simoom.proto

package simoompb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckHealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckHealthRequest) Reset() {
	*x = CheckHealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckHealthRequest) ProtoMessage() {}

func (x *CheckHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckHealthRequest.ProtoReflect.Descriptor instead.
func (*CheckHealthRequest) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{0}
}

type CheckHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision string `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *CheckHealthResponse) Reset() {
	*x = CheckHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckHealthResponse) ProtoMessage() {}

func (x *CheckHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckHealthResponse.ProtoReflect.Descriptor instead.
func (*CheckHealthResponse) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{1}
}

func (x *CheckHealthResponse) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{3}
}

type UpdateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Color      *string `protobuf:"bytes,3,opt,name=color,proto3,oneof" json:"color,omitempty"`
	IsArchived *string `protobuf:"bytes,4,opt,name=is_archived,json=isArchived,proto3,oneof" json:"is_archived,omitempty"`
}

func (x *UpdateProjectRequest) Reset() {
	*x = UpdateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectRequest) ProtoMessage() {}

func (x *UpdateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProjectRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateProjectRequest) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *UpdateProjectRequest) GetIsArchived() string {
	if x != nil && x.IsArchived != nil {
		return *x.IsArchived
	}
	return ""
}

type DeleteProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProjectRequest) Reset() {
	*x = DeleteProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectRequest) ProtoMessage() {}

func (x *DeleteProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color      string                 `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	IsArchived bool                   `protobuf:"varint,4,opt,name=is_archived,json=isArchived,proto3" json:"is_archived,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ProjectResponse) Reset() {
	*x = ProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectResponse) ProtoMessage() {}

func (x *ProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectResponse.ProtoReflect.Descriptor instead.
func (*ProjectResponse) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProjectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *ProjectResponse) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *ProjectResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProjectResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*ProjectResponse `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	HasNext  bool               `protobuf:"varint,2,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
}

func (x *ProjectsResponse) Reset() {
	*x = ProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simoompb_v1_simoom_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectsResponse) ProtoMessage() {}

func (x *ProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simoompb_v1_simoom_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectsResponse.ProtoReflect.Descriptor instead.
func (*ProjectsResponse) Descriptor() ([]byte, []int) {
	return file_simoompb_v1_simoom_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectsResponse) GetProjects() []*ProjectResponse {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *ProjectsResponse) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

var File_simoompb_v1_simoom_proto protoreflect.FileDescriptor

var file_simoompb_v1_simoom_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69,
	0x6d, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x69, 0x6d, 0x6f,
	0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x40,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x69, 0x73, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x22, 0x26, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x67, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f,
	0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e,
	0x65, 0x78, 0x74, 0x32, 0x65, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f,
	0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd1, 0x02, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21,
	0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x20, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x21, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6e,
	0x67, 0x75, 0x75, 0x34, 0x32, 0x2f, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x73, 0x69, 0x6d, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x69, 0x6d,
	0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simoompb_v1_simoom_proto_rawDescOnce sync.Once
	file_simoompb_v1_simoom_proto_rawDescData = file_simoompb_v1_simoom_proto_rawDesc
)

func file_simoompb_v1_simoom_proto_rawDescGZIP() []byte {
	file_simoompb_v1_simoom_proto_rawDescOnce.Do(func() {
		file_simoompb_v1_simoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_simoompb_v1_simoom_proto_rawDescData)
	})
	return file_simoompb_v1_simoom_proto_rawDescData
}

var file_simoompb_v1_simoom_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_simoompb_v1_simoom_proto_goTypes = []interface{}{
	(*CheckHealthRequest)(nil),    // 0: simoompb.v1.CheckHealthRequest
	(*CheckHealthResponse)(nil),   // 1: simoompb.v1.CheckHealthResponse
	(*CreateProjectRequest)(nil),  // 2: simoompb.v1.CreateProjectRequest
	(*ListProjectsRequest)(nil),   // 3: simoompb.v1.ListProjectsRequest
	(*UpdateProjectRequest)(nil),  // 4: simoompb.v1.UpdateProjectRequest
	(*DeleteProjectRequest)(nil),  // 5: simoompb.v1.DeleteProjectRequest
	(*ProjectResponse)(nil),       // 6: simoompb.v1.ProjectResponse
	(*ProjectsResponse)(nil),      // 7: simoompb.v1.ProjectsResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_simoompb_v1_simoom_proto_depIdxs = []int32{
	8, // 0: simoompb.v1.ProjectResponse.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: simoompb.v1.ProjectResponse.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: simoompb.v1.ProjectsResponse.projects:type_name -> simoompb.v1.ProjectResponse
	0, // 3: simoompb.v1.MonitoringService.CheckHealth:input_type -> simoompb.v1.CheckHealthRequest
	2, // 4: simoompb.v1.ProjectService.CreateProject:input_type -> simoompb.v1.CreateProjectRequest
	3, // 5: simoompb.v1.ProjectService.ListProjects:input_type -> simoompb.v1.ListProjectsRequest
	4, // 6: simoompb.v1.ProjectService.UpdateProject:input_type -> simoompb.v1.UpdateProjectRequest
	5, // 7: simoompb.v1.ProjectService.DeleteProject:input_type -> simoompb.v1.DeleteProjectRequest
	1, // 8: simoompb.v1.MonitoringService.CheckHealth:output_type -> simoompb.v1.CheckHealthResponse
	6, // 9: simoompb.v1.ProjectService.CreateProject:output_type -> simoompb.v1.ProjectResponse
	7, // 10: simoompb.v1.ProjectService.ListProjects:output_type -> simoompb.v1.ProjectsResponse
	6, // 11: simoompb.v1.ProjectService.UpdateProject:output_type -> simoompb.v1.ProjectResponse
	9, // 12: simoompb.v1.ProjectService.DeleteProject:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_simoompb_v1_simoom_proto_init() }
func file_simoompb_v1_simoom_proto_init() {
	if File_simoompb_v1_simoom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simoompb_v1_simoom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckHealthRequest); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckHealthResponse); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectRequest); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectRequest); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectResponse); i {
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
		file_simoompb_v1_simoom_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectsResponse); i {
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
	file_simoompb_v1_simoom_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_simoompb_v1_simoom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_simoompb_v1_simoom_proto_goTypes,
		DependencyIndexes: file_simoompb_v1_simoom_proto_depIdxs,
		MessageInfos:      file_simoompb_v1_simoom_proto_msgTypes,
	}.Build()
	File_simoompb_v1_simoom_proto = out.File
	file_simoompb_v1_simoom_proto_rawDesc = nil
	file_simoompb_v1_simoom_proto_goTypes = nil
	file_simoompb_v1_simoom_proto_depIdxs = nil
}
