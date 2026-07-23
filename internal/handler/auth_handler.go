package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"spotsync/internal/dto"
	"spotsync/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Invalid request body",
		})
	}

	res, err := h.service.Register(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"success": true,
		"message": "User registered successfully",
		"data":    res,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Invalid request body",
		})
	}

	res, err := h.service.Login(req)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Login successful",
		"data":    res,
	})
}