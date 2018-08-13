package repository

import (
	"github.com/febriliando/rest-api-go/src/modules/profile/model"
)

// Profile Repository
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindById(string) (*model.Profile, error)
	FindByEmail(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
