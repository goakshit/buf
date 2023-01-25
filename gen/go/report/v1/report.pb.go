// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: report/v1/report.proto

package reportv1

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

// ReportType represents the different types of reports.
type ReportType int32

const (
	ReportType_REPORT_TYPE_UNSPECIFIED ReportType = 0
	ReportType_REPORT_TYPE_ON_DEMAND   ReportType = 1
	ReportType_REPORT_TYPE_HOURLY      ReportType = 2
	ReportType_REPORT_TYPE_DAILY       ReportType = 3
	ReportType_REPORT_TYPE_MONTHLY     ReportType = 4
)

// Enum value maps for ReportType.
var (
	ReportType_name = map[int32]string{
		0: "REPORT_TYPE_UNSPECIFIED",
		1: "REPORT_TYPE_ON_DEMAND",
		2: "REPORT_TYPE_HOURLY",
		3: "REPORT_TYPE_DAILY",
		4: "REPORT_TYPE_MONTHLY",
	}
	ReportType_value = map[string]int32{
		"REPORT_TYPE_UNSPECIFIED": 0,
		"REPORT_TYPE_ON_DEMAND":   1,
		"REPORT_TYPE_HOURLY":      2,
		"REPORT_TYPE_DAILY":       3,
		"REPORT_TYPE_MONTHLY":     4,
	}
)

func (x ReportType) Enum() *ReportType {
	p := new(ReportType)
	*p = x
	return p
}

func (x ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_report_v1_report_proto_enumTypes[0].Descriptor()
}

func (ReportType) Type() protoreflect.EnumType {
	return &file_report_v1_report_proto_enumTypes[0]
}

func (x ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportType.Descriptor instead.
func (ReportType) EnumDescriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{0}
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportType ReportType `protobuf:"varint,1,opt,name=report_type,json=reportType,proto3,enum=report.v1.ReportType" json:"report_type,omitempty"`
	ReportId   string     `protobuf:"bytes,2,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	Name       string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetReportType() ReportType {
	if x != nil {
		return x.ReportType
	}
	return ReportType_REPORT_TYPE_UNSPECIFIED
}

func (x *Report) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *Report) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{1}
}

func (x *GetReportRequest) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

type GetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetReportResponse) Reset() {
	*x = GetReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportResponse) ProtoMessage() {}

func (x *GetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportResponse.ProtoReflect.Descriptor instead.
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type PutReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportType ReportType `protobuf:"varint,1,opt,name=report_type,json=reportType,proto3,enum=report.v1.ReportType" json:"report_type,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PutReportRequest) Reset() {
	*x = PutReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReportRequest) ProtoMessage() {}

func (x *PutReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReportRequest.ProtoReflect.Descriptor instead.
func (*PutReportRequest) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{3}
}

func (x *PutReportRequest) GetReportType() ReportType {
	if x != nil {
		return x.ReportType
	}
	return ReportType_REPORT_TYPE_UNSPECIFIED
}

func (x *PutReportRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PutReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *PutReportResponse) Reset() {
	*x = PutReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReportResponse) ProtoMessage() {}

func (x *PutReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReportResponse.ProtoReflect.Descriptor instead.
func (*PutReportResponse) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{4}
}

func (x *PutReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type DeleteReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *DeleteReportRequest) Reset() {
	*x = DeleteReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReportRequest) ProtoMessage() {}

func (x *DeleteReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReportRequest.ProtoReflect.Descriptor instead.
func (*DeleteReportRequest) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReportRequest) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

type MigrateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId        string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address         string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	City            string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country         string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Pin             string `protobuf:"bytes,6,opt,name=pin,proto3" json:"pin,omitempty"`
	PinCode         string `protobuf:"bytes,7,opt,name=pin_code,json=pinCode,proto3" json:"pin_code,omitempty"`
	Phone           string `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	Description     string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	LongDescription string `protobuf:"bytes,10,opt,name=long_description,json=longDescription,proto3" json:"long_description,omitempty"`
	Latitude        string `protobuf:"bytes,11,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Long            string `protobuf:"bytes,12,opt,name=long,proto3" json:"long,omitempty"`
	Longer          string `protobuf:"bytes,13,opt,name=longer,proto3" json:"longer,omitempty"`
	Longerer        string `protobuf:"bytes,14,opt,name=longerer,proto3" json:"longerer,omitempty"`
	Smarter         string `protobuf:"bytes,15,opt,name=smarter,proto3" json:"smarter,omitempty"`
	Smarterr        string `protobuf:"bytes,16,opt,name=smarterr,proto3" json:"smarterr,omitempty"`
	Smarterrr       string `protobuf:"bytes,17,opt,name=smarterrr,proto3" json:"smarterrr,omitempty"`
	Stupid          string `protobuf:"bytes,18,opt,name=stupid,proto3" json:"stupid,omitempty"`
	Stupider        string `protobuf:"bytes,19,opt,name=stupider,proto3" json:"stupider,omitempty"`
	Stupiderr       string `protobuf:"bytes,20,opt,name=stupiderr,proto3" json:"stupiderr,omitempty"`
	Stupiderrr      string `protobuf:"bytes,21,opt,name=stupiderrr,proto3" json:"stupiderrr,omitempty"`
	Stupiderrrr     string `protobuf:"bytes,22,opt,name=stupiderrrr,proto3" json:"stupiderrrr,omitempty"`
	Stupiderrrrr    string `protobuf:"bytes,23,opt,name=stupiderrrrr,proto3" json:"stupiderrrrr,omitempty"`
	Stupiderrrrre   string `protobuf:"bytes,24,opt,name=stupiderrrrre,proto3" json:"stupiderrrrre,omitempty"`
	Fast            string `protobuf:"bytes,25,opt,name=fast,proto3" json:"fast,omitempty"`
	Faster          string `protobuf:"bytes,26,opt,name=faster,proto3" json:"faster,omitempty"`
	Fasterr         string `protobuf:"bytes,27,opt,name=fasterr,proto3" json:"fasterr,omitempty"`
	Fasterrr        string `protobuf:"bytes,28,opt,name=fasterrr,proto3" json:"fasterrr,omitempty"`
	Fasterrrr       string `protobuf:"bytes,29,opt,name=fasterrrr,proto3" json:"fasterrrr,omitempty"`
	Slow            string `protobuf:"bytes,30,opt,name=slow,proto3" json:"slow,omitempty"`
	Slower          string `protobuf:"bytes,31,opt,name=slower,proto3" json:"slower,omitempty"`
	Slowerr         string `protobuf:"bytes,32,opt,name=slowerr,proto3" json:"slowerr,omitempty"`
	Slowerrr        string `protobuf:"bytes,33,opt,name=slowerrr,proto3" json:"slowerrr,omitempty"`
	Slowerrrr       string `protobuf:"bytes,34,opt,name=slowerrrr,proto3" json:"slowerrrr,omitempty"`
	Test1           string `protobuf:"bytes,35,opt,name=test1,proto3" json:"test1,omitempty"`
	Test2           string `protobuf:"bytes,36,opt,name=test2,proto3" json:"test2,omitempty"`
	Test3           string `protobuf:"bytes,37,opt,name=test3,proto3" json:"test3,omitempty"`
	Test4           string `protobuf:"bytes,38,opt,name=test4,proto3" json:"test4,omitempty"`
	Test5           string `protobuf:"bytes,39,opt,name=test5,proto3" json:"test5,omitempty"`
	Test6           string `protobuf:"bytes,40,opt,name=test6,proto3" json:"test6,omitempty"`
	Test7           string `protobuf:"bytes,41,opt,name=test7,proto3" json:"test7,omitempty"`
	Test8           string `protobuf:"bytes,42,opt,name=test8,proto3" json:"test8,omitempty"`
	Test9           string `protobuf:"bytes,43,opt,name=test9,proto3" json:"test9,omitempty"`
}

func (x *MigrateReportRequest) Reset() {
	*x = MigrateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateReportRequest) ProtoMessage() {}

func (x *MigrateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateReportRequest.ProtoReflect.Descriptor instead.
func (*MigrateReportRequest) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{6}
}

func (x *MigrateReportRequest) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *MigrateReportRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MigrateReportRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MigrateReportRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MigrateReportRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MigrateReportRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

func (x *MigrateReportRequest) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *MigrateReportRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *MigrateReportRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MigrateReportRequest) GetLongDescription() string {
	if x != nil {
		return x.LongDescription
	}
	return ""
}

func (x *MigrateReportRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *MigrateReportRequest) GetLong() string {
	if x != nil {
		return x.Long
	}
	return ""
}

