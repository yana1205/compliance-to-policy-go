// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: api/proto/v1alpha1/messages.proto

package v1alpha1

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
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[0]
	if x != nil {
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
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{0}
}

type GetSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonSchema []byte `protobuf:"bytes,1,opt,name=json_schema,json=jsonSchema,proto3" json:"json_schema,omitempty"`
}

func (x *GetSchemaResponse) Reset() {
	*x = GetSchemaResponse{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaResponse) ProtoMessage() {}

func (x *GetSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetSchemaResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetSchemaResponse) GetJsonSchema() []byte {
	if x != nil {
		return x.JsonSchema
	}
	return nil
}

type ConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ConfigureRequest) Reset() {
	*x = ConfigureRequest{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureRequest) ProtoMessage() {}

func (x *ConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureRequest.ProtoReflect.Descriptor instead.
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigureRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type ConfigureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ConfigureResponse) Reset() {
	*x = ConfigureResponse{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureResponse) ProtoMessage() {}

func (x *ConfigureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureResponse.ProtoReflect.Descriptor instead.
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigureResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RemediationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Findings []*Finding `protobuf:"bytes,1,rep,name=findings,proto3" json:"findings,omitempty"`
}

func (x *RemediationRequest) Reset() {
	*x = RemediationRequest{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemediationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemediationRequest) ProtoMessage() {}

func (x *RemediationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemediationRequest.ProtoReflect.Descriptor instead.
func (*RemediationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RemediationRequest) GetFindings() []*Finding {
	if x != nil {
		return x.Findings
	}
	return nil
}

type RemediationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemediationResponse) Reset() {
	*x = RemediationResponse{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemediationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemediationResponse) ProtoMessage() {}

func (x *RemediationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemediationResponse.ProtoReflect.Descriptor instead.
func (*RemediationResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *RemediationResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateRequest) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *PVPResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResultsResponse) Reset() {
	*x = ResultsResponse{}
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultsResponse) ProtoMessage() {}

func (x *ResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1alpha1_messages_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultsResponse.ProtoReflect.Descriptor instead.
func (*ResultsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1alpha1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *ResultsResponse) GetResult() *PVPResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ResultsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_api_proto_v1alpha1_messages_proto protoreflect.FileDescriptor

var file_api_proto_v1alpha1_messages_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x1a, 0x1f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x2a,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x66,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x2b, 0x0a, 0x13, 0x52,
	0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e,
	0x50, 0x56, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x2d, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x74,
	0x6f, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1alpha1_messages_proto_rawDescOnce sync.Once
	file_api_proto_v1alpha1_messages_proto_rawDescData = file_api_proto_v1alpha1_messages_proto_rawDesc
)

func file_api_proto_v1alpha1_messages_proto_rawDescGZIP() []byte {
	file_api_proto_v1alpha1_messages_proto_rawDescOnce.Do(func() {
		file_api_proto_v1alpha1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1alpha1_messages_proto_rawDescData)
	})
	return file_api_proto_v1alpha1_messages_proto_rawDescData
}

var file_api_proto_v1alpha1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_v1alpha1_messages_proto_goTypes = []any{
	(*Empty)(nil),               // 0: protocols.Empty
	(*GetSchemaResponse)(nil),   // 1: protocols.GetSchemaResponse
	(*ConfigureRequest)(nil),    // 2: protocols.ConfigureRequest
	(*ConfigureResponse)(nil),   // 3: protocols.ConfigureResponse
	(*RemediationRequest)(nil),  // 4: protocols.RemediationRequest
	(*RemediationResponse)(nil), // 5: protocols.RemediationResponse
	(*GenerateRequest)(nil),     // 6: protocols.GenerateRequest
	(*GenerateResponse)(nil),    // 7: protocols.GenerateResponse
	(*ResultsResponse)(nil),     // 8: protocols.ResultsResponse
	(*Finding)(nil),             // 9: protocols.Finding
	(*Policy)(nil),              // 10: protocols.Policy
	(*PVPResult)(nil),           // 11: protocols.PVPResult
}
var file_api_proto_v1alpha1_messages_proto_depIdxs = []int32{
	9,  // 0: protocols.RemediationRequest.findings:type_name -> protocols.Finding
	10, // 1: protocols.GenerateRequest.policy:type_name -> protocols.Policy
	11, // 2: protocols.ResultsResponse.result:type_name -> protocols.PVPResult
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_v1alpha1_messages_proto_init() }
func file_api_proto_v1alpha1_messages_proto_init() {
	if File_api_proto_v1alpha1_messages_proto != nil {
		return
	}
	file_api_proto_v1alpha1_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1alpha1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_v1alpha1_messages_proto_goTypes,
		DependencyIndexes: file_api_proto_v1alpha1_messages_proto_depIdxs,
		MessageInfos:      file_api_proto_v1alpha1_messages_proto_msgTypes,
	}.Build()
	File_api_proto_v1alpha1_messages_proto = out.File
	file_api_proto_v1alpha1_messages_proto_rawDesc = nil
	file_api_proto_v1alpha1_messages_proto_goTypes = nil
	file_api_proto_v1alpha1_messages_proto_depIdxs = nil
}