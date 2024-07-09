// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: api/v1/api.proto

package apiv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
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

type CreateSeasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year   int32  `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Series string `protobuf:"bytes,2,opt,name=series,proto3" json:"series,omitempty"`
}

func (x *CreateSeasonRequest) Reset() {
	*x = CreateSeasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeasonRequest) ProtoMessage() {}

func (x *CreateSeasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeasonRequest.ProtoReflect.Descriptor instead.
func (*CreateSeasonRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSeasonRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CreateSeasonRequest) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

type CreateSeasonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId int32 `protobuf:"varint,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
}

func (x *CreateSeasonResponse) Reset() {
	*x = CreateSeasonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeasonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeasonResponse) ProtoMessage() {}

func (x *CreateSeasonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeasonResponse.ProtoReflect.Descriptor instead.
func (*CreateSeasonResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSeasonResponse) GetSeasonId() int32 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type GetSeasonByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId int32 `protobuf:"varint,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
}

func (x *GetSeasonByIdRequest) Reset() {
	*x = GetSeasonByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeasonByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeasonByIdRequest) ProtoMessage() {}

func (x *GetSeasonByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeasonByIdRequest.ProtoReflect.Descriptor instead.
func (*GetSeasonByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetSeasonByIdRequest) GetSeasonId() int32 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type GetSeasonByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Season *Season `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
}

func (x *GetSeasonByIdResponse) Reset() {
	*x = GetSeasonByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeasonByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeasonByIdResponse) ProtoMessage() {}

func (x *GetSeasonByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeasonByIdResponse.ProtoReflect.Descriptor instead.
func (*GetSeasonByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetSeasonByIdResponse) GetSeason() *Season {
	if x != nil {
		return x.Season
	}
	return nil
}

type GetAllSeasonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllSeasonsRequest) Reset() {
	*x = GetAllSeasonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSeasonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSeasonsRequest) ProtoMessage() {}

func (x *GetAllSeasonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSeasonsRequest.ProtoReflect.Descriptor instead.
func (*GetAllSeasonsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{4}
}

type GetAllSeasonsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seasons []*Season `protobuf:"bytes,1,rep,name=seasons,proto3" json:"seasons,omitempty"`
}

func (x *GetAllSeasonsResponse) Reset() {
	*x = GetAllSeasonsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSeasonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSeasonsResponse) ProtoMessage() {}

func (x *GetAllSeasonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSeasonsResponse.ProtoReflect.Descriptor instead.
func (*GetAllSeasonsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllSeasonsResponse) GetSeasons() []*Season {
	if x != nil {
		return x.Seasons
	}
	return nil
}

type CreateDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string     `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string     `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PlaceOfBirth string     `protobuf:"bytes,3,opt,name=place_of_birth,json=placeOfBirth,proto3" json:"place_of_birth,omitempty"`
	DateOfBirth  *date.Date `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
}

func (x *CreateDriverRequest) Reset() {
	*x = CreateDriverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDriverRequest) ProtoMessage() {}

func (x *CreateDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDriverRequest.ProtoReflect.Descriptor instead.
func (*CreateDriverRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDriverRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateDriverRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateDriverRequest) GetPlaceOfBirth() string {
	if x != nil {
		return x.PlaceOfBirth
	}
	return ""
}

func (x *CreateDriverRequest) GetDateOfBirth() *date.Date {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

type CreateDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *CreateDriverResponse) Reset() {
	*x = CreateDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDriverResponse) ProtoMessage() {}

func (x *CreateDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDriverResponse.ProtoReflect.Descriptor instead.
func (*CreateDriverResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *CreateDriverResponse) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type GetDriverByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *GetDriverByIdRequest) Reset() {
	*x = GetDriverByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriverByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverByIdRequest) ProtoMessage() {}

func (x *GetDriverByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriverByIdRequest.ProtoReflect.Descriptor instead.
func (*GetDriverByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetDriverByIdRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type GetDriverByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver *Driver `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *GetDriverByIdResponse) Reset() {
	*x = GetDriverByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriverByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverByIdResponse) ProtoMessage() {}

func (x *GetDriverByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriverByIdResponse.ProtoReflect.Descriptor instead.
func (*GetDriverByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetDriverByIdResponse) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Base string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *CreateTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTeamRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

type CreateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int32 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *CreateTeamResponse) Reset() {
	*x = CreateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamResponse) ProtoMessage() {}

func (x *CreateTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_api_proto_rawDescGZIP(), []int{11}
}

func (x *CreateTeamResponse) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

var File_api_v1_api_proto protoreflect.FileDescriptor

var file_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x56, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x1a, 0x06, 0x18, 0xb4,
	0x10, 0x20, 0xec, 0x0e, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x3c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a,
	0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x16,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x07, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0xe3, 0x03, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x3d, 0x0a, 0x0d, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x3a, 0x92, 0x02, 0xba, 0x48, 0x8e,
	0x02, 0x1a, 0x58, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x64, 0x6f, 0x62, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x12, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x62, 0x69, 0x72, 0x74, 0x68, 0x20, 0x6d,
	0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x79, 0x65, 0x61, 0x72, 0x1a, 0x1c, 0x68,
	0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x2e, 0x79, 0x65, 0x61, 0x72, 0x29, 0x1a, 0x5b, 0x0a, 0x17, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x62,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x21, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x6f, 0x66, 0x20, 0x62, 0x69, 0x72, 0x74, 0x68, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68,
	0x61, 0x76, 0x65, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x1a, 0x1d, 0x68, 0x61, 0x73, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x2e, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x29, 0x1a, 0x55, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x62, 0x5f, 0x64, 0x61,
	0x79, 0x12, 0x1f, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x64,
	0x61, 0x79, 0x1a, 0x1b, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x2e, 0x64, 0x61, 0x79, 0x29, 0x22,
	0x33, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x32,
	0xad, 0x04, 0x0a, 0x12, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x6d, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x61, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x68,
	0x6f, 0x6c, 0x74, 0x7a, 0x73, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c,
	0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_api_proto_rawDescOnce sync.Once
	file_api_v1_api_proto_rawDescData = file_api_v1_api_proto_rawDesc
)

