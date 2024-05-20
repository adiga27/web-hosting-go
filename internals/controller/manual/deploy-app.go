package manual

import (
	"time"

	"github.com/adiga27/web-hosting-go/internals/model"
	"github.com/adiga27/web-hosting-go/pkg"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeployApp(ctx *fiber.Ctx) error {
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
				"message": "Error Finding App",
				"data":    err,
			})
	}

	deploymentData, err := pkg.CreateDeployment(&appBranch.AppId, &appBranch.BranchName)
	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Creating Deployment",
				"data":    err,
			})
	}
	appBranch.JobId = *deploymentData.JobId

	status, err := pkg.UploadZipFile(ctx, deploymentData.ZipUploadUrl)

	if err != nil && status != 200 {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Uploading Files..",
				"data":    err,
			})
	}

	jobSummary, err := pkg.StartDeployment(&appBranch.AppId, &appBranch.BranchName, &appBranch.JobId)
	if err != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error starting Depolyment",
				"data":    err,
			})
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "jobId", Value: jobSummary.JobId},
		{Key: "status", Value: jobSummary.Status},
		{Key: "updated_at", Value: time.Now()},
	}}}

	updatedApp, err := model.UpdateAppModel(bson.D{{Key: "_id", Value: objectId}}, update)
	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Updating App",
				"data":    err,
			})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": updatedApp,
	})
}
