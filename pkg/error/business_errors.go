package error

import (
	rpc_error "github.com/Freeline95/go-common/pkg/grpc/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBusinessError(code string, message string) error {
	status := status.New(codes.Internal, message)
	status, err := status.WithDetails(&rpc_error.BusinessError{Code: code, Message: message})
	if err != nil {
		return err
	}

	return status.Err()
}

func CheckBusinessErrorCode(err error, code string) bool {
	if st, ok := status.FromError(err); ok {
		for _, detail := range st.Details() {
			if d, ok := detail.(*rpc_error.BusinessError); ok {
				if d.GetCode() == code {
					return true
				}
			}
		}
	}

	return false
}
