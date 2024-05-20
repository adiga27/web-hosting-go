package manual

import (
	"github.com/adiga27/web-hosting-go/internals/model"
	"github.com/adiga27/web-hosting-go/pkg"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteAppBranch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Id not found",
				"data":    id,
			})
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
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

	_, err = pkg.DeleteApp(&appBranch.AppId)
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Deleting App",
				"data":    err,
			})
	}

	err = model.DeleteAppModel(bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Deleting AppBranch",
				"data":    err,
			})
	}

	return ctx.Status(200).JSON(
		fiber.Map{
			"status":  "success",
			"message": "App Deleted",
			"data":    nil,
		})
}
