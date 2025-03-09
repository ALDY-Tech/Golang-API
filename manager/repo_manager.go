package manager

import "submission-project-enigma-laundry/repository"

type RepositoryManager interface {
	CustomerRepository() repository.CustomerRepository
}

type repoManager struct {
	infra InfraManager
}

func (i *repoManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(i.infra.SqlDB())
}

func NewRepoManager(infra InfraManager) RepositoryManager {
	return &repoManager{
		infra: infra,
	}
}