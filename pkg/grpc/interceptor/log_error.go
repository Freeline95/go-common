package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	log "github.com/sirupsen/logrus"
)

func LogErrorInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Errorf("GRPC request failed with code %d and message: %s", s.Code(), s.Message())
		} else {
			log.Errorf("GRPC request failed with error: %v", err)
		}
	}
	return resp, err
}