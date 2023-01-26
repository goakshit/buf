# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: report/v1/report.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("report/v1/report.proto", :syntax => :proto3) do
    add_message "report.v1.Report" do
      optional :report_type, :enum, 1, "report.v1.ReportType", json_name: "reportType"
      optional :report_id, :string, 2, json_name: "reportId"
      optional :name, :string, 3, json_name: "name"
    end
    add_message "report.v1.GetReportRequest" do
      optional :report_id, :string, 1, json_name: "reportId"
    end
    add_message "report.v1.GetReportResponse" do
      optional :report, :message, 1, "report.v1.Report", json_name: "report"
    end
    add_message "report.v1.PutReportRequest" do
      optional :report_type, :enum, 1, "report.v1.ReportType", json_name: "reportType"
      optional :name, :string, 2, json_name: "name"
    end
    add_message "report.v1.PutReportResponse" do
      optional :report, :message, 1, "report.v1.Report", json_name: "report"
    end
    add_message "report.v1.DeleteReportRequest" do
      optional :report_id, :string, 1, json_name: "reportId"
    end
    add_message "report.v1.MigrateReportRequest" do
      optional :report_id, :string, 1, json_name: "reportId"
      optional :name, :string, 2, json_name: "name"
      optional :address, :string, 3, json_name: "address"
      optional :city, :string, 4, json_name: "city"
      optional :country, :string, 5, json_name: "country"
      optional :pin, :string, 6, json_name: "pin"
      optional :pin_code, :string, 7, json_name: "pinCode"
      optional :phone, :string, 8, json_name: "phone"
      optional :description, :string, 9, json_name: "description"
      optional :long_description, :string, 10, json_name: "longDescription"
      optional :latitude, :string, 11, json_name: "latitude"
      optional :long, :string, 12, json_name: "long"
      optional :longer, :string, 13, json_name: "longer"
      optional :longerer, :string, 14, json_name: "longerer"
      optional :smarter, :string, 15, json_name: "smarter"
      optional :smarterr, :string, 16, json_name: "smarterr"
      optional :smarterrr, :string, 17, json_name: "smarterrr"
      optional :stupid, :string, 18, json_name: "stupid"
      optional :stupider, :string, 19, json_name: "stupider"
      optional :stupiderr, :string, 20, json_name: "stupiderr"
      optional :stupiderrr, :string, 21, json_name: "stupiderrr"
      optional :stupiderrrr, :string, 22, json_name: "stupiderrrr"
      optional :stupiderrrrr, :string, 23, json_name: "stupiderrrrr"
      optional :stupiderrrrre, :string, 24, json_name: "stupiderrrrre"
      optional :fast, :string, 25, json_name: "fast"
      optional :faster, :string, 26, json_name: "faster"
      optional :fasterr, :string, 27, json_name: "fasterr"
      optional :fasterrr, :string, 28, json_name: "fasterrr"
      optional :fasterrrr, :string, 29, json_name: "fasterrrr"
      optional :slow, :string, 30, json_name: "slow"
      optional :slower, :string, 31, json_name: "slower"
      optional :slowerr, :string, 32, json_name: "slowerr"
      optional :slowerrr, :string, 33, json_name: "slowerrr"
      optional :slowerrrr, :string, 34, json_name: "slowerrrr"
      optional :test1, :string, 35, json_name: "test1"
      optional :test2, :string, 36, json_name: "test2"
      optional :test3, :string, 37, json_name: "test3"
      optional :test4, :string, 38, json_name: "test4"
      optional :test5, :string, 39, json_name: "test5"
      optional :test6, :string, 40, json_name: "test6"
      optional :test7, :string, 41, json_name: "test7"
      optional :test8, :string, 42, json_name: "test8"
      optional :test9, :string, 43, json_name: "test9"
      optional :test10, :string, 44, json_name: "test10"
      optional :test11, :string, 45, json_name: "test11"
      optional :test12, :string, 46, json_name: "test12"
      optional :test13, :string, 47, json_name: "test13"
      optional :test14, :string, 48, json_name: "test14"
      optional :test15, :string, 49, json_name: "test15"
      optional :test16, :string, 50, json_name: "test16"
      optional :test17, :string, 51, json_name: "test17"
      optional :test18, :string, 52, json_name: "test18"
    end
    add_message "report.v1.DeleteReportResponse" do
    end
    add_enum "report.v1.ReportType" do
      value :REPORT_TYPE_UNSPECIFIED, 0
      value :REPORT_TYPE_ON_DEMAND, 1
      value :REPORT_TYPE_HOURLY, 2
      value :REPORT_TYPE_DAILY, 3
      value :REPORT_TYPE_MONTHLY, 4
    end
  end
end

module Report
  module V1
    Report = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.Report").msgclass
    GetReportRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.GetReportRequest").msgclass
    GetReportResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.GetReportResponse").msgclass
    PutReportRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.PutReportRequest").msgclass
    PutReportResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.PutReportResponse").msgclass
    DeleteReportRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.DeleteReportRequest").msgclass
    MigrateReportRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.MigrateReportRequest").msgclass
    DeleteReportResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.DeleteReportResponse").msgclass
    ReportType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("report.v1.ReportType").enummodule
  end
end
