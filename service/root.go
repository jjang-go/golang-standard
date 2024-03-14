package service

import (
	"sync"

	"github.com/jjang-go/golang-standard/repository"
)

// Network, Repository의 다리역할

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	//repository
	repository *repository.Repository
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
	})

	return serviceInstance
}
