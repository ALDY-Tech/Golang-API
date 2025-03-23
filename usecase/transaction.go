package usecase

import (
	"errors"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
	// "submission-project-enigma-laundry/utils"
)

type TransactionUseCase interface {
	Insert(transaction *model.Transaction) error
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func (uc *transactionUsecase) Insert(transaction *model.Transaction) error {
	if transaction.EmployeeID == "" || transaction.CustomerID == "" {
		return errors.New("employeeId and customerId are required")
	}

	// Insert transaksi ke repository
	err := uc.transactionRepo.Insert(transaction)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository) TransactionUseCase {
	tu := new(transactionUsecase)
	tu.transactionRepo = transactionRepo
	return tu
}
