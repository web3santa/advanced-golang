package main

import (
	"gin-gorm/config"
	"gin-gorm/controller"
	"gin-gorm/helper"
	"gin-gorm/model"
	"gin-gorm/repository"
	"gin-gorm/router"
	"gin-gorm/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repository.NewTagsRepositoryImpl(db)

	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
