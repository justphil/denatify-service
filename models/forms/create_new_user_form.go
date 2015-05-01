package forms

import (
	"github.com/mholt/binding"
)

type CreateNewUserForm struct {
	Fullname string
	Email    string
	Password string
}

func (f *CreateNewUserForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&f.Fullname: binding.Field{
			Form:     "fullname",
			Required: true,
		},
		&f.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&f.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}
