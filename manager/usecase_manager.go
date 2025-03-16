package manager

import "submission-project-enigma-laundry/usecase"

type UseCaseManager interface {
    CustomerUseCase() usecase.CustomerUseCase
    EmployeeUseCase() usecase.EmployeeUseCase
	ProductUsecase() usecase.ProductUseCase
}

type useCaseManager struct {
    repo RepositoryManager
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
    return usecase.NewCustomerUseCase(u.repo.CustomerRepository())
}

func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
    return usecase.NewEmployeeUseCase(u.repo.EmployeeRepository())
}

func (u *useCaseManager) ProductUsecase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repo.ProductRepository())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
    return &useCaseManager{repo: repo}
}