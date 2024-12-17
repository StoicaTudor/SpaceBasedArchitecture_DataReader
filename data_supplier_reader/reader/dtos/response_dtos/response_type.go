package response_dtos

type ResponseType string

const (
	GetUser     ResponseType = "GET_USER"
	GetAllUsers ResponseType = "GET_ALL_USERS"
)