func (x *MigrateReportRequest) GetLonger() string {
	if x != nil {
		return x.Longer
	}
	return ""
}

func (x *MigrateReportRequest) GetLongerer() string {
	if x != nil {
		return x.Longerer
	}
	return ""
}

func (x *MigrateReportRequest) GetSmarter() string {
	if x != nil {
		return x.Smarter
	}
	return ""
}

func (x *MigrateReportRequest) GetSmarterr() string {
	if x != nil {
		return x.Smarterr
	}
	return ""
}

func (x *MigrateReportRequest) GetSmarterrr() string {
	if x != nil {
		return x.Smarterrr
	}
	return ""
}

func (x *MigrateReportRequest) GetStupid() string {
	if x != nil {
		return x.Stupid
	}
	return ""
}

func (x *MigrateReportRequest) GetStupider() string {
	if x != nil {
		return x.Stupider
	}
	return ""
}

func (x *MigrateReportRequest) GetStupiderr() string {
	if x != nil {
		return x.Stupiderr
	}
	return ""
}

func (x *MigrateReportRequest) GetStupiderrr() string {
	if x != nil {
		return x.Stupiderrr
	}
	return ""
}

func (x *MigrateReportRequest) GetStupiderrrr() string {
	if x != nil {
		return x.Stupiderrrr
	}
	return ""
}

func (x *MigrateReportRequest) GetStupiderrrrr() string {
	if x != nil {
		return x.Stupiderrrrr
	}
	return ""
}

func (x *MigrateReportRequest) GetStupiderrrrre() string {
	if x != nil {
		return x.Stupiderrrrre
	}
	return ""
}

func (x *MigrateReportRequest) GetFast() string {
	if x != nil {
		return x.Fast
	}
	return ""
}

func (x *MigrateReportRequest) GetFaster() string {
	if x != nil {
		return x.Faster
	}
	return ""
}

func (x *MigrateReportRequest) GetFasterr() string {
	if x != nil {
		return x.Fasterr
	}
	return ""
}

func (x *MigrateReportRequest) GetFasterrr() string {
	if x != nil {
		return x.Fasterrr
	}
	return ""
}

func (x *MigrateReportRequest) GetFasterrrr() string {
	if x != nil {
		return x.Fasterrrr
	}
	return ""
}

func (x *MigrateReportRequest) GetSlow() string {
	if x != nil {
		return x.Slow
	}
	return ""
}

func (x *MigrateReportRequest) GetSlower() string {
	if x != nil {
		return x.Slower
	}
	return ""
}

func (x *MigrateReportRequest) GetSlowerr() string {
	if x != nil {
		return x.Slowerr
	}
	return ""
}

func (x *MigrateReportRequest) GetSlowerrr() string {
	if x != nil {
		return x.Slowerrr
	}
	return ""
}

func (x *MigrateReportRequest) GetSlowerrrr() string {
	if x != nil {
		return x.Slowerrrr
	}
	return ""
}

func (x *MigrateReportRequest) GetTest1() string {
	if x != nil {
		return x.Test1
	}
	return ""
}

func (x *MigrateReportRequest) GetTest2() string {
	if x != nil {
		return x.Test2
	}
	return ""
}

func (x *MigrateReportRequest) GetTest3() string {
	if x != nil {
		return x.Test3
	}
	return ""
}

func (x *MigrateReportRequest) GetTest4() string {
	if x != nil {
		return x.Test4
	}
	return ""
}

func (x *MigrateReportRequest) GetTest5() string {
	if x != nil {
		return x.Test5
	}
	return ""
}

func (x *MigrateReportRequest) GetTest6() string {
	if x != nil {
		return x.Test6
	}
	return ""
}

func (x *MigrateReportRequest) GetTest7() string {
	if x != nil {
		return x.Test7
	}
	return ""
}

func (x *MigrateReportRequest) GetTest8() string {
	if x != nil {
		return x.Test8
	}
	return ""
}

func (x *MigrateReportRequest) GetTest9() string {
	if x != nil {
		return x.Test9
	}
	return ""
}

type DeleteReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteReportResponse) Reset() {
	*x = DeleteReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_v1_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReportResponse) ProtoMessage() {}

