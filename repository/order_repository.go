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
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		if len(order.Order_items) > 0 {
			for i := range order.Order_items {
				order.Order_items[i].Order_id = order.Order_id
			}
			if err := tx.Create(&order.Order_items).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) Update(order entity.Order) (entity.Order, error) {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&order).Error; err != nil {
			return err
		}
		if len(order.Order_items) > 0 {
			if err := tx.Where("order_id = ?", order.Order_id).Delete(&entity.OrderItem{}).Error; err != nil {
				return err
			}
			for i := range order.Order_items {
				order.Order_items[i].Order_id = order.Order_id
			}
			if err := tx.Create(&order.Order_items).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) Delete(orderId string) error {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", orderId).Delete(&entity.OrderItem{}).Error; err != nil {
			return err
		}
		if err := tx.Where("order_id = ?", orderId).Delete(&entity.Order{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (o *orderRepository) FindById(orderId string) (entity.Order, error) {
	var order entity.Order
	err := o.db.Preload("Order_items").
		Where("order_id = ?", orderId).
		First(&order).Error
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (o *orderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := o.db.Preload("Order_items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
