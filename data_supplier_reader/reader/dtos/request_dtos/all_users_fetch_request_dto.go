package request_dtos

import "DataReader/data_supplier_reader/reader/dtos"

type AllUsersFetchRequestDTO struct {
}

func (dto AllUsersFetchRequestDTO) GetAction() dtos.ActionType {
	return dtos.GetAll
}

func (dto AllUsersFetchRequestDTO) GetReadCommandPayloadType() dtos.ReadCommandPayloadType {
	return dtos.FetchAllUsersRequestDTO
}

func (dto AllUsersFetchRequestDTO) GetFilter() dtos.FilterType {
	return dtos.NoFilter
}
