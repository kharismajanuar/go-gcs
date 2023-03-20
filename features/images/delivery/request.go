package delivery

import (
	"go-gcs/features/images"
	"mime/multipart"
)

type ImageRequest struct {
	ImageFile multipart.File `json:"image" form:"image"`
}

func requestToCore(dataRequest ImageRequest) images.Core {
	return images.Core{
		ImageFile: dataRequest.ImageFile,
	}
}
