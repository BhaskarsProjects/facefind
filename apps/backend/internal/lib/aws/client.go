package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/bhaskarvb28/facefind/internal/server"
)

type S3Client struct {
	client *s3.Client
}

func NewS3Client(
	server *server.Server,
) (*S3Client, error) {

	r2Config := server.Config.R2

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("auto"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				r2Config.AccessKeyID,
				r2Config.SecretAccessKey,
				"",
			),
		),
		config.WithRequestChecksumCalculation(
			aws.RequestChecksumCalculationWhenRequired,
		),
		config.WithResponseChecksumValidation(
			aws.ResponseChecksumValidationWhenRequired,
		),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"load R2 config: %w",
			err,
		)
	}

	client := s3.NewFromConfig(
		cfg,
		func(options *s3.Options) {
			options.BaseEndpoint = aws.String(
				fmt.Sprintf(
					"https://%s.r2.cloudflarestorage.com",
					r2Config.AccountID,
				),
			)
		},
	)

	return &S3Client{
		client: client,
	}, nil
}