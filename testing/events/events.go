package events

import (
	paymentv1 "github.com/goakshit/buf/gen/go/payment/v1"
	"google.golang.org/protobuf/proto"
)

var (
	Events = []proto.Message{
		&paymentv1.Payment{},
		&paymentv1.PaymentCopy{},
	}
)
