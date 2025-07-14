package aws

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	accountId   = "169032304403"
	buildClient sync.Once
	client      *s3.Client
)

type ListS3Req struct {
	Bucket string
	Prefix string
}

func InitializeClient() *s3.Client {
	buildClient.Do(func() {
		cfg, _ := config.LoadDefaultConfig(context.TODO())
		client = s3.NewFromConfig(cfg)
	})
	return client
}

func ListS3(bucket string) []types.Object {
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Fatal(err)
	}
	return output.Contents
}
