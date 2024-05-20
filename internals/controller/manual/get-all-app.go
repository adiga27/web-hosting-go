package manual

import (
	"github.com/adiga27/web-hosting-go/internals/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllApp(ctx *fiber.Ctx) error {
	appBranches, err := model.GetAllAppModel()
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Retrieving Apps",
				"data":    err,
			})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": appBranches,
	})
}
