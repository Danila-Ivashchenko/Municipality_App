package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (m *Middleware) WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := parser.ParseBearerToken(c)

		if token == "" {
			response.Error(c, errors.New("auth token not provided"))
			c.Abort()
			return
		}

		ctx := c.Request.Context()

		userToken, err := m.Params.UserAuthService.GetByTokenWithValidation(ctx, token)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		user, err := m.Params.UserService.GetUserByID(ctx, userToken.UserID)
		if err != nil {
			response.Error(c, errors.New("invalid bearer token"))
			c.Abort()
			return
		}

		slog.DebugContext(ctx, "parsed user", slog.Any("user_id", user.ID))
		c.Set("user", user)
		c.Set("user_auth_token", userToken)

		ctx = context_paylod_parser.SetUserToContext(ctx, user)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
