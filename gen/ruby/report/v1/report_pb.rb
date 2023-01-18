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
