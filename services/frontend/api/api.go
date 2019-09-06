package api

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"github.com/mewil/portal/common/logger"
	"google.golang.org/grpc"
)

type FrontendService struct {
	log             logger.Logger
	jwtSecret       string
	authServiceAddr string
	authServiceConn *grpc.ClientConn
	userServiceAddr string
	userServiceConn *grpc.ClientConn
	postServiceAddr string
	postServiceConn *grpc.ClientConn
}

func NewAPI(log logger.Logger, baseRouter *gin.RouterGroup, jwtSecret string) error {
	s := FrontendService{
		log:       log,
		jwtSecret: jwtSecret,
	}
	if err := s.createServiceConnections(); err != nil {
		return err
	}
	s.createViews(baseRouter)
	return nil
}

const (
	authServiceArrEnvKey = "AUTH_SERVICE_ADDR"
	userServiceArrEnvKey = "USER_SERVICE_ADDR"
	postServiceArrEnvKey = "POST_SERVICE_ADDR"
)

func (s *FrontendService) createServiceConnections() error {
	ctx := context.Background()
	s.authServiceAddr = os.Getenv(authServiceArrEnvKey)
	if err := createGrpcConnection(ctx, &s.authServiceConn, s.authServiceAddr); err != nil {
		return err
	}
	s.userServiceAddr = os.Getenv(userServiceArrEnvKey)
	if err := createGrpcConnection(ctx, &s.userServiceConn, s.userServiceAddr); err != nil {
		return err
	}
	s.postServiceAddr = os.Getenv(postServiceArrEnvKey)
	if err := createGrpcConnection(ctx, &s.postServiceConn, s.postServiceAddr); err != nil {
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

const rateLimitTPS = 0.5

func (s *FrontendService) createViews(baseRouter *gin.RouterGroup) {
	rateLimitMiddleware := tollbooth_gin.LimitHandler(tollbooth.NewLimiter(rateLimitTPS, nil))
	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/signin", rateLimitMiddleware, s.PostAuthSignIn())
	authRouter.POST("/signup", rateLimitMiddleware, s.PostAuthSignUp())
	userRouter := baseRouter.Group("/user")
	userRouter.GET("/:id", s.UserAuthMiddleware(), s.GetUser())
	userRouter.GET("/:id/followers", s.UserAuthMiddleware(), s.GetUserFollowers())
	userRouter.GET("/:id/following", s.UserAuthMiddleware(), s.GetUserFollowing())
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
