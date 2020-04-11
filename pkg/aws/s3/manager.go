package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	s3Service "github.com/aws/aws-sdk-go-v2/service/s3"
)

// Manager is a container for an AWS SDK S3 client and provides various S3 services.
type Manager struct {
	client *s3Service.Client
}

// NewManager creates a new Manager configured with aws.Config values.
func NewManager(config *aws.Config) *Manager {
	return &Manager{client: s3Service.New(*config)}
}

// GetBucketManager returns a BucketManager or an error.
func (manager *Manager) GetBucketManager(bucketName string) (*BucketManager, error) {
	bucketManagers, err := manager.ListBucketManagers()
	if err != nil {
		return &BucketManager{}, err
	}

	for _, bucketManager := range bucketManagers {
		if *bucketManager.bucket.Name == bucketName {
			return bucketManager, nil
		}
	}

	return &BucketManager{}, errors.New(fmt.Sprintf("s3: no bucket found for bucket '%s'", bucketName))
}

// ListBucketManagers returns a slice of BucketManager(s) for buckets that are visible to the AWS credentials from the
// default credential loading chain. See "Specifying Credentials" in
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html.
func (manager *Manager) ListBucketManagers() ([]*BucketManager, error) {
	req := manager.client.ListBucketsRequest(nil)

	resp, err := req.Send(context.Background())
	if err != nil {
		return []*BucketManager{}, nil
	}

	var bucketManagers = make([]*BucketManager, len(resp.Buckets))
	for i, bucket := range resp.Buckets {
		bucketManagers[i] = NewBucketManager(&bucket)
	}

	return bucketManagers, nil
}
