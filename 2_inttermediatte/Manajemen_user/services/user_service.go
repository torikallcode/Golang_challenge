package services

import (
	"golang_challenge/2_inttermediatte/Manajemen_user/models"
)

var index = 1
var users []models.User

func AddUser(name, email string) models.User {
	user := models.User{
		ID:    index,
		Name:  name,
		Email: email,
	}
	index++
	users = append(users, user)
	return user
}

func GetAllUser() []models.User {
	return users
}

func EditEmail(id int, newEmail string) bool {
	for i := range users {
		if users[i].ID == id {
			users[i].UpdateEmail(newEmail)
			return true
		}
	}
	return false
}

func DeleteUser(id int) bool {
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
