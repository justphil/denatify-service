package controllers

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/justphil/denatify-service/models"
	"github.com/justphil/denatify-service/models/forms"
	"github.com/justphil/denatify-service/templates"
)

type UsersController struct {
	AppController
}

func (c *UsersController) NewUser() error {
	return templates.Layout(c.ResponseWriter, func() {
		templates.UserForm(c.ResponseWriter, "create", nil, nil)
	})
}

func (c *UsersController) CreateNewUser() error {
	userForm := new(forms.UserForm)
	errs := c.Bind(userForm)

	if errs.Len() > 0 {
		return c.renderUserForm("create", errors.New("An error occurred! (1)"), "", userForm)
	} else {
		u := &models.User{Fullname: userForm.Fullname, Email: userForm.Email, Password: userForm.Password}
		_, err := c.store.CreateUser(u)
		if err != nil {
			return c.renderUserForm("create", err, "", userForm)
		}

		return c.Redirect("/users", http.StatusSeeOther)
	}
}

func (c *UsersController) Users() error {
	users, err := c.store.GetUsers()
	if err != nil {
		return err
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.Users(c.ResponseWriter, users, err)
	})
}

func (c *UsersController) User() error {
	var u *models.User
	var err error

	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		return errors.New("The provided ID is invalid.")
	}

	if err == nil {
		u, err = c.store.GetUserById(userId)
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.User(c.ResponseWriter, u, err)
	})
}

func (c *UsersController) UpdateUser() error {
	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		return errors.New("The provided ID is invalid.")
	}

	u, err := c.store.GetUserById(userId)
	if err != nil {
		return err
	}

	data := map[string]string{
		"userId":   u.Id.Hex(),
		"fullname": u.Fullname,
		"email":    u.Email,
		"password": u.Password,
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.UserForm(c.ResponseWriter, "update", err, data)
	})
}

func (c *UsersController) PerformUserUpdate() error {
	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		return errors.New("The provided ID is invalid.")
	}

	userForm := new(forms.UserForm)
	errs := c.Bind(userForm)
	var err error

	if errs.Len() > 0 {
		err = errors.New("An error occurred! (1)")
		return c.renderUserForm("update", err, userId, userForm)
	} else {
		u := &models.User{
			Id:       bson.ObjectIdHex(userId),
			Fullname: userForm.Fullname,
			Email:    userForm.Email,
			Password: userForm.Password,
		}

		err = c.store.UpdateUser(u)
		if err != nil {
			return c.renderUserForm("update", err, userId, userForm)
		}

		return c.Redirect("/users/"+userId, http.StatusSeeOther)
	}
}

func (c *UsersController) DeleteUser() error {
	var u *models.User
	var err error

	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		return errors.New("The provided ID is invalid.")
	}

	if err == nil {
		u, err = c.store.GetUserById(userId)
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.DeleteUser(c.ResponseWriter, u, err)
	})
}

func (c *UsersController) PerformUserDeletion() error {
	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		return errors.New("The provided ID is invalid.")
	}

	err := c.store.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return c.Redirect("/users", http.StatusSeeOther)
}

func (c *UsersController) renderUserForm(action string, err error, userId string, form *forms.UserForm) error {
	data := map[string]string{
		"userId":   userId,
		"fullname": form.Fullname,
		"email":    form.Email,
		"password": form.Password,
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.UserForm(c.ResponseWriter, action, err, data)
	})
}
