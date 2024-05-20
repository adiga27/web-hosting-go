package pkg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

func CreateBranch(appId *string) (*types.Branch, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}

	data, err := client.CreateBranch(context.Background(), &amplify.CreateBranchInput{
		AppId:       appId,
		BranchName:  aws.String("dev"),
		DisplayName: aws.String("adiga"),
	})
	if err != nil {
		return nil, err
	}

	return data.Branch, nil
}
