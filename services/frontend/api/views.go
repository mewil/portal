package api

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

const rateLimitTPS = 0.5

func (s *FrontendSvc) createViews(baseRouter *gin.RouterGroup) {
	rateLimitMiddleware := tollbooth_gin.LimitHandler(tollbooth.NewLimiter(rateLimitTPS, nil))
	authRouter := baseRouter.Group("/auth")
	authRouter.POST("/signin", rateLimitMiddleware, s.PostAuthSignIn(s.injectAuthSvcClient()))
	authRouter.POST("/signup", rateLimitMiddleware, s.PostAuthSignUp(s.injectAuthSvcClient(), s.injectUserSvcClient()))

	userRouter := baseRouter.Group("/user")
	userRouter.Use(s.UserAuthMiddleware())
	userRouter.GET("/:user_id", s.GetUser(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/followers", s.GetUserFollowers(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/following", s.GetUserFollowing(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/profile", s.GetUserProfile(s.injectPostSvcClient(), s.injectPostSvcClient()))
	userRouter.POST("/:user_id/follow", s.PostUserFollow(s.injectUserSvcClient()))
	userRouter.POST("/:user_id/unfollow", s.PostUserUnfollow(s.injectUserSvcClient()))

	postRouter := baseRouter.Group("/post")
	postRouter.Use(s.UserAuthMiddleware())
	postRouter.GET("/", s.GetFeed(s.injectPostSvcClient()))
	postRouter.GET("/:post_id", s.GetPost(s.injectPostSvcClient()))
	postRouter.GET("/:post_id/likes", s.GetPostLikes(s.injectPostSvcClient()))
	postRouter.GET("/:post_id/comments", s.GetPostComments(s.injectPostSvcClient()))
	postRouter.POST("/", s.PostPost(s.injectPostSvcClient()))
	postRouter.POST("/:post_id/like", s.PostPostLike(s.injectPostSvcClient()))
	postRouter.POST("/:post_id/comment", s.PostComment(s.injectPostSvcClient()))

	commentRouter := baseRouter.Group("/comment")
	commentRouter.GET("/:comment_id", s.GetComment(s.injectPostSvcClient()))
	commentRouter.GET("/:comment_id/likes", s.GetCommentLikes(s.injectPostSvcClient()))
	commentRouter.POST("/:comment_id/like", s.PostCommentLike(s.injectPostSvcClient()))
}
