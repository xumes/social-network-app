package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	IsAuthRequired bool
}

func Config(router *mux.Router) *mux.Router {
	appRoutes := loginRoute
	appRoutes = append(appRoutes, userRoutes...)

	fmt.Println("Inside router config")

	for _, route := range appRoutes {
		fmt.Println(route.Uri)
		router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer((http.Dir("./assets/")))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
