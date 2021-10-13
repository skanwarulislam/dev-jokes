package router

import (
	"github.com/gorilla/pat"
	"makeajoke/handlers"
)

func SetupRouter() *pat.Router {
	router := pat.New()
	router.Get("/auth/{provider}/callback", handlers.DoCallback)
	router.Get("/auth/{provider}", handlers.DoAuth)
	router.Get("/logout/{provider}", handlers.DoLogout)
	router.Get("/", handlers.ShowIndex)

	return router
}
