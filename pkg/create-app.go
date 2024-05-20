package pkg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

func CreateApp() (*types.App, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}

	data, err := client.CreateApp(context.Background(), &amplify.CreateAppInput{
		Name:     aws.String("Krishna"),
		Platform: types.PlatformWeb,
	})

	if err != nil {
		return nil, err
	}

	return data.App, nil
}
