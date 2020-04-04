package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	s3service "github.com/aws/aws-sdk-go-v2/service/s3"
)

type Manager struct {
	client *s3service.Client
}

func NewManager(config *aws.Config) *Manager {
	return &Manager{client: s3service.New(*config)}
}

func (manager *Manager) GetBucket(bucketName string) (s3service.Bucket, error) {
	buckets, err := manager.ListBuckets()
	if err != nil {
		return s3service.Bucket{}, err
	}

	for _, bucket := range buckets {
		if *bucket.Name == bucketName {
			return bucket, nil
		}
	}

	return s3service.Bucket{}, nil
}

func (manager *Manager) ListBuckets() ([]s3service.Bucket, error) {
	req := manager.client.ListBucketsRequest(nil)

	resp, err := req.Send(context.Background())
	if err != nil {
		return []s3service.Bucket{}, nil
	}

	return resp.Buckets, nil
}
