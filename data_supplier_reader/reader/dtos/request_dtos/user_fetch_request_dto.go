package request_dtos

import "DataReader/data_supplier_reader/reader/dtos"

type UserFetchRequestDTO struct {
	ID string `json:"id"`
}

func (dto UserFetchRequestDTO) GetAction() dtos.ActionType {
	return dtos.Get
}

func (dto UserFetchRequestDTO) GetReadCommandPayloadType() dtos.ReadCommandPayloadType {
	return dtos.FetchUserRequestDTO
}

func (dto UserFetchRequestDTO) GetFilter() dtos.FilterType {
	return dtos.ID
}
