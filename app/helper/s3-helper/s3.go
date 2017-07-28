package s3_helper

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToS3(from string, to string) (string, error) {
	sess := session.Must(session.NewSession())
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	zipFile, err := os.Open(from)
	if err != nil {
		return "", err
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET_NAME")),
		Key:    aws.String(to),
		Body:   zipFile,
	})
	if err != nil {
		return "", err
	}

	return aws.StringValue(&result.Location), nil
}