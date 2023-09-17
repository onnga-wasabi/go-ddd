package repository

import (
	"context"

	"github.com/onnga-wasabi/go-ddd/sample/domain/model/user"
	"github.com/onnga-wasabi/go-ddd/sample/infrastructure/xo/persistence/data_model"
)

type userRepositoryXO struct {
	db data_model.DB
}

func NewUserRepositoryXO(db data_model.DB) user.UserRepository {
	return userRepositoryXO{db: db}
}

func (r userRepositoryXO) Save(u user.User) error {
	dataModelUser := data_model.User{
		ID:   string(u.ID()),
		Name: u.Name(),
	}
	return dataModelUser.Save(context.Background(), r.db)
}
