package repository

import (
	"rhmn-coffe/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(id string) error
	FindById(id string) (entity.Product, error)
	FindByProductName(productName string) (entity.Product, error)
	FindAll() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Create(product entity.Product) (entity.Product, error) {
	err := p.db.Create(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *productRepository) Update(product entity.Product) (entity.Product, error) {
	if err := p.db.Save(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *productRepository) Delete(id string) error {
	err := p.db.Where("product_id = ?", id).Delete(&entity.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) FindById(id string) (entity.Product, error) {
	var product entity.Product
	err := p.db.Where("product_id = ?", id).First(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *productRepository) FindByProductName(productName string) (entity.Product, error) {
	var product entity.Product
	err := p.db.Where("product_name = ?", productName).First(&product).Error
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
