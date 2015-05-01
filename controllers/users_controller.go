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
		templates.UserForm(c.ResponseWriter, "create", "", nil)
	})
}

func (c *UsersController) CreateNewUser() error {
	createNewUserForm := new(forms.CreateNewUserForm)
	errs := c.Bind(createNewUserForm)
	var errMsg string

	if errs.Len() > 0 {
		errMsg = "An error occurred! (1)"
		return c.renderUserForm("create", errMsg, "", createNewUserForm)
	} else {
		u := &models.User{Fullname: createNewUserForm.Fullname, Email: createNewUserForm.Email, Password: createNewUserForm.Password}
		_, err := c.store.CreateUser(u)
		if err != nil {
			errMsg = "An error occurred! (2)"
			return c.renderUserForm("create", errMsg, "", createNewUserForm)
		}

		return c.Redirect("/users", http.StatusSeeOther)
	}
}

func (c *UsersController) Users() error {
	var errMsg string
	users, err := c.store.GetUsers()
	if err != nil {
		errMsg = err.Error()
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.Users(c.ResponseWriter, users, errMsg)
	})
}

func (c *UsersController) User() error {
	var u *models.User
	var err error

	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		err = errors.New("There is no user with the specified ID.")
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

	var errMsg string
	var data map[string]string
	u, err := c.store.GetUserById(userId)
	if err != nil {
		errMsg = err.Error()
	} else {
		data = map[string]string{
			"userId":   u.Id.Hex(),
			"fullname": u.Fullname,
			"email":    u.Email,
			"password": u.Password,
		}
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.UserForm(c.ResponseWriter, "update", errMsg, data)
	})
}

func (c *UsersController) PerformUserUpdate() error {
	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		return errors.New("The provided ID is invalid.")
	}

	updateUserForm := new(forms.CreateNewUserForm)
	errs := c.Bind(updateUserForm)
	var errMsg string

	if errs.Len() > 0 {
		errMsg = "An error occurred! (1)"
		return c.renderUserForm("update", errMsg, userId, updateUserForm)
	} else {
		u := &models.User{
			Id:       bson.ObjectIdHex(userId),
			Fullname: updateUserForm.Fullname,
			Email:    updateUserForm.Email,
			Password: updateUserForm.Password,
		}

		err := c.store.UpdateUser(u)
		if err != nil {
			errMsg = "An error occurred! (2)"
			return c.renderUserForm("update", errMsg, userId, updateUserForm)
		}

		return c.Redirect("/users/"+userId, http.StatusSeeOther)
	}
}

func (c *UsersController) DeleteUser() error {
	var u *models.User
	var err error

	userId := c.Params.ByName("userId")
	if !bson.IsObjectIdHex(userId) {
		err = errors.New("There is no user with the specified ID.")
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
		return errors.New("There is no user with the specified ID.")
	}

	err := c.store.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return c.Redirect("/users", http.StatusSeeOther)
}

func (c *UsersController) renderUserForm(action string, errMsg string, userId string, form *forms.CreateNewUserForm) error {
	data := map[string]string{
		"userId":   userId,
		"fullname": form.Fullname,
		"email":    form.Email,
		"password": form.Password,
	}

	return templates.Layout(c.ResponseWriter, func() {
		templates.UserForm(c.ResponseWriter, action, errMsg, data)
	})
}
