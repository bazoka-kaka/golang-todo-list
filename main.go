package main

import (
	"fmt"
	"net/http"

	"github.com/bazoka-kaka/golang-todo-list/controllers"
	"github.com/bazoka-kaka/golang-todo-list/middlewares"
)

func main() {
	http.Handle("/register", middlewares.AllowOnlyPOST(http.HandlerFunc(controllers.HandleRegister)))
	http.Handle("/login", middlewares.AllowOnlyPOST(http.HandlerFunc(controllers.HandleLogin)))
	http.Handle("/logout", middlewares.AllowOnlyGET(http.HandlerFunc(controllers.HandleLogout)))

	// api
	http.Handle("/users", middlewares.Authenticate(middlewares.AllowOnlyGET(http.HandlerFunc(controllers.GetAllUsers))))

	// listen to server
	fmt.Println("server running on port 3500")
	http.ListenAndServe(":3500", nil)
}