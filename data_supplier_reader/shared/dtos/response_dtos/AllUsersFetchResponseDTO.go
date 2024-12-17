package response_dtos

type AllUsersFetchResponseDTO struct {
	AllUsers []UserFetchResponseDTO `json:"users"`
}
