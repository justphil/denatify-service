package controllers

import (
	"github.com/justphil/denatify-service/templates"
)

type LandingPageController struct {
	AppController
}

func (c *LandingPageController) Index() error {
	return templates.Layout(c.ResponseWriter, func() {
		templates.LandingPage(c.ResponseWriter)
	})
}
