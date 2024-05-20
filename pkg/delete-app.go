package pkg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func DeleteApp(appId *string) (*amplify.DeleteAppOutput, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}

	data, err := client.DeleteApp(context.Background(), &amplify.DeleteAppInput{
		AppId: appId,
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}
