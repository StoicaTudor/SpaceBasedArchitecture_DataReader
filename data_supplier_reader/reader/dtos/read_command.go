package dtos

type ActionType string

const (
	Get    ActionType = "GET"
	GetAll ActionType = "GET_ALL"
)

type ReadCommandPayloadType string

const (
	FetchUserRequestDTO     ReadCommandPayloadType = "FetchUserRequestDTO"
	FetchAllUsersRequestDTO ReadCommandPayloadType = "FetchAllUsersRequestDTO"
)

type FilterType string

const (
	ID       FilterType = "ID"
	NoFilter FilterType = "NoFilter"
)

type ReadCommand interface {
	GetAction() ActionType
	GetReadCommandPayloadType() ReadCommandPayloadType
	GetFilter() FilterType
}
