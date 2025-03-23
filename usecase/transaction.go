package usecase

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
	"submission-project-enigma-laundry/utils"
)

type TransactionUseCase interface {
	Insert(transaction *model.Transaction) error
	findById(id string) (*model.Transaction, error)
	// FindAll(page int, totalRows int) ([]model.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func (uc *transactionUsecase) Insert(transaction *model.Transaction) error {
	transaction.ID = utils.GenerateID()
	transaction.BillDetails[0].ID = utils.GenerateID()
	return uc.transactionRepo.Insert(transaction)
}

func (uc *transactionUsecase) findById(id string) (*model.Transaction, error) {
	return uc.transactionRepo.FindById(id)
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository) TransactionUseCase {
	tu := new(transactionUsecase)
	tu.transactionRepo = transactionRepo
	return tu
}
