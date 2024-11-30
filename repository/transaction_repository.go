package repository

import (
	"rhmn-coffe/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction entity.Transaction) (entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
	FindByOrderId(orderId string) ([]entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
	Delete(transactionId string) error
}

type transactionRepository struct {
	db *gorm.DB
}

func (t *transactionRepository) Create(transaction entity.Transaction) (entity.Transaction, error) {
	err := t.db.Create(&transaction).Error
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (t *transactionRepository) FindAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := t.db.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) FindByOrderId(orderId string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := t.db.Where("order_id = ?", orderId).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) Update(transaction entity.Transaction) (entity.Transaction, error) {
	err := t.db.Save(&transaction).Error
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (t *transactionRepository) Delete(transactionId string) error {
	err := t.db.Where("transaction_id = ?", transactionId).Delete(&entity.Transaction{}).Error
	return err
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}
