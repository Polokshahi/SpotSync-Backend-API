package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(roles ...string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			role, ok := c.Get("role").(string)

			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"success": false,
					"message": "Unauthorized",
				})
			}

			for _, allowed := range roles {
				if role == allowed {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]any{
				"success": false,
				"message": "Forbidden",
			})
		}
	}
}