func file_api_v1_api_proto_rawDescGZIP() []byte {
	file_api_v1_api_proto_rawDescOnce.Do(func() {
		file_api_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_api_proto_rawDescData)
	})
	return file_api_v1_api_proto_rawDescData
}

var file_api_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_v1_api_proto_goTypes = []interface{}{
	(*CreateSeasonRequest)(nil),   // 0: api.v1.CreateSeasonRequest
	(*CreateSeasonResponse)(nil),  // 1: api.v1.CreateSeasonResponse
	(*GetSeasonByIdRequest)(nil),  // 2: api.v1.GetSeasonByIdRequest
	(*GetSeasonByIdResponse)(nil), // 3: api.v1.GetSeasonByIdResponse
	(*GetAllSeasonsRequest)(nil),  // 4: api.v1.GetAllSeasonsRequest
	(*GetAllSeasonsResponse)(nil), // 5: api.v1.GetAllSeasonsResponse
	(*CreateDriverRequest)(nil),   // 6: api.v1.CreateDriverRequest
	(*CreateDriverResponse)(nil),  // 7: api.v1.CreateDriverResponse
	(*GetDriverByIdRequest)(nil),  // 8: api.v1.GetDriverByIdRequest
	(*GetDriverByIdResponse)(nil), // 9: api.v1.GetDriverByIdResponse
	(*CreateTeamRequest)(nil),     // 10: api.v1.CreateTeamRequest
	(*CreateTeamResponse)(nil),    // 11: api.v1.CreateTeamResponse
	(*Season)(nil),                // 12: api.v1.Season
	(*date.Date)(nil),             // 13: google.type.Date
	(*Driver)(nil),                // 14: api.v1.Driver
}
var file_api_v1_api_proto_depIdxs = []int32{
	12, // 0: api.v1.GetSeasonByIdResponse.season:type_name -> api.v1.Season
	12, // 1: api.v1.GetAllSeasonsResponse.seasons:type_name -> api.v1.Season
	13, // 2: api.v1.CreateDriverRequest.date_of_birth:type_name -> google.type.Date
	14, // 3: api.v1.GetDriverByIdResponse.driver:type_name -> api.v1.Driver
	0,  // 4: api.v1.FormulaDataService.CreateSeason:input_type -> api.v1.CreateSeasonRequest
	2,  // 5: api.v1.FormulaDataService.GetSeasonById:input_type -> api.v1.GetSeasonByIdRequest
	4,  // 6: api.v1.FormulaDataService.GetAllSeasons:input_type -> api.v1.GetAllSeasonsRequest
	6,  // 7: api.v1.FormulaDataService.CreateDriver:input_type -> api.v1.CreateDriverRequest
	8,  // 8: api.v1.FormulaDataService.GetDriverById:input_type -> api.v1.GetDriverByIdRequest
	10, // 9: api.v1.FormulaDataService.CreateTeam:input_type -> api.v1.CreateTeamRequest
	1,  // 10: api.v1.FormulaDataService.CreateSeason:output_type -> api.v1.CreateSeasonResponse
	3,  // 11: api.v1.FormulaDataService.GetSeasonById:output_type -> api.v1.GetSeasonByIdResponse
	5,  // 12: api.v1.FormulaDataService.GetAllSeasons:output_type -> api.v1.GetAllSeasonsResponse
	7,  // 13: api.v1.FormulaDataService.CreateDriver:output_type -> api.v1.CreateDriverResponse
	9,  // 14: api.v1.FormulaDataService.GetDriverById:output_type -> api.v1.GetDriverByIdResponse
	11, // 15: api.v1.FormulaDataService.CreateTeam:output_type -> api.v1.CreateTeamResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_api_proto_init() }
func file_api_v1_api_proto_init() {
	if File_api_v1_api_proto != nil {
		return
	}
	file_api_v1_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeasonRequest); i {
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
		file_api_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeasonResponse); i {
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
		file_api_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeasonByIdRequest); i {
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
		file_api_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeasonByIdResponse); i {
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
		file_api_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSeasonsRequest); i {
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
		file_api_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSeasonsResponse); i {
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
		file_api_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDriverRequest); i {
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
		file_api_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDriverResponse); i {
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
		file_api_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriverByIdRequest); i {
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
		file_api_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriverByIdResponse); i {
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
		file_api_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamRequest); i {
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
		file_api_v1_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamResponse); i {
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
			RawDescriptor: file_api_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_api_proto_goTypes,
		DependencyIndexes: file_api_v1_api_proto_depIdxs,
		MessageInfos:      file_api_v1_api_proto_msgTypes,
	}.Build()
	File_api_v1_api_proto = out.File
	file_api_v1_api_proto_rawDesc = nil
	file_api_v1_api_proto_goTypes = nil
	file_api_v1_api_proto_depIdxs = nil
}
