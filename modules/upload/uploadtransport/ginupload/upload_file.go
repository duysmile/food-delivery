package ginupload

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/upload/uploadbiz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// fileName := fileHeader.Filename
		// destination := fmt.Sprintf("./%s/%s", folder, fileName)
		// log.Println(destination)
		// errUpload := c.SaveUploadedFile(fileHeader, destination)
		// if errUpload != nil {
		// 	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 		"error": errUpload,
		// 	})
		// 	return
		// }

		// imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.GetUploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
