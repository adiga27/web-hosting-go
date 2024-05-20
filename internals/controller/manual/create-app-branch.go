package manual

import (
	"github.com/adiga27/web-hosting-go/internals/model"
	"github.com/adiga27/web-hosting-go/pkg"
	"github.com/gofiber/fiber/v2"
)

func CreateAppBranch(ctx *fiber.Ctx) error {
	appBranch := new(model.App)

	err := ctx.BodyParser(appBranch)
	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Something's wrong with your input",
				"data":    err,
			})
	}

	app, err := pkg.CreateApp()
	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error creating app",
				"data":    err,
			})
	}
	appBranch.AppId = *app.AppId

	branch, err := pkg.CreateBranch(&appBranch.AppId)
	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Creating Branch",
				"data":    err,
			})
	}

	appBranch.DisplayName = *branch.DisplayName
	appBranch.Url = ctx.Protocol() + "://" + appBranch.DisplayName + "." + appBranch.AppId + ".amplifyapp.com"

	err = appBranch.CreateAppModel()
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Creating AppBranch",
				"data":    err,
			})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"status": "success",
		"message": fiber.Map{
			"id":         appBranch.Id,
			"appName":    appBranch.AppName,
			"appId":      appBranch.AppId,
			"branchName": appBranch.BranchName,
			"url":        appBranch.Url,
		},
	})
}
