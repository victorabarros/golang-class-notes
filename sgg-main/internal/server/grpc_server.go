package server

import (
	"context"
	"fmt"
	"net"

	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/core"
	grpct "github.com/ricardoerikson/sgg/internal/transport/grpc"
	"github.com/ricardoerikson/sgg/internal/transport/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	server         *grpc.Server
	articleHandler *grpct.ArticleHandler
	config         *cfg.Configuration
}

// NewGRPCServer creates a new gRPC server
func NewGRPCServer(_ context.Context,
	articleHandler *grpct.ArticleHandler,
	config *cfg.Configuration,
) core.Server {

	return &grpcServer{
		articleHandler: articleHandler,
		config:         config,
	}
}

func (s *grpcServer) Name() string {
	return "grpc-server"
}

func (s *grpcServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.GRPC.Port))
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			ContextPropagationUnaryServerInterceptor(ctx),
		),
	}

	s.server = grpc.NewServer(opts...)
	pb.RegisterArticlesServer(s.server, s.articleHandler)
	reflection.Register(s.server)
	return s.server.Serve(lis)
}

func (s *grpcServer) Shutdown(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}

func ContextPropagationUnaryServerInterceptor(baseContext context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// if md, ok := metadata.FromIncomingContext(ctx); ok {
		// 	ctx = metadata.NewOutgoingContext(ctx, md)
		// }

		return handler(baseContext, req)
	}
}
