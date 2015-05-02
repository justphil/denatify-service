package controllers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justphil/httprouter-controller"
	"github.com/mholt/binding"
	"github.com/unrolled/render"

	"github.com/justphil/denatify-service/models/store"
)

type AppController struct {
	controller.Base
	store store.Store

	// we use render as a field (instead of embedding it) because we want to reuse the global render.Render instance
	render *render.Render
}

// lifecycle methods
func (ac *AppController) Init(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	// reuse the same render.Render instance for every request
	// assumption: it is thread safe.
	ac.render = renderInstance
	ac.store = storeInstance
	return ac.Base.Init(rw, r, params)
}

// helper methods for handling request data
func (ac *AppController) Bind(formStruct binding.FieldMapper) binding.Errors {
	return binding.Bind(ac.Request, formStruct)
}

// helper methods for creating responses
func (ac *AppController) Redirect(url string, code int) error {
	http.Redirect(ac.ResponseWriter, ac.Request, url, code)
	return nil
}

// helper methods for context access
func (ac *AppController) ContextGet(key interface{}) interface{} {
	return context.Get(ac.Request, key)
}

func (ac *AppController) ContextSet(key, value interface{}) {
	context.Set(ac.Request, key, value)
}
