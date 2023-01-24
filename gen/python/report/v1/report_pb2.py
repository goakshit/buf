# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: report/v1/report.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16report/v1/report.proto\x12\treport.v1\"q\n\x06Report\x12\x36\n\x0breport_type\x18\x01 \x01(\x0e\x32\x15.report.v1.ReportTypeR\nreportType\x12\x1b\n\treport_id\x18\x02 \x01(\tR\x08reportId\x12\x12\n\x04name\x18\x03 \x01(\tR\x04name\"/\n\x10GetReportRequest\x12\x1b\n\treport_id\x18\x01 \x01(\tR\x08reportId\">\n\x11GetReportResponse\x12)\n\x06report\x18\x01 \x01(\x0b\x32\x11.report.v1.ReportR\x06report\"^\n\x10PutReportRequest\x12\x36\n\x0breport_type\x18\x01 \x01(\x0e\x32\x15.report.v1.ReportTypeR\nreportType\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\">\n\x11PutReportResponse\x12)\n\x06report\x18\x01 \x01(\x0b\x32\x11.report.v1.ReportR\x06report\"2\n\x13\x44\x65leteReportRequest\x12\x1b\n\treport_id\x18\x01 \x01(\tR\x08reportId\"\xd7\x03\n\x14MigrateReportRequest\x12\x1b\n\treport_id\x18\x01 \x01(\tR\x08reportId\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n\x07\x61\x64\x64ress\x18\x03 \x01(\tR\x07\x61\x64\x64ress\x12\x12\n\x04\x63ity\x18\x04 \x01(\tR\x04\x63ity\x12\x18\n\x07\x63ountry\x18\x05 \x01(\tR\x07\x63ountry\x12\x10\n\x03pin\x18\x06 \x01(\tR\x03pin\x12\x19\n\x08pin_code\x18\x07 \x01(\tR\x07pinCode\x12\x14\n\x05phone\x18\x08 \x01(\tR\x05phone\x12 \n\x0b\x64\x65scription\x18\t \x01(\tR\x0b\x64\x65scription\x12)\n\x10long_description\x18\n \x01(\tR\x0flongDescription\x12\x1a\n\x08latitude\x18\x0b \x01(\tR\x08latitude\x12\x12\n\x04long\x18\x0c \x01(\tR\x04long\x12\x16\n\x06longer\x18\r \x01(\tR\x06longer\x12\x1a\n\x08longerer\x18\x0e \x01(\tR\x08longerer\x12\x18\n\x07smarter\x18\x0f \x01(\tR\x07smarter\x12\x1a\n\x08smarterr\x18\x10 \x01(\tR\x08smarterr\x12\x1c\n\tsmarterrr\x18\x11 \x01(\tR\tsmarterrr\"\x16\n\x14\x44\x65leteReportResponse*\x8c\x01\n\nReportType\x12\x1b\n\x17REPORT_TYPE_UNSPECIFIED\x10\x00\x12\x19\n\x15REPORT_TYPE_ON_DEMAND\x10\x01\x12\x16\n\x12REPORT_TYPE_HOURLY\x10\x02\x12\x15\n\x11REPORT_TYPE_DAILY\x10\x03\x12\x17\n\x13REPORT_TYPE_MONTHLY\x10\x04\x32\xf6\x01\n\rReportService\x12H\n\tGetReport\x12\x1b.report.v1.GetReportRequest\x1a\x1c.report.v1.GetReportResponse\"\x00\x12H\n\tPutReport\x12\x1b.report.v1.PutReportRequest\x1a\x1c.report.v1.PutReportResponse\"\x00\x12Q\n\x0c\x44\x65leteReport\x12\x1e.report.v1.DeleteReportRequest\x1a\x1f.report.v1.DeleteReportResponse\"\x00\x42\x8f\x01\n\rcom.report.v1B\x0bReportProtoH\x02P\x01Z*github.com/goakshit/buf/report/v1;reportv1\xa2\x02\x03RXX\xaa\x02\tReport.V1\xca\x02\tReport\\V1\xe2\x02\x15Report\\V1\\GPBMetadata\xea\x02\nReport::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'report.v1.report_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\rcom.report.v1B\013ReportProtoH\002P\001Z*github.com/goakshit/buf/report/v1;reportv1\242\002\003RXX\252\002\tReport.V1\312\002\tReport\\V1\342\002\025Report\\V1\\GPBMetadata\352\002\nReport::V1'
  _REPORTTYPE._serialized_start=976
  _REPORTTYPE._serialized_end=1116
  _REPORT._serialized_start=37
  _REPORT._serialized_end=150
  _GETREPORTREQUEST._serialized_start=152
  _GETREPORTREQUEST._serialized_end=199
  _GETREPORTRESPONSE._serialized_start=201
  _GETREPORTRESPONSE._serialized_end=263
  _PUTREPORTREQUEST._serialized_start=265
  _PUTREPORTREQUEST._serialized_end=359
  _PUTREPORTRESPONSE._serialized_start=361
  _PUTREPORTRESPONSE._serialized_end=423
  _DELETEREPORTREQUEST._serialized_start=425
  _DELETEREPORTREQUEST._serialized_end=475
  _MIGRATEREPORTREQUEST._serialized_start=478
  _MIGRATEREPORTREQUEST._serialized_end=949
  _DELETEREPORTRESPONSE._serialized_start=951
  _DELETEREPORTRESPONSE._serialized_end=973
  _REPORTSERVICE._serialized_start=1119
  _REPORTSERVICE._serialized_end=1365
# @@protoc_insertion_point(module_scope)
