package controllers

import "TestCodelite/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	//s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/forgotpassword", middlewares.SetMiddlewareJSON(s.ForgotPassword)).Methods("POST")
	s.Router.HandleFunc("/resetpassword",middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePassword))).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/article", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateArticle))).Methods("POST")
	s.Router.HandleFunc("/articles", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetArticles))).Methods("GET")
	s.Router.HandleFunc("/article/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetArticle))).Methods("GET")
	s.Router.HandleFunc("/article/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateArticle))).Methods("PUT")
	s.Router.HandleFunc("/article/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteArticle)).Methods("DELETE")
}