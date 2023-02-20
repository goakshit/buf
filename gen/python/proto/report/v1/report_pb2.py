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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16report/v1/report.proto\x12\treport.v1\"q\n\x06Report\x12\x36\n\x0breport_type\x18\x01 \x01(\x0e\x32\x15.report.v1.ReportTypeR\nreportType\x12\x1b\n\treport_id\x18\x02 \x01(\tR\x08reportId\x12\x12\n\x04name\x18\x03 \x01(\tR\x04name\"/\n\x10GetReportRequest\x12\x1b\n\treport_id\x18\x01 \x01(\tR\x08reportId\">\n\x11GetReportResponse\x12)\n\x06report\x18\x01 \x01(\x0b\x32\x11.report.v1.ReportR\x06report\"^\n\x10PutReportRequest\x12\x36\n\x0breport_type\x18\x01 \x01(\x0e\x32\x15.report.v1.ReportTypeR\nreportType\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\">\n\x11PutReportResponse\x12)\n\x06report\x18\x01 \x01(\x0b\x32\x11.report.v1.ReportR\x06report\"2\n\x13\x44\x65leteReportRequest\x12\x1b\n\treport_id\x18\x01 \x01(\tR\x08reportId\"\xab\x0c\n\x14MigrateReportRequest\x12\x1b\n\treport_id\x18\x01 \x01(\tR\x08reportId\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x18\n\x07\x61\x64\x64ress\x18\x03 \x01(\tR\x07\x61\x64\x64ress\x12\x12\n\x04\x63ity\x18\x04 \x01(\tR\x04\x63ity\x12\x18\n\x07\x63ountry\x18\x05 \x01(\tR\x07\x63ountry\x12\x10\n\x03pin\x18\x06 \x01(\tR\x03pin\x12\x19\n\x08pin_code\x18\x07 \x01(\tR\x07pinCode\x12\x14\n\x05phone\x18\x08 \x01(\tR\x05phone\x12 \n\x0b\x64\x65scription\x18\t \x01(\tR\x0b\x64\x65scription\x12)\n\x10long_description\x18\n \x01(\tR\x0flongDescription\x12\x1a\n\x08latitude\x18\x0b \x01(\tR\x08latitude\x12\x12\n\x04long\x18\x0c \x01(\tR\x04long\x12\x16\n\x06longer\x18\r \x01(\tR\x06longer\x12\x1a\n\x08longerer\x18\x0e \x01(\tR\x08longerer\x12\x18\n\x07smarter\x18\x0f \x01(\tR\x07smarter\x12\x1a\n\x08smarterr\x18\x10 \x01(\tR\x08smarterr\x12\x1c\n\tsmarterrr\x18\x11 \x01(\tR\tsmarterrr\x12\x16\n\x06stupid\x18\x12 \x01(\tR\x06stupid\x12\x1a\n\x08stupider\x18\x13 \x01(\tR\x08stupider\x12\x1c\n\tstupiderr\x18\x14 \x01(\tR\tstupiderr\x12\x1e\n\nstupiderrr\x18\x15 \x01(\tR\nstupiderrr\x12 \n\x0bstupiderrrr\x18\x16 \x01(\tR\x0bstupiderrrr\x12\"\n\x0cstupiderrrrr\x18\x17 \x01(\tR\x0cstupiderrrrr\x12$\n\rstupiderrrrre\x18\x18 \x01(\tR\rstupiderrrrre\x12\x12\n\x04\x66\x61st\x18\x19 \x01(\tR\x04\x66\x61st\x12\x16\n\x06\x66\x61ster\x18\x1a \x01(\tR\x06\x66\x61ster\x12\x18\n\x07\x66\x61sterr\x18\x1b \x01(\tR\x07\x66\x61sterr\x12\x1a\n\x08\x66\x61sterrr\x18\x1c \x01(\tR\x08\x66\x61sterrr\x12\x1c\n\tfasterrrr\x18\x1d \x01(\tR\tfasterrrr\x12\x12\n\x04slow\x18\x1e \x01(\tR\x04slow\x12\x16\n\x06slower\x18\x1f \x01(\tR\x06slower\x12\x18\n\x07slowerr\x18  \x01(\tR\x07slowerr\x12\x1a\n\x08slowerrr\x18! \x01(\tR\x08slowerrr\x12\x1c\n\tslowerrrr\x18\" \x01(\tR\tslowerrrr\x12\x14\n\x05test1\x18# \x01(\tR\x05test1\x12\x14\n\x05test2\x18$ \x01(\tR\x05test2\x12\x14\n\x05test3\x18% \x01(\tR\x05test3\x12\x14\n\x05test4\x18& \x01(\tR\x05test4\x12\x14\n\x05test5\x18\' \x01(\tR\x05test5\x12\x14\n\x05test6\x18( \x01(\tR\x05test6\x12\x14\n\x05test7\x18) \x01(\tR\x05test7\x12\x14\n\x05test8\x18* \x01(\tR\x05test8\x12\x14\n\x05test9\x18+ \x01(\tR\x05test9\x12\x16\n\x06test10\x18, \x01(\tR\x06test10\x12\x16\n\x06test11\x18- \x01(\tR\x06test11\x12\x16\n\x06test12\x18. \x01(\tR\x06test12\x12\x16\n\x06test13\x18/ \x01(\tR\x06test13\x12\x16\n\x06test14\x18\x30 \x01(\tR\x06test14\x12\x16\n\x06test15\x18\x31 \x01(\tR\x06test15\x12\x16\n\x06test16\x18\x32 \x01(\tR\x06test16\x12\x16\n\x06test17\x18\x33 \x01(\tR\x06test17\x12\x16\n\x06test18\x18\x34 \x01(\tR\x06test18\x12\x16\n\x06test19\x18\x35 \x01(\tR\x06test19\x12\x16\n\x06test20\x18\x36 \x01(\tR\x06test20\x12\x16\n\x06test21\x18\x37 \x01(\tR\x06test21\x12\x16\n\x06test22\x18\x38 \x01(\tR\x06test22\x12\x16\n\x06test23\x18\x39 \x01(\tR\x06test23\x12\x16\n\x06test24\x18: \x01(\tR\x06test24\x12\x16\n\x06test25\x18; \x01(\tR\x06test25\x12\x16\n\x06test26\x18< \x01(\tR\x06test26\x12\x16\n\x06test27\x18= \x01(\tR\x06test27\"\x16\n\x14\x44\x65leteReportResponse*\x8c\x01\n\nReportType\x12\x1b\n\x17REPORT_TYPE_UNSPECIFIED\x10\x00\x12\x19\n\x15REPORT_TYPE_ON_DEMAND\x10\x01\x12\x16\n\x12REPORT_TYPE_HOURLY\x10\x02\x12\x15\n\x11REPORT_TYPE_DAILY\x10\x03\x12\x17\n\x13REPORT_TYPE_MONTHLY\x10\x04\x42\x8f\x01\n\rcom.report.v1B\x0bReportProtoH\x02P\x01Z*github.com/goakshit/buf/report/v1;reportv1\xa2\x02\x03RXX\xaa\x02\tReport.V1\xca\x02\tReport\\V1\xe2\x02\x15Report\\V1\\GPBMetadata\xea\x02\nReport::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'report.v1.report_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\rcom.report.v1B\013ReportProtoH\002P\001Z*github.com/goakshit/buf/report/v1;reportv1\242\002\003RXX\252\002\tReport.V1\312\002\tReport\\V1\342\002\025Report\\V1\\GPBMetadata\352\002\nReport::V1'
  _REPORTTYPE._serialized_start=2084
  _REPORTTYPE._serialized_end=2224
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
  _MIGRATEREPORTREQUEST._serialized_end=2057
  _DELETEREPORTRESPONSE._serialized_start=2059
  _DELETEREPORTRESPONSE._serialized_end=2081
# @@protoc_insertion_point(module_scope)
