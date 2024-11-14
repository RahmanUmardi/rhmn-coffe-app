package common

import (
	"rhmn-coffe/shared/model"

	"github.com/gofiber/fiber/v2"
)

func SendSingleResponseCreated(ctx *fiber.Ctx, data interface{}, message string) error {
	return ctx.Status(fiber.StatusCreated).JSON(&model.SingleResponse{
		Status: model.Status{
			Code:    fiber.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}

func SendSingleResponseOk(ctx *fiber.Ctx, data interface{}, message string) error {
	return ctx.Status(fiber.StatusOK).JSON(&model.SingleResponse{
		Status: model.Status{
			Code:    fiber.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(&model.Status{
		Code:    code,
		Message: message,
	})
}
