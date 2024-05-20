package pkg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func UpdateApp(appId *string, name *string) (*amplify.UpdateAppOutput, error) {
	client, err := AmplifyClient()
	if err != nil {
		return nil, err
	}
	updatedApp, err := client.UpdateApp(context.Background(), &amplify.UpdateAppInput{
		AppId: appId,
		Name:  name,
	})
	if err != nil {
		return nil, err
	}

	return updatedApp, nil
}
