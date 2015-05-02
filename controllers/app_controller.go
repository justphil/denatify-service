package controllers

import (
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justphil/httprouter-controller"
	"github.com/mholt/binding"
	"github.com/unrolled/render"

	"github.com/justphil/denatify-service/models/store"
)

const (
	IS_DEV_MODE = true
)

var renderInstance *render.Render
var storeInstance store.Store

func init() {
	renderInstance = render.New(render.Options{
		Layout:        "_layout",
		IndentJSON:    IS_DEV_MODE, // if true, output human readable JSON
		IsDevelopment: IS_DEV_MODE, // if true, recompile templates on every request
	})

	storeInstance = store.NewMongoStore(getMongoUrl())
}

func getMongoUrl() string {
	url := os.Getenv("MONGO_URL")
	if url == "" {
		url = "127.0.0.1"
	}

	return url
}

type AppController struct {
	controller.Base
	render *render.Render // we use render as a field (instead of embedding it) because we want to reuse the global render.Render instance
	store  store.Store
}

// lifecycle methods
func (ac *AppController) Init(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	// reuse the same render.Render instance for every request
	// assumption: it is thread safe.
	ac.render = renderInstance
	ac.store = storeInstance
	return ac.Base.Init(rw, r, params)
}

// helper methods
func (ac *AppController) HTML(status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) {
	ac.render.HTML(ac.ResponseWriter, status, name, binding, htmlOpt...)
}

func (ac *AppController) Bind(formStruct binding.FieldMapper) binding.Errors {
	return binding.Bind(ac.Request, formStruct)
}

func (ac *AppController) Redirect(url string, code int) error {
	http.Redirect(ac.ResponseWriter, ac.Request, url, code)
	return nil
}

func (ac *AppController) ContextGet(key interface{}) interface{} {
	return context.Get(ac.Request, key)
}

func (ac *AppController) ContextSet(key, value interface{}) {
	context.Set(ac.Request, key, value)
}
