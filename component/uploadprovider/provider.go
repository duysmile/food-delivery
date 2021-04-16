package uploadprovider

import (
	"200lab/food-delivery/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