func (x *DeleteReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_v1_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReportResponse.ProtoReflect.Descriptor instead.
func (*DeleteReportResponse) Descriptor() ([]byte, []int) {
	return file_report_v1_report_proto_rawDescGZIP(), []int{7}
}

var File_report_v1_report_proto protoreflect.FileDescriptor

var file_report_v1_report_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x22, 0x71, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x36, 0x0a,
	0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x5e, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0xfb, 0x08, 0x0a, 0x14,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f,
	0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e,
	0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x6f, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x65, 0x72, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x65, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x65, 0x72, 0x72, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x65, 0x72, 0x72, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75,
	0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x70, 0x69,
	0x64, 0x65, 0x72, 0x72, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x75,
	0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x75, 0x70, 0x69,
	0x64, 0x65, 0x72, 0x72, 0x72, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x75,
	0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x72, 0x72, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x72, 0x72, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x72, 0x72, 0x72, 0x65, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x75, 0x70, 0x69, 0x64, 0x65, 0x72, 0x72, 0x72,
	0x72, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x73, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x61, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x72, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72,
	0x72, 0x72, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x72, 0x72, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x77, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x72, 0x72, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x72,
	0x72, 0x72, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x72, 0x72, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x31, 0x18, 0x23, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73,
	0x74, 0x32, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x32, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x33, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x65, 0x73, 0x74, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x34, 0x18, 0x26,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x65, 0x73, 0x74, 0x35, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74,
	0x35, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x36, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x36, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x37,
	0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x37, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x65, 0x73, 0x74, 0x38, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65,
	0x73, 0x74, 0x38, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x39, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x39, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2a, 0x8c, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x4c, 0x59, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x04,
	0x42, 0x8f, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48,
	0x02, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x61, 0x6b, 0x73, 0x68, 0x69, 0x74, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_v1_report_proto_rawDescOnce sync.Once
	file_report_v1_report_proto_rawDescData = file_report_v1_report_proto_rawDesc
)

func file_report_v1_report_proto_rawDescGZIP() []byte {
	file_report_v1_report_proto_rawDescOnce.Do(func() {
		file_report_v1_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_v1_report_proto_rawDescData)
	})
	return file_report_v1_report_proto_rawDescData
}

var file_report_v1_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_report_v1_report_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_report_v1_report_proto_goTypes = []interface{}{
	(ReportType)(0),              // 0: report.v1.ReportType
	(*Report)(nil),               // 1: report.v1.Report
	(*GetReportRequest)(nil),     // 2: report.v1.GetReportRequest
	(*GetReportResponse)(nil),    // 3: report.v1.GetReportResponse
	(*PutReportRequest)(nil),     // 4: report.v1.PutReportRequest
	(*PutReportResponse)(nil),    // 5: report.v1.PutReportResponse
	(*DeleteReportRequest)(nil),  // 6: report.v1.DeleteReportRequest
	(*MigrateReportRequest)(nil), // 7: report.v1.MigrateReportRequest
	(*DeleteReportResponse)(nil), // 8: report.v1.DeleteReportResponse
}
var file_report_v1_report_proto_depIdxs = []int32{
	0, // 0: report.v1.Report.report_type:type_name -> report.v1.ReportType
	1, // 1: report.v1.GetReportResponse.report:type_name -> report.v1.Report
	0, // 2: report.v1.PutReportRequest.report_type:type_name -> report.v1.ReportType
	1, // 3: report.v1.PutReportResponse.report:type_name -> report.v1.Report
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_report_v1_report_proto_init() }
func file_report_v1_report_proto_init() {
	if File_report_v1_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_v1_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_report_v1_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportRequest); i {
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
		file_report_v1_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportResponse); i {
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
		file_report_v1_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReportRequest); i {
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
		file_report_v1_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReportResponse); i {
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
		file_report_v1_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReportRequest); i {
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
		file_report_v1_report_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateReportRequest); i {
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
		file_report_v1_report_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReportResponse); i {
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
			RawDescriptor: file_report_v1_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_report_v1_report_proto_goTypes,
		DependencyIndexes: file_report_v1_report_proto_depIdxs,
		EnumInfos:         file_report_v1_report_proto_enumTypes,
		MessageInfos:      file_report_v1_report_proto_msgTypes,
	}.Build()
	File_report_v1_report_proto = out.File
	file_report_v1_report_proto_rawDesc = nil
	file_report_v1_report_proto_goTypes = nil
	file_report_v1_report_proto_depIdxs = nil
}
