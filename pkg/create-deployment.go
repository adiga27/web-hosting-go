package pkg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func CreateDeployment(appId *string, branchName *string) (*amplify.CreateDeploymentOutput, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}

	data, err := client.CreateDeployment(context.Background(), &amplify.CreateDeploymentInput{
		AppId:      appId,
		BranchName: branchName,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
