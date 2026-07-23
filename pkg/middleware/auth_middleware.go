package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"spotsync/internal/auth"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"success": false,
				"message": "Missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"success": false,
				"message": "Invalid authorization header",
			})
		}

		claims, err := auth.ValidateJWT(parts[1])

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"success": false,
				"message": "Invalid or expired token",
			})
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		return next(c)
	}
}