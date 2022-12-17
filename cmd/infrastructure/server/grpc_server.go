package server

import (
	grpcService "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/grpc/pb/reader_service"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/utils/constants"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func (s *server) NewGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}

	grpcServer := newGrpcServer()

	grpcService := grpcService.NewReaderGrpcService(s.cfg, s.ps, s.log, s.m)
	reader_service.RegisterProductReaderServiceServer(grpcServer, grpcService)
	grpc_prometheus.Register(grpcServer)

	if s.cfg.GRPC.Development {
		reflection.Register(grpcServer)
	}

	go func() {
		s.log.Infof("Reader gRPC server is listening on port: %s", s.cfg.GRPC.Port)
		s.log.Fatal(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}

func newGrpcServer() *grpc.Server {
	return grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: constants.MAX_CONNECTION_IDLE * time.Minute,
			Timeout:           constants.GRPC_TIMEOUT * time.Second,
			MaxConnectionAge:  constants.MAX_CONNECTION_AGE * time.Minute,
			Time:              constants.GRPC_TIME * time.Minute,
		}),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)
}
