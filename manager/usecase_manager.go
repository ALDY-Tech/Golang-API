package manager

import "submission-project-enigma-laundry/usecase"

type UseCaseManager interface {
    CustomerUseCase() usecase.CustomerUseCase
    EmployeeUseCase() usecase.EmployeeUseCase
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

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
    return &useCaseManager{repo: repo}
}