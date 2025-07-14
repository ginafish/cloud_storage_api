package aws

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type mockListS3 func(ctx context.Context, params *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)

func (m mockListS3) ListObjects(ctx context.Context, params *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return m(ctx, params, opts...)
}

func mockS3Client() {
	return func(t *testing.T) S3ListObjectsAPI {
		return mockListS3(func(ctx context.Context, params *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
			t.Helper()
			return &s3.ListObjectsV2Output{}, nil
		})
	}
}

func TestListS3(t *testing.T) {

}
