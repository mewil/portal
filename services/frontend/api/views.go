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
	userRouter.GET("/:user_id", s.UserAuthMiddleware(), s.GetUser(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/followers/:page", s.UserAuthMiddleware(), s.GetUserFollowers(s.injectUserSvcClient()))
	userRouter.GET("/:user_id/following/:page", s.UserAuthMiddleware(), s.GetUserFollowing(s.injectUserSvcClient()))
}
