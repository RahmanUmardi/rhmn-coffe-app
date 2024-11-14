package handler

import (
	"fmt"
	"rhmn-coffe/config"
	"rhmn-coffe/entity"
	"rhmn-coffe/entity/dto"
	"rhmn-coffe/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase usecase.AuthUseCase
	rg          fiber.Router
}

func (a *AuthHandler) loginHandler(ctx *fiber.Ctx) error {
	var payload dto.AuthRequestDto

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Sprintf("Invalid input: %s", err.Error())})
	}

	token, err := a.authUsecase.Login(payload)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": fmt.Sprintf("Failed to login: %s", err.Error())})
	}

	return ctx.Status(fiber.StatusOK).JSON(token)
}

func (a *AuthHandler) registerHandler(ctx *fiber.Ctx) error {
	var payload dto.AuthRequestDto

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Sprintf("Invalid input: %s", err.Error())})
	}

	user, err := a.authUsecase.Register(payload)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": fmt.Sprintf("Failed to register: %s", err.Error())})
	}
	response := struct {
		Message string      `json:"message"`
		Data    entity.User `json:"data"`
	}{
		Message: "Success Register User",
		Data:    user,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (a *AuthHandler) Route() {
	a.rg.Post(config.Login, a.loginHandler)
	a.rg.Post(config.Register, a.registerHandler)
}

func NewAuthHandler(authUc usecase.AuthUseCase, rg fiber.Router) *AuthHandler {
	return &AuthHandler{authUsecase: authUc, rg: rg}
}
