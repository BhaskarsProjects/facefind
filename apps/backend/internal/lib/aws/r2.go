package aws

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *S3Client) Upload(
	ctx context.Context,
	bucket string,
	key string,
	reader io.Reader,
	contentType string,
	size int64,
) error {

	_, err := s.client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:        aws.String(bucket),
			Key:           aws.String(key),
			Body:          reader,
			ContentType:   aws.String(contentType),
			ContentLength: aws.Int64(size),
		},
	)

	if err != nil {
		return fmt.Errorf(
			"upload object %q: %w",
			key,
			err,
		)
	}

	return nil
}

func (s *S3Client) Delete(
	ctx context.Context,
	bucket string,
	key string,
) error {

	_, err := s.client.DeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	)

	if err != nil {
		return fmt.Errorf(
			"delete object %q: %w",
			key,
			err,
		)
	}

	return nil
}

func (s *S3Client) CreatePresignedURL(
	ctx context.Context,
	bucket string,
	objectKey string,
) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	expiration := time.Hour

	presignedURL, err := presignClient.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(objectKey),
		},
		s3.WithPresignExpires(expiration),
	)
	if err != nil {
		return "", err
	}

	return presignedURL.URL, nil
}