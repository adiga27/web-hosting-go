package pkg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

func StartDeployment(appId *string, branchName *string, jobId *string) (*types.JobSummary, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}
	start, err := client.StartDeployment(context.Background(), &amplify.StartDeploymentInput{
		AppId:      appId,
		BranchName: branchName,
		JobId:      jobId,
	})

	return start.JobSummary, err
}
