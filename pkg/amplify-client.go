package pkg

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

func AmplifyClient() (*amplify.Client, error) {
	staticProvider := credentials.NewStaticCredentialsProvider(
		os.Getenv("W_AWS_ACCESS_KEY"),
		os.Getenv("W_AWS_SECRET_KEY"),
		"",
	)

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithDefaultRegion(os.Getenv("W_AWS_REGION")),
		config.WithCredentialsProvider(staticProvider),
	)

	if err != nil {
		return nil, err
	}

	client := amplify.NewFromConfig(cfg)

	return client, nil
}
