package controllers

import (
	"os"
	"strings"

	"github.com/unrolled/render"

	"github.com/justphil/denatify-service/models/store"
)

const (
	DEV_MODE  = "DEV_MODE"
	MONGO_URL = "MONGO_URL"
)

var IS_DEV_MODE = getDevMode()
var renderInstance *render.Render
var storeInstance store.Store

func init() {
	renderInstance = render.New(render.Options{
		IndentJSON:    IS_DEV_MODE, // if true, output human readable JSON
		IsDevelopment: IS_DEV_MODE, // if true, recompile templates on every request
	})

	storeInstance = store.NewMongoStore(getMongoUrl())
}

func getDevMode() bool {
	devMode := false

	if devModeStr := os.Getenv(DEV_MODE); len(devModeStr) > 0 && strings.ToLower(devModeStr) == "true" {
		devMode = true
	}

	return devMode
}

func getMongoUrl() string {
	url := os.Getenv(MONGO_URL)
	if url == "" {
		url = "127.0.0.1"
	}

	return url
}
