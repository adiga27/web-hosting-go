package manual

import (
	"time"

	"github.com/adiga27/web-hosting-go/internals/model"
	"github.com/adiga27/web-hosting-go/pkg"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateAppBranch(ctx *fiber.Ctx) error {
	app := new(model.App)
	err := ctx.BodyParser(app)
	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Something's wrong with your input",
				"data":    err,
			})
	}

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Id not found",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error creating Object ID",
			"data":    err,
		})
	}

	appBranch, err := model.GetAppModel(bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Finding AppBranch",
				"data":    err,
			})
	}

	updateApp, err := pkg.UpdateApp(&appBranch.AppId, &app.AppName)
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Updating App",
				"data":    err,
			})
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "appName", Value: updateApp.App.Name},
		{Key: "updated_at", Value: time.Now()},
	}}}

	updatedAppBranch, err := model.UpdateAppModel(bson.D{{Key: "_id", Value: objectId}}, update)
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Updating AppBranch",
				"data":    err,
			})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": updatedAppBranch,
	})
}
