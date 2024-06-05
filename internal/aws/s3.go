package cloud_aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3() *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApSouth1RegionID),
	}))
	// creds := stscreds.NewCredentials(sess, "arn:aws:s3:::66049c07d9e8546699fe0872fd32d8f6")
	return s3.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ID"),
			os.Getenv("AWS_SECRET"),
			"",
		),
	})
}
