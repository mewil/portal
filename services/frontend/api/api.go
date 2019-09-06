package api

import (
	"context"
	"net/http"
	"os"
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
}

func NewAPI(log logger.Logger, baseRouter *gin.RouterGroup, jwtSecret string) error {
	s := FrontendSvc{
		log:       log,
		jwtSecret: jwtSecret,
	}
	s.createViews(baseRouter)
	return s.createSvcConnections()
}

const (
	authSvcArrEnvKey = "AUTH_SERVICE_ADDR"
	userSvcArrEnvKey = "USER_SERVICE_ADDR"
	postSvcArrEnvKey = "POST_SERVICE_ADDR"
)

func (s *FrontendSvc) createSvcConnections() error {
	ctx := context.Background()
	s.authSvcAddr = os.Getenv(authSvcArrEnvKey)
	if err := createGrpcConnection(ctx, &s.authSvcConn, s.authSvcAddr); err != nil {
		return err
	}
	s.userSvcAddr = os.Getenv(userSvcArrEnvKey)
	if err := createGrpcConnection(ctx, &s.userSvcConn, s.userSvcAddr); err != nil {
		return err
	}
	s.postSvcAddr = os.Getenv(postSvcArrEnvKey)
	if err := createGrpcConnection(ctx, &s.postSvcConn, s.postSvcAddr); err != nil {
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

const (
	statusKey  = "status"
	messageKey = "message"
	dataKey    = "data"
)

// ResponseOK is a helper function to make sure all valid HTTP responses
// follow the same format
func ResponseOK(c *gin.Context, message string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		statusKey:  true,
		messageKey: message,
		dataKey:    data,
	})
}

// ResponseError is a helper function to make sure all invalid HTTP responses
// follow the same format
func ResponseError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		statusKey:  false,
		messageKey: message,
	})
}

// ResponseInternalError is a helper function to make sure all invalid HTTP responses
// follow the same format
func ResponseInternalError(c *gin.Context) {
	ResponseError(c, http.StatusInternalServerError, "an unexpected error occurred, please try again")
}
