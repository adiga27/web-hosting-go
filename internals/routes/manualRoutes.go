package routes

import (
	"github.com/adiga27/web-hosting-go/internals/controller/manual"
	"github.com/gofiber/fiber/v2"
)

func ManualRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1/manual")
	v1.Post("/createApp", manual.CreateAppBranch)
	v1.Post("/deployApp/:id", manual.DeployApp)
	v1.Delete("/deleteApp/:id", manual.DeleteAppBranch)
	v1.Get("/getApp", manual.GetAllApp)
	v1.Get("/getStatus/:id", manual.GetAppStatus)
	v1.Patch("/updateApp/:id", manual.UpdateAppBranch)
}
