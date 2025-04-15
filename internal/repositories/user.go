package repositories

import (
	"go-rest-api/internal/models"
)

type UserRepository struct {
	Store storeInterface
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.Store.DB().QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES (?, ?) RETURNING id", 
		user.Email, 
		user.EncPassword,
	).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	return nil, nil
}
