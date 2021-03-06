package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc"
)

type FrontendSvc struct {
	log         logger.Logger
	jwtSecret   string
	authSvcAddr string
	authSvcConn *grpc.ClientConn
	userSvcAddr string
	userSvcConn *grpc.ClientConn
	postSvcAddr string
	postSvcConn *grpc.ClientConn
	fileSvcAddr string
	fileSvcConn *grpc.ClientConn
}

const (
	fileBufferSize = 4096
	fileSizeLimit  = 2000 * fileBufferSize
)

func NewAPI(log logger.Logger, baseRouter *gin.RouterGroup, jwtSecret, authSvcAddr, userSvcAddr, postSvcAddr, fileSvcAddr string) error {
	s := FrontendSvc{
		log:         log,
		jwtSecret:   jwtSecret,
		authSvcAddr: authSvcAddr,
		userSvcAddr: userSvcAddr,
		postSvcAddr: postSvcAddr,
		fileSvcAddr: fileSvcAddr,
	}
	s.createViews(baseRouter)
	return s.createSvcConnections()
}

func NewTestAPI() FrontendSvc {
	log, _ := logger.NewLogger("test")
	s := FrontendSvc{
		log:       log,
		jwtSecret: "jwtSecret",
	}
	return s
}

func (s *FrontendSvc) createSvcConnections() error {
	ctx := context.Background()
	if err := createGrpcConnection(ctx, &s.authSvcConn, s.authSvcAddr); err != nil {
		return err
	}
	if err := createGrpcConnection(ctx, &s.userSvcConn, s.userSvcAddr); err != nil {
		return err
	}
	if err := createGrpcConnection(ctx, &s.postSvcConn, s.postSvcAddr); err != nil {
		return err
	}
	if err := createGrpcConnection(ctx, &s.fileSvcConn, s.fileSvcAddr); err != nil {
		return err
	}
	return nil
}

func createGrpcConnection(ctx context.Context, conn **grpc.ClientConn, addr string) error {
	err := *new(error)
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(3*time.Second))
	return err
}

type AuthSvcInjector func() pb.AuthSvcClient

func (s *FrontendSvc) injectAuthSvcClient() AuthSvcInjector {
	return func() pb.AuthSvcClient {
		return pb.NewAuthSvcClient(s.authSvcConn)
	}
}

type UserSvcInjector func() pb.UserSvcClient

func (s *FrontendSvc) injectUserSvcClient() UserSvcInjector {
	return func() pb.UserSvcClient {
		return pb.NewUserSvcClient(s.userSvcConn)
	}
}

type PostSvcInjector func() pb.PostSvcClient

func (s *FrontendSvc) injectPostSvcClient() PostSvcInjector {
	return func() pb.PostSvcClient {
		return pb.NewPostSvcClient(s.postSvcConn)
	}
}

type FileSvcInjector func() pb.FileSvcClient

func (s *FrontendSvc) injectFileSvcClient() FileSvcInjector {
	return func() pb.FileSvcClient {
		return pb.NewFileSvcClient(s.fileSvcConn)
	}
}
