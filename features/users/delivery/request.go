package delivery

import (
	"go-gcs/features/users"
	"mime/multipart"
)

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	DisplayImage string
	ImageFile    multipart.File `json:"image" form:"image"`
	ImageName    string
}

func requestToCore(dataRequest UserRequest) users.Core {
	return users.Core{
		Name:         dataRequest.Name,
		DisplayImage: dataRequest.DisplayImage,
		ImageFile:    dataRequest.ImageFile,
		ImageName:    dataRequest.ImageName,
	}
}
