package app

import "github.com/aashish-sport/bookstore_users-apii/controllers"

func mapurls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", controllers.SearchUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)
	router.PATCH("/users/:user_id", controllers.UpdateUser)
	router.DELETE("/users/:user_id", controllers.DeleteUser)
	router.GET("/internal/users/search", controllers.Search)
	router.POST("/users", controllers.CreateUser)
}
