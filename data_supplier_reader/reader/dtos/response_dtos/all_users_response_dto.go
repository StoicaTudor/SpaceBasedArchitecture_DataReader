package response_dtos

import "DataReader/data"

type AllUsersResponseDTO struct {
	AllUsers     []data.User  `json:"users"`
	ResponseType ResponseType `json:"responseType"`
}
