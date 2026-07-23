package routes

import (
	"github.com/labstack/echo/v4"

	"spotsync/internal/handler"
	"spotsync/pkg/middleware"
)

func RegisterRoutes(
	e *echo.Echo,
	authHandler *handler.AuthHandler,
	zoneHandler *handler.ZoneHandler,
	reservationHandler *handler.ReservationHandler,
) {


	api := e.Group("/api/v1")


	// Public routes
	api.GET("/", func(c echo.Context) error {

		return c.JSON(
			200,
			map[string]any{
				"success": true,
				"message": "Welcome to SpotSync API",
			},
		)
	})


	api.POST(
		"/auth/register",
		authHandler.Register,
	)


	api.POST(
		"/auth/login",
		authHandler.Login,
	)



	api.GET(
		"/zones",
		zoneHandler.GetAll,
	)


	api.GET(
		"/zones/:id",
		zoneHandler.GetByID,
	)



	// Protected routes

	protected := api.Group("")

	protected.Use(
		middleware.JWTMiddleware,
	)



	protected.POST(
		"/reservations",
		reservationHandler.Create,
	)


	protected.GET(
		"/reservations/my-reservations",
		reservationHandler.MyReservations,
	)


	protected.DELETE(
		"/reservations/:id",
		reservationHandler.Cancel,
	)



	// Admin routes

	admin := api.Group("")

	admin.Use(
		middleware.JWTMiddleware,
	)


	admin.Use(
		middleware.RoleMiddleware("admin"),
	)



	admin.POST(
		"/zones",
		zoneHandler.Create,
	)


	admin.GET(
		"/reservations",
		reservationHandler.GetAll,
	)
}