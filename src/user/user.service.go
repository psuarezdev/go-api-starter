package user

import "github.com/psuarezdev/go-api-starter/src/database"

func GetById(id uint) (*User, error) {
	db := database.GetConnection()

	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetByUsername(username string) (*User, error) {
	db := database.GetConnection()

	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func Create(user *User) error {
	db := database.GetConnection()

	if err := db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
