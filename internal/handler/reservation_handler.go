package handler

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"spotsync/internal/dto"
	"spotsync/internal/service"
)

type ReservationHandler struct {
	service *service.ReservationService
}


func NewReservationHandler(
	s *service.ReservationService,
) *ReservationHandler {

	return &ReservationHandler{
		service: s,
	}
}



// Create Reservation
func (h *ReservationHandler) Create(
	c echo.Context,
) error {


	var req dto.CreateReservationRequest


	if err := c.Bind(&req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"success": false,
				"message": "Invalid request",
			},
		)
	}



	userID, ok := c.Get("userID").(uint)

	if !ok {

		return c.JSON(
			http.StatusUnauthorized,
			map[string]any{
				"success": false,
				"message": "Unauthorized",
			},
		)
	}



	res, err := h.service.CreateReservation(
		req.ZoneID,
		userID,
		req.LicensePlate,
	)


	if err != nil {

		return c.JSON(
			http.StatusConflict,
			map[string]any{
				"success": false,
				"message": err.Error(),
			},
		)
	}



	return c.JSON(
		http.StatusCreated,
		map[string]any{
			"success": true,
			"message": "Reservation confirmed successfully",
			"data":    res,
		},
	)
}



// Get My Reservations
func (h *ReservationHandler) MyReservations(
	c echo.Context,
) error {


	userID, ok := c.Get("userID").(uint)

	if !ok {

		return c.JSON(
			http.StatusUnauthorized,
			map[string]any{
				"success": false,
				"message": "Unauthorized",
			},
		)
	}



	res, err := h.service.GetMyReservations(userID)


	if err != nil {

		return c.JSON(
			http.StatusInternalServerError,
			map[string]any{
				"success": false,
				"message": err.Error(),
			},
		)
	}



	return c.JSON(
		http.StatusOK,
		map[string]any{
			"success": true,
			"message": "My reservations retrieved successfully",
			"data":    res,
		},
	)
}



// Cancel Reservation
func (h *ReservationHandler) Cancel(
	c echo.Context,
) error {


	userID, ok := c.Get("userID").(uint)

	if !ok {

		return c.JSON(
			http.StatusUnauthorized,
			map[string]any{
				"success": false,
				"message": "Unauthorized",
			},
		)
	}



	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]any{
				"success": false,
				"message": "Invalid reservation id",
			},
		)
	}



	err = h.service.CancelReservation(
		uint(id),
		userID,
	)


	if err != nil {

		return c.JSON(
			http.StatusForbidden,
			map[string]any{
				"success": false,
				"message": err.Error(),
			},
		)
	}



	return c.JSON(
		http.StatusOK,
		map[string]any{
			"success": true,
			"message": "Reservation cancelled successfully",
		},
	)
}



// Get All Reservations (Admin)
func (h *ReservationHandler) GetAll(
	c echo.Context,
) error {


	res, err := h.service.GetAllReservations()


	if err != nil {

		return c.JSON(
			http.StatusInternalServerError,
			map[string]any{
				"success": false,
				"message": err.Error(),
			},
		)
	}



	return c.JSON(
		http.StatusOK,
		map[string]any{
			"success": true,
			"message": "Reservations retrieved successfully",
			"data":    res,
		},
	)
}