package model

import (
	"time"
)

type Profile struct {
	ID        string    `bson:"id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"creatd_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Profiles []Profile

// helpers just check the valid pass in db
func (p *Profile) IsValidPassword(password string) bool {
	return p.Password == password
}
