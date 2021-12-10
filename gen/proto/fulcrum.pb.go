// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: fulcrum.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{0}
}

type MergeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line string `protobuf:"bytes,1,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *MergeRequest) Reset() {
	*x = MergeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeRequest) ProtoMessage() {}

func (x *MergeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeRequest.ProtoReflect.Descriptor instead.
func (*MergeRequest) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{1}
}

func (x *MergeRequest) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

type MergeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MergeResponse) Reset() {
	*x = MergeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeResponse) ProtoMessage() {}

func (x *MergeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeResponse.ProtoReflect.Descriptor instead.
func (*MergeResponse) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{2}
}

func (x *MergeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AddCityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
	Soldados      int32  `protobuf:"varint,3,opt,name=soldados,proto3" json:"soldados,omitempty"`
}

func (x *AddCityRequest) Reset() {
	*x = AddCityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCityRequest) ProtoMessage() {}

func (x *AddCityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCityRequest.ProtoReflect.Descriptor instead.
func (*AddCityRequest) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{3}
}

func (x *AddCityRequest) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *AddCityRequest) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

func (x *AddCityRequest) GetSoldados() int32 {
	if x != nil {
		return x.Soldados
	}
	return 0
}

type VectorClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X  int32  `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y  int32  `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z  int32  `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
	Ip string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *VectorClock) Reset() {
	*x = VectorClock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClock) ProtoMessage() {}

func (x *VectorClock) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClock.ProtoReflect.Descriptor instead.
func (*VectorClock) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{4}
}

func (x *VectorClock) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *VectorClock) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *VectorClock) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *VectorClock) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type UpdateNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
	NuevoNombre   string `protobuf:"bytes,3,opt,name=nuevo_nombre,json=nuevoNombre,proto3" json:"nuevo_nombre,omitempty"`
}

func (x *UpdateNameRequest) Reset() {
	*x = UpdateNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNameRequest) ProtoMessage() {}

func (x *UpdateNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNameRequest.ProtoReflect.Descriptor instead.
func (*UpdateNameRequest) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNameRequest) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *UpdateNameRequest) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

func (x *UpdateNameRequest) GetNuevoNombre() string {
	if x != nil {
		return x.NuevoNombre
	}
	return ""
}

type UpdateNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *UpdateNameResponse) Reset() {
	*x = UpdateNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNameResponse) ProtoMessage() {}

func (x *UpdateNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNameResponse.ProtoReflect.Descriptor instead.
func (*UpdateNameResponse) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNameResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *UpdateNameResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *UpdateNameResponse) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type UpdateNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
	NuevoNumero   int32  `protobuf:"varint,3,opt,name=nuevo_numero,json=nuevoNumero,proto3" json:"nuevo_numero,omitempty"`
}

func (x *UpdateNumberRequest) Reset() {
	*x = UpdateNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNumberRequest) ProtoMessage() {}

func (x *UpdateNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNumberRequest.ProtoReflect.Descriptor instead.
func (*UpdateNumberRequest) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateNumberRequest) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *UpdateNumberRequest) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

func (x *UpdateNumberRequest) GetNuevoNumero() int32 {
	if x != nil {
		return x.NuevoNumero
	}
	return 0
}

type UpdateNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *UpdateNumberResponse) Reset() {
	*x = UpdateNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNumberResponse) ProtoMessage() {}

func (x *UpdateNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNumberResponse.ProtoReflect.Descriptor instead.
func (*UpdateNumberResponse) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateNumberResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *UpdateNumberResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *UpdateNumberResponse) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type DeleteCityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
}

func (x *DeleteCityRequest) Reset() {
	*x = DeleteCityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCityRequest) ProtoMessage() {}

func (x *DeleteCityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCityRequest.ProtoReflect.Descriptor instead.
func (*DeleteCityRequest) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCityRequest) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *DeleteCityRequest) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

type DeleteCityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *DeleteCityResponse) Reset() {
	*x = DeleteCityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCityResponse) ProtoMessage() {}

func (x *DeleteCityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCityResponse.ProtoReflect.Descriptor instead.
func (*DeleteCityResponse) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCityResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *DeleteCityResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *DeleteCityResponse) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_fulcrum_proto protoreflect.FileDescriptor

