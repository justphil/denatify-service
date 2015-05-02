package forms

import (
	"github.com/mholt/binding"
)

type UserForm struct {
	Fullname string
	Email    string
	Password string
}

func (f *UserForm) FieldMap() binding.FieldMap {
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
