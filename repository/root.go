package repository

import (
	"sync"

	"github.com/jjang-go/golang-standard/types"
)

// 3티어 아키텍쳐
// 세션정의

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	UserMap []*types.User
	//repository
	//db mongo.database
}

func Newrepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{}
	})

	return repositoryInstance
}
