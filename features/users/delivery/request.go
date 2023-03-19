package delivery

import "go-gcs/features/users"

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	DisplayImage string `json:"display_image" form:"display_image"`
}

func requestToCore(dataRequest UserRequest) users.Core {
	return users.Core{
		Name:         dataRequest.Name,
		DisplayImage: dataRequest.DisplayImage,
	}
}
