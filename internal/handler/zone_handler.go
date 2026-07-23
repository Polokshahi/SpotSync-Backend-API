package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"spotsync/internal/dto"
	"spotsync/internal/service"
)

type ZoneHandler struct {
	service *service.ZoneService
}

func NewZoneHandler(s *service.ZoneService) *ZoneHandler {
	return &ZoneHandler{
		service: s,
	}
}

func (h *ZoneHandler) Create(c echo.Context) error {

	var req dto.CreateZoneRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Invalid request",
		})
	}

	res, err := h.service.CreateZone(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"success": true,
		"message": "Parking zone created successfully",
		"data":    res,
	})
}

func (h *ZoneHandler) GetAll(c echo.Context) error {

	res, err := h.service.GetAllZones()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Parking zones retrieved successfully",
		"data":    res,
	})
}

func (h *ZoneHandler) GetByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	res, err := h.service.GetZoneByID(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Parking zone retrieved successfully",
		"data":    res,
	})
}