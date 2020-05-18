package error

import (
	"fmt"

	"github.com/gogo/status"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

// CustomError ...
type CustomError struct {
	Code         codes.Code
	Message      string
	Err          error
	ProtoDetails []proto.Message
}

// NewCustomErrorf - Creates a new Custom Error Object
func NewCustomErrorf(c codes.Code, format string, a ...interface{}) *CustomError {
	return NewCustomError(c, fmt.Sprintf(format, a...))
}

// NewCustomError ...
func NewCustomError(c codes.Code, message string) *CustomError {
	return &CustomError{Code: c, Message: message, Err: errors.New(message)}
}

// NewCustomWrapError ...
func (ce *CustomError) NewCustomWrapError(c codes.Code, err error) *CustomError {
	return &CustomError{Code: c, Message: err.Error(), Err: err}
}

// WithDetails - Adds Proto Error Details so Clients can get additional information about the error
/**
The way this works for pushing additional information to clients.
		customProtoError := &data.Error{
			Message:         "Check username",
			Code:            1404,
			Type:            "Skyflow",
			DetailedMessage: desc,
		}
		customError = customError.WithDetails(customProtoError)

To use Standard Errors Defined by grpc for additional information
		customError := myErrors.NewCustomErrorf(codes.NotFound, "user %s doesn't exist", req.Name)
		v := &rpc.BadRequest_FieldViolation{
			Field:       "username",
			Description: desc,
		}
		br := &rpc.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		customError = customError.WithDetails(br)

*/
func (ce *CustomError) WithDetails(details ...proto.Message) *CustomError {
	for _, detail := range details {
		ce.ProtoDetails = append(ce.ProtoDetails, detail)
	}
	return ce
}

// BaseError ...
func (ce *CustomError) BaseError() error {
	return ce.Err
}

// CustomError ...
func (ce *CustomError) Error() string {
	return ce.Err.Error()
}

// GetStatusError ...
func (ce *CustomError) GetStatusError() *status.Status {
	status := status.Newf(ce.Code, ce.Message)

	if len(ce.ProtoDetails) > 0 {
		for _, detail := range ce.ProtoDetails {
			status, _ = status.WithDetails(detail)
		}
	}

	return status
}

// ConvertToGRPCError ...
func ConvertToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	switch castedError := err.(type) {
	case *CustomError:
		fmt.Printf("Custom Base Error was %+v\n", castedError.BaseError())
		if len(castedError.ProtoDetails) > 0 {
			for _, detail := range castedError.ProtoDetails {
				fmt.Printf("Custom Detail Error was %+v\n", detail)
			}
		}
		return castedError.GetStatusError().Err()
	default:
		fmt.Printf("Default Error was %+v\n", castedError)
		s, ok := status.FromError(err)
		if ok {
			return s.Err()
		}
		// Create a new Status Error
		return status.Newf(codes.Internal, err.Error()).Err()
	}
}
