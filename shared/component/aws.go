package sharedcomponent

import (
	"context"
	"io"
	
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type s3Uploader struct {
	apiKey     string
	bucketName string
	domain     string
	region     string
	secretKey  string
	client     *s3.Client
}

func NewS3Uploader(apiKey, bucketName, domain, region, secretKey string) (*s3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(apiKey, secretKey, ""),
		),
	)
	
	if err != nil {
		return nil, err
	}
	
	client := s3.NewFromConfig(cfg)
	
	return &s3Uploader{
		apiKey:     apiKey,
		bucketName: bucketName,
		domain:     domain,
		region:     region,
		secretKey:  secretKey,
		client:     client,
	}, nil
}

func (c *s3Uploader) SaveFileUpload(ctx context.Context, ioReader io.Reader, dst, contentType string, length int64) error {
	
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(c.bucketName),
		Key:           aws.String(dst),
		ACL:           types.ObjectCannedACLPrivate,
		Body:          ioReader,
		ContentType:   aws.String(contentType),
		ContentLength: &length,
	})
	
	if err != nil {
		return err
	}
	return nil
}

func (c *s3Uploader) GetDomain() string {
	return c.domain
}
