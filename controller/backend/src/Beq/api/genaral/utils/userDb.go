package utils

import (
	"Beq/api/genaral/model"
	"errors"
	"sync"
)

//UsersDB Initiating list database
type UsersDB map[int]model.UserInfo

//Instance Initiating list database
var instance UsersDB
var userIndex int

var once sync.Once

//GetUserStore Initiating list database
func GetUserStore() *UsersDB {
	once.Do(func() {
		instance = make(map[int]model.UserInfo)
		userIndex = 0
	})
	return &instance
}

//AddUser add user to data base
func (*UsersDB) AddUser(User model.UserInfo) error {
	if instance != nil {
		instance[userIndex] = User
		userIndex++
		return nil
	}
	return errors.New("No Data Base Initiate")

}

//FindUser add user to data base
func (*UsersDB) FindUser(Email string) (*model.UserInfo, error) {
	if instance != nil {
		for _, userInfo := range instance {
			if userInfo.Email == Email {
				return &userInfo, nil
			}
		}
		return nil, errors.New("User NOT FOUND! ")
	}
	return nil, errors.New("No Data Base Initiate")
}
