# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: payment/v1/payment.proto for package 'Payment.V1'

require 'grpc'
require 'payment/v1/payment_pb'

module Payment
  module V1
    module PaymentService
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'payment.v1.PaymentService'

        rpc :GetPayment, ::Payment::V1::GetPaymentRequest, ::Payment::V1::GetPaymentResponse
        rpc :PutPayment, ::Payment::V1::PutPaymentRequest, ::Payment::V1::PutPaymentResponse
        rpc :DeletePayment, ::Payment::V1::DeletePaymentRequest, ::Payment::V1::DeletePaymentResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end