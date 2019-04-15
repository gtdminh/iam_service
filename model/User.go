package model

import (
	"time"
)

/*
User table
*/
type User struct {
	ID       string `bson:"id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`

	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Address   string `bson:"address"`
	City      string `bson:"city"`
	Country   string `bson:"country"`
	Zip       string `bson:"zip"`

	CreatedOn          time.Time `bson:"createdon"`
	ModifiedOn         time.Time `bson:"modifiedon"`
	EmailConfirmed     bool      `bson:"email_confirmed"`
	PasswordReset      bool      `bson:"password_reset"`
	EmailConfirmSecret string    `bson:"email_confirm_secret"`
}

/*
Role table
*/
type Role struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}

/*
Permission table
*/
type Permission struct {
	ID           string `bson:"id"`
	Name         string `bson:"name"`
	RoutePattern string `bson:"pattern"`
}

func NewUser(email, password, firstName, lastName, address, city, country, zip string) User {
	return User{
		Email:          email,
		Password:       password,
		FirstName:      firstName,
		LastName:       lastName,
		Address:        address,
		City:           city,
		Country:        country,
		Zip:            zip,
		CreatedOn:      time.Now(),
		ModifiedOn:     time.Now(),
		EmailConfirmed: false,
		PasswordReset:  true,
	}
}
