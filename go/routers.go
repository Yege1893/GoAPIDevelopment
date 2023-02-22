/*
 * Todo app OAS
 *
 * OpenApi specification for a todo application
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateTodo",
		strings.ToUpper("Post"),
		"/todos",
		CreateTodo,
	},

	Route{
		"DeleteTodo",
		strings.ToUpper("Delete"),
		"/todos/{todoId}",
		DeleteTodo,
	},

	Route{
		"GetTodos",
		strings.ToUpper("Get"),
		"/todos",
		GetTodos,
	},

	Route{
		"GetTodosWithStatus",
		strings.ToUpper("Get"),
		"/todo/status/{status}",
		GetTodosWithStatus,
	},

	Route{
		"UpdateTodo",
		strings.ToUpper("Put"),
		"/todos/{todoId}",
		UpdateTodo,
	},
}