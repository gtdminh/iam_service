package service

import (
	"auth_service/config"
	"auth_service/model"

	"github.com/globalsign/mgo/bson"
)

/*
GetUserByID returns user by mongo id
*/
func GetUserByID(id string) *model.User {
	db := config.GetMongoDb()

	defer db.Session.Close()

	var user *model.User
	if db != nil {
		err := db.C("users").Find(bson.M{id: id}).One(user)
		if err != nil {
			panic(err)
		}
	}

	return user
}

/*
GetUserByEmail finds user by email
*/
func GetUserByEmail(email string) *model.User {
	db := config.GetMongoDb()

	defer db.Session.Close()

	var user *model.User
	if db != nil {
		err := db.C("users").Find(bson.M{email: email}).One(user)
		if err != nil {
			panic(err)
		}
	}

	return user
}

/*
InsertUser adds user to db
*/
func InsertUser(u *model.User) bool {
	db := config.GetMongoDb()

	defer db.Session.Close()

	if db != nil {

		err := db.C("users").Insert(u)
		if err != nil {
			return false
		}

		return true
	}

	return false
}
