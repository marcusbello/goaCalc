// Code generated by goa v3.7.13, DO NOT EDIT.
//
// calc gRPC client types
//
// Command:
// $ goa gen goaCalc/design

package client

import (
	calc "goaCalc/gen/calc"
	calcpb "goaCalc/gen/grpc/calc/pb"
)

// NewProtoMultiplyRequest builds the gRPC request type from the payload of the
// "multiply" endpoint of the "calc" service.
func NewProtoMultiplyRequest(payload *calc.MultiplyPayload) *calcpb.MultiplyRequest {
	message := &calcpb.MultiplyRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewMultiplyResult builds the result type of the "multiply" endpoint of the
// "calc" service from the gRPC response type.
func NewMultiplyResult(message *calcpb.MultiplyResponse) int {
	result := int(message.Field)
	return result
}
