package handler

import (
	"fmt"
	"rhmn-coffe/config"
	"rhmn-coffe/entity"
	"rhmn-coffe/middleware"
	"rhmn-coffe/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUc      usecase.ProductUsecase
	rg             fiber.Router
	authMiddleware middleware.AuthMiddleware
}

func (p *ProductHandler) Create(c *fiber.Ctx) error {
	var payload entity.Product

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Sprintf("Invalid input: %s", err.Error())})
	}

	Product, err := p.productUc.Create(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": fmt.Sprintf("Failed to create product: %s", err.Error())})
	}

	response := struct {
		Message string         `json:"message"`
		Data    entity.Product `json:"data"`
	}{
		Message: "Product Created",
		Data:    Product,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (p *ProductHandler) FindAll(ctx *fiber.Ctx) error {

	products, err := p.productUc.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprint("Product is empty")})
	}

	response := struct {
		Message string           `json:"message"`
		Data    []entity.Product `json:"data"`
	}{
		Message: "Succes Get All Product",
		Data:    products,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (p *ProductHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	product, err := p.productUc.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("Product with id %s not found", id)})
	}

	response := struct {
		Message string         `json:"message"`
		Data    entity.Product `json:"data"`
	}{
		Message: "Success Get Product By Id",
		Data:    product,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (p *ProductHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var input entity.UpdateProduct
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	updateProduct, err := p.productUc.Update(id, input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := struct {
		Message string         `json:"message"`
		Data    entity.Product `json:"data"`
	}{
		Message: "Success Update Product By Id",
		Data:    updateProduct,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (p *ProductHandler) Delete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	err := p.productUc.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("Product with ID %s not found", id)})
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Product deleted successfully",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (p *ProductHandler) Route() {
	p.rg.Post(config.PostProduct, p.authMiddleware.RequireToken("employee"), p.Create)
	p.rg.Get(config.GetProductList, p.authMiddleware.RequireToken("employee"), p.FindAll)
	p.rg.Get(config.GetProduct, p.authMiddleware.RequireToken("employee"), p.FindById)
	p.rg.Put(config.PutProduct, p.authMiddleware.RequireToken("employee"), p.Update)
	p.rg.Delete(config.DeleteProduct, p.authMiddleware.RequireToken("employee"), p.Delete)
}

func NewProductHandler(productUc usecase.ProductUsecase, authMiddleware middleware.AuthMiddleware, rg fiber.Router) *ProductHandler {
	return &ProductHandler{productUc: productUc, authMiddleware: authMiddleware, rg: rg}
}
