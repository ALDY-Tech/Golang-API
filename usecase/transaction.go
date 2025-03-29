package usecase

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
	"submission-project-enigma-laundry/utils"
)

type TransactionUseCase interface {
	Insert(transaction *model.Transaction) error
	FindById(id string) (model.TransactionDetail, error)
	FindAll(startDate string, endDate string, productName string) ([]model.TransactionDetail, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func (uc *transactionUsecase) Insert(transaction *model.Transaction) error {
	transaction.ID = utils.GenerateID()
	transaction.BillDetails[0].ID = utils.GenerateID()
	return uc.transactionRepo.Insert(transaction)
}

func (uc *transactionUsecase) FindById(id string) (model.TransactionDetail, error) {
	return uc.transactionRepo.FindById(id)
}

func (uc *transactionUsecase) FindAll(startDate string, endDate string, productName string) ([]model.TransactionDetail, error) {
	return uc.transactionRepo.FindAll(startDate, endDate, productName)
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository) TransactionUseCase {
	tu := new(transactionUsecase)
	tu.transactionRepo = transactionRepo
	return tu
}
