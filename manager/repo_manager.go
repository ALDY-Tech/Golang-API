package manager

import "submission-project-enigma-laundry/repository"

type RepositoryManager interface {
	CustomerRepository() repository.CustomerRepository
	EmployeeRepository() repository.EmployeeRepository
	ProductRepository() repository.ProductRepository
}

type repoManager struct {
	infra InfraManager
}

func (i *repoManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(i.infra.SqlDB())
}

func (i *repoManager) EmployeeRepository() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(i.infra.SqlDB())
}

func (i *repoManager) ProductRepository() repository.ProductRepository {
	return repository.NewProductRepository(i.infra.SqlDB())
}

func NewRepoManager(infra InfraManager) RepositoryManager {
	return &repoManager{
		infra: infra,
	}
}