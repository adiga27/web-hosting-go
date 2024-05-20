package manual

import (
	"time"

	"github.com/adiga27/web-hosting-go/internals/model"
	"github.com/adiga27/web-hosting-go/pkg"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAppStatus(ctx *fiber.Ctx) error {
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
	var job *types.Job
	err = nil
	ticker := time.NewTicker(3 * time.Second)
	done := make(chan bool)

	if (appBranch.Status != "SUCCEED") && (appBranch.Status != "FAILED") {
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					job, err = pkg.GetJob(&appBranch.AppId, &appBranch.BranchName, &appBranch.JobId)
					if err != nil {
						ticker.Stop()
						done <- true
						return
					}
					if (job.Summary.Status == types.JobStatusSucceed) || (job.Summary.Status == types.JobStatusFailed) {
						ticker.Stop()
						done <- true
						return
					}
				}
			}
		}()
		time.Sleep(9 * time.Second)
	}

	if err != nil {
		return ctx.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Error Retreiving App Status",
				"data":    err,
			})
	}

	var updatedApp *model.App
	if job != nil {
		update := bson.D{{Key: "$set", Value: bson.D{
			{Key: "status", Value: job.Summary.Status},
			{Key: "updated_at", Value: time.Now()},
		}}}

		updatedApp, err = model.UpdateAppModel(bson.D{{Key: "_id", Value: objectId}}, update)
		if err != nil {
			return ctx.Status(500).JSON(
				fiber.Map{
					"status":  "error",
					"message": "Error Updating App",
					"data":    err,
				})
		}
	}
	if updatedApp != nil {
		return ctx.Status(200).JSON(
			fiber.Map{
				"status":  "success",
				"message": updatedApp,
			})
	}
	if appBranch.Status == "SUCCEED" || appBranch.Status == "FAILED" {
		return ctx.Status(200).JSON(
			fiber.Map{
				"status":  "success",
				"message": appBranch,
			})
	}
	return ctx.Status(500).JSON(
		fiber.Map{
			"status":  "error",
			"message": appBranch,
		})
}
