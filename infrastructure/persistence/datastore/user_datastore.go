package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userDatastore struct {
	db *gorm.DB
}

// NewUserDatastore : UserPersistenseを生成.
func NewUserDatastore(db *gorm.DB) repository.UserRepository {
	return &userDatastore{db}
}

func (uD userDatastore) FindAllUser() ([]*models.User, error) {
	users := []*models.User{}

	err := uD.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uD userDatastore) FindUserByID(id int64) (*models.User, error) {
	User := models.User{ID: id}
	err := uD.db.Take(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (uD userDatastore) FindUserByProviderID(provider, id string) (*models.User, error) {
	user := models.User{}
	switch provider {
	case "google":
		user.GoogleID = id
	}

	err := uD.db.Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uD userDatastore) InsertUser(user *models.User) error {
	return uD.db.Create(user).Error
}

func (uD userDatastore) UpdateUser(user *models.User) error {
	return uD.db.Updates(user).Error
}

func (uD userDatastore) DeleteUser(user *models.User) error {
	err := uD.db.Take(&user).Error
	if err != nil {
		return err
	}
	return uD.db.Delete(user).Error
}
