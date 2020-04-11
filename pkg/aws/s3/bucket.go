package s3

import (
	"context"
	s3Service "github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

type BucketManager struct {
	bucket *s3Service.Bucket
}

func NewBucketManager(bucket *s3Service.Bucket) *BucketManager {
	return &BucketManager{bucket}
}

func (manager *BucketManager) QueueObjects(
	objChan chan *s3Service.Object,
	errChan chan error,
	ctx context.Context,
) {
	input := &s3Service.ListObjectsV2Input{Bucket: manager.bucket.Name}
	request := s3Service.ListObjectsV2Request{Input: input}
	paginator := s3Service.NewListObjectsV2Paginator(request)

	for paginator.Next(ctx) {
		output := paginator.CurrentPage()
		for _, obj := range output.Contents {
			objChan <- &obj
		}
	}

	if err := paginator.Err(); err != nil {
		errChan <- err
	}
}

func (manager *BucketManager) Summarize() *Prefix {
	objChan := make(chan *s3Service.Object, 8)
	errChan := make(chan error)
	ctx := context.Background()

	go func() {
		for {
			select {
			case s3Object := <-objChan:
				manager.updateSummary(s3Object)
			case err := <- errChan:
				ctx.Done()
				log.Fatalln(err)
			}
		}
	}()

	go func() {
		manager.QueueObjects(objChan, errChan, ctx)
	}()


}

func (manager *BucketManager) updateSummary(s3Object *s3Service.Object) {

}
