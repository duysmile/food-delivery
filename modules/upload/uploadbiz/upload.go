package uploadbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component/uploadprovider"
	"200lab/food-delivery/modules/upload/uploadmodel"
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStore interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBiz(
	provider uploadprovider.UploadProvider,
	store CreateImageStore,
) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: store}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDemension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                              // img.jpg => .jpg
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt) // 123123123123.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	if err := biz.imgStore.CreateImage(ctx, img); err != nil {
		// delete img on S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}

func getImageDemension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
