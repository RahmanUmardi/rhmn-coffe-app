package repository

import (
	"rhmn-coffe/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	Update(order entity.Order) (entity.Order, error)
	Delete(orderId string) error
	FindById(orderId string) (entity.Order, error)
	FindAll() ([]entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) Create(order entity.Order) (entity.Order, error) {
	err := o.db.Create(&order).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) Update(order entity.Order) (entity.Order, error) {
	err := o.db.Save(&order).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) Delete(orderId string) error {
	err := o.db.Where("order_id = ?", orderId).Delete(&entity.Order{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) FindById(orderId string) (entity.Order, error) {
	var order entity.Order
	err := o.db.Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := o.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
