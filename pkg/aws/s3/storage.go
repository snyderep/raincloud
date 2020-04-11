package s3

import (
	s3Service "github.com/aws/aws-sdk-go-v2/service/s3"
	"sync"
)

// StorageSummary contains a total number and size of objects.
type StorageSummary struct {
	Count int
	Size int64
	lock sync.Locker
}

func NewStorageSummary() *StorageSummary {
	return &StorageSummary{lock: &sync.Mutex{}}
}

// Update increments the summary and adds the size to the total.
func (summary *StorageSummary) Update(s3Object *s3Service.Object) {
	summary.lock.Lock()
	defer summary.lock.Unlock()

	summary.Count++
	summary.Size += *s3Object.Size
}

type Prefix struct {
	Children []*Prefix
	Parent *Prefix
	PrefixTail string
	StorageSummaries map[s3Service.StorageClass]StorageSummary
	storageClassLock sync.Locker
}