var file_fulcrum_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x0a, 0x0c,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0x21, 0x0a, 0x0d, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x78, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x75, 0x64, 0x61,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x22, 0x47, 0x0a,
	0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69,
	0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x65, 0x76,
	0x6f, 0x5f, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x75, 0x65, 0x76, 0x6f, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x22, 0x84, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x6f, 0x22, 0x40, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x7a, 0x22, 0x5f, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x75, 0x64, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43,
	0x69, 0x75, 0x64, 0x61, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x7a, 0x32, 0xb4, 0x03, 0x0a, 0x07, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x12, 0x32, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x38, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x5b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x0b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fulcrum_proto_rawDescOnce sync.Once
	file_fulcrum_proto_rawDescData = file_fulcrum_proto_rawDesc
)

func file_fulcrum_proto_rawDescGZIP() []byte {
	file_fulcrum_proto_rawDescOnce.Do(func() {
		file_fulcrum_proto_rawDescData = protoimpl.X.CompressGZIP(file_fulcrum_proto_rawDescData)
	})
	return file_fulcrum_proto_rawDescData
}

var file_fulcrum_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_fulcrum_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: grpc.Empty
	(*MergeRequest)(nil),              // 1: grpc.MergeRequest
	(*MergeResponse)(nil),             // 2: grpc.MergeResponse
	(*AddCityRequest)(nil),            // 3: grpc.AddCityRequest
	(*VectorClock)(nil),               // 4: grpc.VectorClock
	(*UpdateNameRequest)(nil),         // 5: grpc.UpdateNameRequest
	(*UpdateNameResponse)(nil),        // 6: grpc.UpdateNameResponse
	(*UpdateNumberRequest)(nil),       // 7: grpc.UpdateNumberRequest
	(*UpdateNumberResponse)(nil),      // 8: grpc.UpdateNumberResponse
	(*DeleteCityRequest)(nil),         // 9: grpc.DeleteCityRequest
	(*DeleteCityResponse)(nil),        // 10: grpc.DeleteCityResponse
	(*GetNumberRebeldesRequest)(nil),  // 11: grpc.GetNumberRebeldesRequest
	(*GetNumberRebeldesResponse)(nil), // 12: grpc.GetNumberRebeldesResponse
}
var file_fulcrum_proto_depIdxs = []int32{
	3,  // 0: grpc.Fulcrum.AddCity:input_type -> grpc.AddCityRequest
	9,  // 1: grpc.Fulcrum.DeleteCity:input_type -> grpc.DeleteCityRequest
	5,  // 2: grpc.Fulcrum.UpdateName:input_type -> grpc.UpdateNameRequest
	7,  // 3: grpc.Fulcrum.UpdateNumber:input_type -> grpc.UpdateNumberRequest
	11, // 4: grpc.Fulcrum.GetNumberRebeldesFulcrum:input_type -> grpc.GetNumberRebeldesRequest
	4,  // 5: grpc.Fulcrum.VectorClockMerge:input_type -> grpc.VectorClock
	1,  // 6: grpc.Fulcrum.Merge:input_type -> grpc.MergeRequest
	4,  // 7: grpc.Fulcrum.AddCity:output_type -> grpc.VectorClock
	4,  // 8: grpc.Fulcrum.DeleteCity:output_type -> grpc.VectorClock
	4,  // 9: grpc.Fulcrum.UpdateName:output_type -> grpc.VectorClock
	4,  // 10: grpc.Fulcrum.UpdateNumber:output_type -> grpc.VectorClock
	12, // 11: grpc.Fulcrum.GetNumberRebeldesFulcrum:output_type -> grpc.GetNumberRebeldesResponse
	0,  // 12: grpc.Fulcrum.VectorClockMerge:output_type -> grpc.Empty
	2,  // 13: grpc.Fulcrum.Merge:output_type -> grpc.MergeResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_fulcrum_proto_init() }
func file_fulcrum_proto_init() {
	if File_fulcrum_proto != nil {
		return
	}
	file_broker_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fulcrum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_fulcrum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeRequest); i {
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
		file_fulcrum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeResponse); i {
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
		file_fulcrum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCityRequest); i {
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
		file_fulcrum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectorClock); i {
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
		file_fulcrum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNameRequest); i {
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
		file_fulcrum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNameResponse); i {
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
		file_fulcrum_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNumberRequest); i {
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
		file_fulcrum_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNumberResponse); i {
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
		file_fulcrum_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCityRequest); i {
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
		file_fulcrum_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCityResponse); i {
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
			RawDescriptor: file_fulcrum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fulcrum_proto_goTypes,
		DependencyIndexes: file_fulcrum_proto_depIdxs,
		MessageInfos:      file_fulcrum_proto_msgTypes,
	}.Build()
	File_fulcrum_proto = out.File
	file_fulcrum_proto_rawDesc = nil
	file_fulcrum_proto_goTypes = nil
	file_fulcrum_proto_depIdxs = nil
}
