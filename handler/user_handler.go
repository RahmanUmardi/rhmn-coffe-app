package handler

import (
	"fmt"
	"rhmn-coffe/config"
	"rhmn-coffe/entity"
	"rhmn-coffe/middleware"
	"rhmn-coffe/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUc         usecase.UserUsecase
	rg             fiber.Router
	authMiddleware middleware.AuthMiddleware
}

func (u *UserHandler) FindAll(ctx *fiber.Ctx) error {

	users, err := u.userUc.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprint("User is empty")})
	}

	response := struct {
		Message string        `json:"message"`
		Data    []entity.User `json:"data"`
	}{
		Message: "List of user is empty",
		Data:    users,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (u *UserHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := u.userUc.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("User with id %s not found", id)})
	}

	response := struct {
		Message string      `json:"message"`
		Data    entity.User `json:"data"`
	}{
		Message: "Success Get User By Id",
		Data:    user,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (u *UserHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var input entity.UpdateUser
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	updatedUser, err := u.userUc.Update(id, input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully", "data": updatedUser})
}

func (u *UserHandler) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	err := u.userUc.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("User with ID %s not found", id)})
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User deleted successfully",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (u *UserHandler) Route() {
	u.rg.Get(config.GetUserList, u.authMiddleware.RequireToken("admin"), u.FindAll)
	u.rg.Get(config.GetUser, u.authMiddleware.RequireToken("admin"), u.FindById)
	u.rg.Put(config.PutUser, u.authMiddleware.RequireToken("admin"), u.Update)
	u.rg.Delete(config.DeleteUser, u.authMiddleware.RequireToken("admin"), u.Delete)
}

func NewUserHandler(userUc usecase.UserUsecase, authMiddleware middleware.AuthMiddleware, rg fiber.Router) *UserHandler {
	return &UserHandler{userUc: userUc, authMiddleware: authMiddleware, rg: rg}
}
