package main

import (
	"github.com/fahrilhadi/golang-restful-api/app"
	"github.com/fahrilhadi/golang-restful-api/controller"
	"github.com/fahrilhadi/golang-restful-api/repository"
	"github.com/fahrilhadi/golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main()  {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.POST("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}