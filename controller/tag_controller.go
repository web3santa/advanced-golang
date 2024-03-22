package controller

import (
	"gin-gorm/data/request"
	"gin-gorm/data/response"
	"gin-gorm/helper"
	"gin-gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagReponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagReponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagReponse := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagReponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
