package api

import (
	"net/http"
	"strconv"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

const rateLimitTPS = 0.5

func (s *FrontendSvc) createViews(baseRouter *gin.RouterGroup) {
	rateLimitMiddleware := tollbooth_gin.LimitHandler(tollbooth.NewLimiter(rateLimitTPS, nil))
	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/signin", rateLimitMiddleware, s.PostAuthSignIn(s.injectAuthSvcClient(), s.injectUserSvcClient()))
	authRouter.POST("/signup", rateLimitMiddleware, s.PostAuthSignUp(s.injectAuthSvcClient(), s.injectUserSvcClient()))

	userRouter := baseRouter.Group("/user")
	userRouter.Use(s.UserAuthMiddleware())
	userRouter.GET("/:user_id", s.GetUser(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/followers", s.GetUserFollowers(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/following", s.GetUserFollowing(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/profile", s.GetUserProfile(s.injectUserSvcClient(), s.injectPostSvcClient()))
	userRouter.POST("/:user_id/follow", s.PostUserFollow(s.injectUserSvcClient()))
	userRouter.POST("/:user_id/unfollow", s.PostUserUnfollow(s.injectUserSvcClient()))

	postRouter := baseRouter.Group("/post")
	postRouter.Use(s.UserAuthMiddleware())
	postRouter.GET("/", s.GetFeed(s.injectPostSvcClient()))
	postRouter.GET("/:post_id", s.GetPost(s.injectPostSvcClient()))
	postRouter.GET("/:post_id/likes", s.GetPostLikes(s.injectPostSvcClient()))
	postRouter.GET("/:post_id/comments", s.GetPostComments(s.injectPostSvcClient()))
	postRouter.POST("/", s.PostPost(s.injectPostSvcClient(), s.injectFileSvcClient()))
	postRouter.POST("/:post_id/like", s.PostPostLike(s.injectPostSvcClient()))
	postRouter.POST("/:post_id/comment", s.PostComment(s.injectPostSvcClient()))

	commentRouter := baseRouter.Group("/comment")
	commentRouter.GET("/:comment_id", s.GetComment(s.injectPostSvcClient()))
	commentRouter.GET("/:comment_id/likes", s.GetCommentLikes(s.injectPostSvcClient()))
	commentRouter.POST("/:comment_id/like", s.PostCommentLike(s.injectPostSvcClient()))

	fileRouter := baseRouter.Group("/file")
	fileRouter.GET("/:file_id", s.GetFile(s.injectFileSvcClient()))
}

const (
	statusKey  = "status"
	messageKey = "message"
	dataKey    = "data"
)

// ResponseOK is a helper function to make sure all valid HTTP responses
// follow the same format
func ResponseOK(c *gin.Context, message string, data gin.H) {
	ResponseValid(c, http.StatusOK, message, data)
}

// ResponseValid is a helper function to make sure all valid HTTP responses
// follow the same format
func ResponseValid(c *gin.Context, statusCode int, message string, data gin.H) {
	c.JSON(statusCode, gin.H{
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

// GetPageQueryParam is a helper function that returns the value of a `page` query parameter
func GetPageQueryParam(c *gin.Context) (uint32, error) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 32)
	return uint32(page), err
}
