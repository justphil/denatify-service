package controllers

import (
	"fmt"
	//"net/http"

	"github.com/justphil/denatify-service/templates"
)

type LandingPageController struct {
	AppController
}

func (c *LandingPageController) Index() error {
	fmt.Printf("Name: %s\n", c.Params.ByName("name"))
	//c.HTML(http.StatusOK, "index", nil)
	return templates.Layout(c.ResponseWriter, func() {
		templates.LandingPage(c.ResponseWriter)
	})
}
