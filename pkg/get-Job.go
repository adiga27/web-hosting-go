package pkg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

func GetJob(appId *string, branchName *string, jobId *string) (*types.Job, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}
	job, err := client.GetJob(context.Background(), &amplify.GetJobInput{
		AppId:      appId,
		BranchName: branchName,
		JobId:      jobId,
	})
	if err != nil {
		return nil, err
	}

	return job.Job, nil
}
