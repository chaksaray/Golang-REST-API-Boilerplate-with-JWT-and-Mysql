package startup

import (
	"net/http"

	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/controllers"
	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/middlewares"

	"github.com/jinzhu/gorm"
)


func (a *App) InitializeRoutes() {
	// Get default routing
	a.Get("/", a.HandleRequest(controllers.Home))

	// Routing for auth
	a.Post("/login", a.HandleRequest(controllers.Login))

	// Routing for handling the users
	a.Get("/users", a.HandleRequest(controllers.GetAllUsers))
	a.Post("/users", a.HandleRequest(controllers.CreateUser))
	a.Get("/users/{id:[0-9]+}", a.HandleRequest(controllers.GetUserById))
	a.Put("/users/{id:[0-9]+}", a.HandleRequest(controllers.UpdateUser))
	a.Delete("/users/{id:[0-9]+}", a.HandleRequest(controllers.DeleteUser))

	// Routing for handling the posts
	a.Get("/posts", a.HandleRequest(controllers.GetAllPosts))
	a.Post("/posts", a.HandleRequest(controllers.CreatePost))
	a.Get("/posts/{id:[0-9]+}", a.HandleRequest(controllers.GetPostById))
	a.Put("/posts/{id:[0-9]+}", middlewares.SetMiddlewareAuthentication(a.HandleRequest(controllers.UpdatePost)))
	a.Delete("/posts/{id:[0-9]+}", middlewares.SetMiddlewareAuthentication(a.HandleRequest(controllers.DeletePost)))
}

// wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}