package usecase

import (
	"fmt"
	"rhmn-coffe/entity"
	"rhmn-coffe/repository"
	"strconv"
	"strings"
)

type ProductUsecase interface {
	Create(product entity.Product) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	FindById(id string) (entity.Product, error)
	FindByProductName(productName string) (entity.Product, error)
	Update(id string, input entity.UpdateProduct) (entity.Product, error)
	Delete(id string) error
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func (p *productUsecase) Create(product entity.Product) (entity.Product, error) {
	priceStr := strconv.FormatFloat(product.Price, 'f', -1, 64)
	if strings.TrimSpace(product.Product_name) == "" || strings.TrimSpace(priceStr) == "" {
		return entity.Product{}, fmt.Errorf("productName and price can't be empty")
	}

	exitProduct, _ := p.productRepository.FindByProductName(product.Product_name)
	if exitProduct.Product_name != "" {
		return entity.Product{}, fmt.Errorf("productName already exist")
	}

	return p.productRepository.Create(product)
}

func (p *productUsecase) FindAll() ([]entity.Product, error) {
	return p.productRepository.FindAll()
}

func (p *productUsecase) FindById(id string) (entity.Product, error) {
	return p.productRepository.FindById(id)
}

func (p *productUsecase) FindByProductName(productName string) (entity.Product, error) {
	return p.productRepository.FindByProductName(productName)
}

func (p *productUsecase) Update(id string, input entity.UpdateProduct) (entity.Product, error) {
	product, err := p.productRepository.FindById(id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("product not found: %v", err)
	}

	if strings.TrimSpace(input.Product_name) != "" {
		product.Product_name = input.Product_name
	}
	priceStr := strconv.FormatFloat(product.Price, 'f', -1, 64)
	if strings.TrimSpace(priceStr) != "" {
		product.Price = input.Price
	}

	UpdateProduct, err := p.productRepository.Update(product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to update product: %v", err)
	}

	return UpdateProduct, nil
}

func (p *productUsecase) Delete(id string) error {
	_, err := p.FindById(id)
	if err != nil {
		return err
	}

	err = p.productRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete product : %v", err)
	}

	return nil
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepository: productRepository}
}
