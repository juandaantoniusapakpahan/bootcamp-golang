package main

import (
	"belajar-gin/controller"
	"belajar-gin/repository"
	"belajar-gin/routes"
	"belajar-gin/service"
)

func main() {
	newCategoryRepo := repository.NewCategory()
	newCategoryService := service.NewCategoryServicelement(newCategoryRepo)
	newCategoryController := controller.NewCategoryControllerImplement(newCategoryService)

	route := routes.NewRoute(newCategoryController)

	route.Run()
}

// Initiate [i]
// Request Method [-]
// Routing [-]
// Query parameter [-]
// Path parameter [-]
// Request Body [-]
// Cookie [-]
// Header [-]
// Serve File
// File Server [-]
// Download [-]
// Upload [-]
// Redirect [-]
// Gimana cara implementasi middleware (template middleware, cara pakai, grouping) [-]
