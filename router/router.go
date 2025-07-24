package router

import (
	"homeassignment/controllers"
	"net/http"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

var routes = []Route{
	{Path: "/accounts/hello", Handler: controllers.HelloWorld, Method: "GET"},
	{Path: "/accounts", Handler: controllers.GetAccountInfo, Method: "POST"},
	{Path: "/accounts/{account_id}", Handler: controllers.GetAccountInfo, Method: "GET"},
	{Path: "/transactions", Handler: controllers.GetAccountInfo, Method: "POST"},
}

func RegisterRoutes() {
	for _, route := range routes {
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != route.Method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			route.Handler(w, r)
		})
	}
}
