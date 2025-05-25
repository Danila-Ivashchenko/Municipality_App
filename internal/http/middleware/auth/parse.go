package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
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

		if user.IsBlocked {
			response.Error(c, core_errors.UserNotFound)
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

func (m *Middleware) WithAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := context_paylod_parser.GetUserFromContext(ctx)
		if user == nil {
			response.Unauthorized(c, core_errors.UserNotAuth)
			c.Abort()
			return
		}

		if !user.IsAdmin {
			response.Forbidden(c, core_errors.UserIsNotAdmin)
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(ctx)
	}
}

func (m *Middleware) WithCanEdit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := context_paylod_parser.GetUserFromContext(ctx)
		if user == nil {
			response.Unauthorized(c, core_errors.UserNotAuth)
			c.Abort()
			return
		}

		permissions, err := m.Params.UserPermissionService.GetUserPermissions(ctx, user.ID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		canEdit := false

		for _, permission := range permissions {
			if permission == entity.Write {
				canEdit = true
				break
			}
		}

		if !canEdit && !user.IsAdmin {
			response.Forbidden(c, core_errors.UserHasNoPermissionToWrite)
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(ctx)
	}
}

func (m *Middleware) WithCanDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := context_paylod_parser.GetUserFromContext(ctx)
		if user == nil {
			response.Unauthorized(c, core_errors.UserNotAuth)
			c.Abort()
			return
		}

		permissions, err := m.Params.UserPermissionService.GetUserPermissions(ctx, user.ID)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		canDelete := false

		for _, permission := range permissions {
			if permission == entity.Delete {
				canDelete = true
				break
			}
		}

		if !canDelete && !user.IsAdmin {
			response.Forbidden(c, core_errors.UserHasNoPermissionToDelete)
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(ctx)
	}
}
