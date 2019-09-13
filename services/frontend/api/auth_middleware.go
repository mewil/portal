package api

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	authCookieKey           = "auth_cookie"
	authHeaderKey           = "Authorization"
	authHeaderValuePrefix   = "Bearer "
	userIDClaimsKey         = "userID"
	authGroupClaimsKey      = "authGroup"
	expiresAtClaimsKey      = "expiresAt"
	expirationPeriod        = 7 * 24 * time.Hour
	userAuthorizationGroup  = "user"
	adminAuthorizationGroup = "admin"
)

func GetUserID(c *gin.Context) string {
	return c.GetString(userIDClaimsKey)
}

func (s *FrontendSvc) UserAuthMiddleware() gin.HandlerFunc {
	return s.authMiddleware(userAuthorizationGroup)
}

func (s *FrontendSvc) AdminAuthMiddleware() gin.HandlerFunc {
	return s.authMiddleware(adminAuthorizationGroup)
}

func (s *FrontendSvc) authMiddleware(authGroup string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := s.validateToken(c, authGroup); err != nil {
			s.log.Info("unauthorized request", err)
			ResponseError(c, http.StatusUnauthorized, "please provide a valid authorization cookie or token")
			return
		}
		c.Next()
	}
}

func parseToken(c *gin.Context) (string, error) {
	cookieToken, err := c.Cookie(authCookieKey)
	if err == nil {
		return cookieToken, nil
	}
	authHeader := c.Request.Header.Get(authHeaderKey)
	if authHeader != "" {
		return strings.ReplaceAll(authHeader, authHeaderValuePrefix, ""), nil
	}
	return "", errors.New("no auth token provided")
}

func (s *FrontendSvc) validateToken(c *gin.Context, authGroup string) error {
	tokenString, err := parseToken(c)
	if err != nil {
		return err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(jwt.SigningMethodHS512.Name) != token.Method {
			return nil, errors.New("invalid signing algorithm")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	group, ok := claims[authGroupClaimsKey].(string)
	if !ok || group != authGroup {
		return errors.New("invalid auth group")
	}
	userIDString, ok := claims[userIDClaimsKey].(string)
	if !ok {
		return errors.New("invalid user UUID value")
	}
	c.Set(userIDClaimsKey, userIDString)
	expiresAtString, ok := claims[expiresAtClaimsKey].(string)
	if !ok {
		return errors.New("invalid expired at value")
	}
	expiresAt, err := time.Parse(time.RFC3339Nano, expiresAtString)
	if err != nil {
		return err
	}
	if expiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}
	return nil
}

func (s *FrontendSvc) GenerateUserAuthToken(userID string) (string, error) {
	return s.generateToken(userID, userAuthorizationGroup)
}

func (s *FrontendSvc) GenerateAdminAuthToken(userID string) (string, error) {
	return s.generateToken(userID, adminAuthorizationGroup)
}

func (s *FrontendSvc) generateToken(userID string, authGroup string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.MapClaims{
		userIDClaimsKey:    userID,
		authGroupClaimsKey: authGroup,
		expiresAtClaimsKey: time.Now().Add(expirationPeriod).Format(time.RFC3339Nano),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
