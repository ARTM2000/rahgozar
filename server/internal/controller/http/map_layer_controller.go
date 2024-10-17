package http

import (
	"github.com/ARTM2000/rahgozar/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type mapLayerController struct {
	mapLayersService port.IMapLayersService
}

func NewMapLayerController(mapLayersService port.IMapLayersService) Controller {
	return &mapLayerController{
		mapLayersService: mapLayersService,
	}
}

func (mlc *mapLayerController) InitRoutes(app *fiber.App) {
	app.Get("/map-layers/v1/layers-list/", mlc.getLayersList)
	app.Get("/map-layers/v1/layer-data/", mlc.getLayerData)
}

func (mlc *mapLayerController) getLayersList(ctx *fiber.Ctx) error {
	activeLayers, _ := mlc.mapLayersService.GetActiveMapLayersList()

	return ctx.Status(fiber.StatusOK).JSON(FormatResponse(ctx, ResponseData{
		IsError: false,
		Data: map[string]interface{}{
			"layers": activeLayers,
		},
	}))
}

func (mlc *mapLayerController) getLayerData(ctx *fiber.Ctx) error {
	queries := struct {
		LayerName string `query:"layer_name"`
	}{}
	if err := ctx.QueryParser(&queries); err != nil {
		return &fiber.Error{Message: "bad request", Code: fiber.StatusBadRequest}
	}

	layerData, _ := mlc.mapLayersService.GetActiveMapLayerByName(queries.LayerName)

	return ctx.Status(fiber.StatusOK).JSON(FormatResponse(ctx, ResponseData{
		Data: map[string]interface{}{
			"layer_data": layerData,
		},
	}))
}
