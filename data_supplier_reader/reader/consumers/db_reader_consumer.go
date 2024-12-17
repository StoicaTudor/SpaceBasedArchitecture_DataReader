package consumers

import (
	"DataReader/data_supplier_reader/reader/dtos"
	"DataReader/data_supplier_reader/reader/dtos/request_dtos"
	"DataReader/data_supplier_reader/reader/dtos/response_dtos"
	"DataReader/data_supplier_reader/supplier"
	"DataReader/service"
	"DataReader/util"
)

type DBReaderConsumer struct{}

func (consumer DBReaderConsumer) Consume(command dtos.ReadCommand) {
	switch command.GetReadCommandPayloadType() {
	case dtos.FetchUserRequestDTO:
		dto, err := util.Cast[dtos.ReadCommand, request_dtos.UserFetchRequestDTO](command)
		if err != nil {
			return
		}
		user, _ := service.GetUser(dto)
		supplier.Send(user, user.ID)
		break

	case dtos.FetchAllUsersRequestDTO:
		dto, err := util.Cast[dtos.ReadCommand, request_dtos.AllUsersFetchRequestDTO](command)
		if err != nil {
			return
		}
		allUsers, _ := service.GetAllUsers(dto)
		supplier.Send(response_dtos.AllUsersResponseDTO{AllUsers: allUsers, ResponseType: response_dtos.GetAllUsers}, "allUsers")
		break

	default:
		panic("bazdmeg default")
	}
}
