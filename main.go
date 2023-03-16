package main

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"golang-api/controller"
	"golang-api/helper"
	"golang-api/repository"
	"golang-api/service"
	"golang-api/app"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	server := http.Server {
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}