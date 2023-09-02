package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/slavik22/chat/token"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.JWTMaker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authorizationHeader := ctx.Request().Header.Get(authorizationHeaderKey)

			if len(authorizationHeader) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "auth header is empty")
			}

			fields := strings.Fields(authorizationHeader)

			if len(fields) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth header")
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				return echo.NewHTTPError(http.StatusUnauthorized, "auth type is incorrect")
			}

			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is incorrect")
			}

			ctx.Set(authorizationPayloadKey, payload.UserId)
			return next(ctx)
		}
	}
}

func getUserId(c echo.Context) (int64, error) {
	id, ok := c.Get(authorizationPayloadKey).(int64)

	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return id, nil
}
