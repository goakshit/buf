# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: payment/v1/payment.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("payment/v1/payment.proto", :syntax => :proto3) do
    add_message "payment.v1.Payment" do
      optional :payment_type, :enum, 1, "payment.v1.PaymentType", json_name: "paymentType"
      optional :payment_id, :string, 2, json_name: "paymentId"
      optional :name, :string, 3, json_name: "name"
    end
    add_message "payment.v1.GetPaymentRequest" do
      optional :payment_id, :string, 1, json_name: "paymentId"
    end
    add_message "payment.v1.GetPaymentResponse" do
      optional :payment, :message, 1, "payment.v1.Payment", json_name: "payment"
    end
    add_message "payment.v1.PutPaymentRequest" do
      optional :payment_type, :enum, 1, "payment.v1.PaymentType", json_name: "paymentType"
      optional :name, :string, 2, json_name: "name"
    end
    add_message "payment.v1.PutPaymentResponse" do
      optional :payment, :message, 1, "payment.v1.Payment", json_name: "payment"
    end
    add_message "payment.v1.DeletePaymentRequest" do
      optional :payment_id, :string, 1, json_name: "paymentId"
    end
    add_message "payment.v1.DeletePaymentResponse" do
    end
    add_enum "payment.v1.PaymentType" do
      value :PAYMENT_TYPE_UNSPECIFIED, 0
      value :PAYMENT_TYPE_ON_DEMAND, 1
      value :PAYMENT_TYPE_HOURLY, 2
      value :PAYMENT_TYPE_DAILY, 3
      value :PAYMENT_TYPE_MONTHLY, 4
    end
  end
end

module Payment
  module V1
    Payment = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.Payment").msgclass
    GetPaymentRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.GetPaymentRequest").msgclass
    GetPaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.GetPaymentResponse").msgclass
    PutPaymentRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.PutPaymentRequest").msgclass
    PutPaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.PutPaymentResponse").msgclass
    DeletePaymentRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.DeletePaymentRequest").msgclass
    DeletePaymentResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.DeletePaymentResponse").msgclass
    PaymentType = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("payment.v1.PaymentType").enummodule
  end
end
