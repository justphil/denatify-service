package middlewares

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
)

func FooMiddleware() negroni.Handler {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		context.Set(r, "foo", "bar")
		next(rw, r)
	})
}
