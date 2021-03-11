package controllers

import (
	"fmt"
	"net/http"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/models"
	"github.com/GuilhermeFirmiano/street-market/pkg/services"
	"github.com/gin-gonic/gin"
)

//StreetMarketController ...
type StreetMarketController struct {
	service *services.StreetMarketService
}

//NewStreetMarketController ...
func NewStreetMarketController(service *services.StreetMarketService) *StreetMarketController {
	return &StreetMarketController{
		service: service,
	}
}

//RegisterRoutes ...
func (controller *StreetMarketController) RegisterRoutes(router *gin.RouterGroup) {
	streetMarket := router.Group("street-market")
	{
		streetMarket.POST("", controller.post)
		streetMarket.PUT(":id", controller.put)
		streetMarket.GET(":id", controller.getByID)
		streetMarket.GET("", controller.filter)
		streetMarket.DELETE(":id", controller.deleteByID)
		streetMarket.DELETE("", controller.deleteByFilter)
	}
}

//post ...
func (controller *StreetMarketController) post(ctx *gin.Context) {
	r := new(models.StreetMarketPostRequest)
	err := ctx.ShouldBindJSON(r)

	if err != nil {
		grok.BindingError(ctx, err)
		return
	}

	response, err := controller.service.Create(ctx, r)

	if err != nil {
		grok.ResolveError(ctx, err)
		return
	}

	ctx.Writer.Header().Set("Location", fmt.Sprintf("%s/%s", ctx.Request.URL.Path, response.ID))

	ctx.JSON(http.StatusCreated, response)
}

//put ...
func (controller *StreetMarketController) put(ctx *gin.Context) {
	id := ctx.Param("id")

	r := new(models.StreetMarketPutRequest)
	err := ctx.ShouldBindJSON(r)

	if err != nil {
		grok.BindingError(ctx, err)
		return
	}

	response, err := controller.service.Update(ctx, id, r)

	if err != nil {
		grok.ResolveError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//getByID ...
func (controller *StreetMarketController) getByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := controller.service.GetByID(ctx, id)

	if err != nil {
		grok.ResolveError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//filter
func (controller *StreetMarketController) filter(ctx *gin.Context) {
	params := new(models.FilterStreetMarket)

	if err := ctx.ShouldBindQuery(params); err != nil {
		grok.BindingError(ctx, err)
		return
	}

	result, err := controller.service.Filter(ctx, params)

	if err != nil {
		grok.ResolveError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (controller *StreetMarketController) deleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := controller.service.DeleteByID(ctx, id)

	if err != nil {
		grok.ResolveError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controller *StreetMarketController) deleteByFilter(ctx *gin.Context) {
	filter := new(models.FilterStreetMarketDelete)

	if err := ctx.ShouldBindQuery(filter); err != nil {
		grok.BindingError(ctx, err)
		return
	}

	err := controller.service.DeleteByFilter(ctx, filter)

	if err != nil {
		grok.ResolveError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
