package controllers

import "github.com/tee-stark/go-blog/api/middlewares"

func (server *Server) InitializeRoutes() {
	// for the index route
	server.Router.HandleFunc("/", middlewares.ResToJSONMiddleware(server.Home)).Methods("GET")
	// login route
	server.Router.HandleFunc("/login", middlewares.ResToJSONMiddleware(server.Login)).Methods("POST")
	// user routes
	server.Router.HandleFunc("/users", middlewares.ResToJSONMiddleware(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.ResToJSONMiddleware(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.ResToJSONMiddleware(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.ResToJSONMiddleware(middlewares.AuthMiddleware(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.ResToJSONMiddleware(middlewares.AuthMiddleware(server.DeleteUser))).Methods("DELETE")

	// post routes
	server.Router.HandleFunc("/posts", middlewares.ResToJSONMiddleware(server.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/posts", middlewares.ResToJSONMiddleware(server.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/posts/{id}", middlewares.ResToJSONMiddleware(server.GetPost)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middlewares.ResToJSONMiddleware(middlewares.AuthMiddleware(server.UpdatePost))).Methods("PUT")
	server.Router.HandleFunc("/posts/{id}", middlewares.AuthMiddleware(server.DeletePost)).Methods("DELETE")
}
