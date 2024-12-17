package adapter

import (
	dtos2 "DataReader/data_supplier_reader/reader/dtos"
	"DataReader/data_supplier_reader/reader/dtos/request_dtos"
	"encoding/json"
)

type DeserializedCommand struct {
	//User json.RawMessage  `json:"user"`
	Type dtos2.ActionType `json:"type"`
}

func ConvertToCommand(serializedCommand []byte) dtos2.ReadCommand {
	var deserializedCommand DeserializedCommand
	commandUnmarshallError := json.Unmarshal(serializedCommand, &deserializedCommand)

	if commandUnmarshallError != nil {
		panic("bazdmeg0")
	}

	switch deserializedCommand.Type {
	case dtos2.Get:
		var dto request_dtos.UserFetchRequestDTO
		//dtoUnmarshallError := json.Unmarshal(deserializedCommand.User, &dto)
		//command, error1 := util.Cast[data_contracts.User, data_contracts.UserCreateDTO](deserializedCommand.User)
		//if dtoUnmarshallError != nil {
		//	panic("bazdmeg1")
		//}
		return dto

	case dtos2.GetAll:
		var dto request_dtos.AllUsersFetchRequestDTO
		//dtoUnmarshallError := json.Unmarshal(User, &dto)
		//if dtoUnmarshallError != nil {
		//	panic("bazdmeg1")
		//}
		return dto

	default:
		panic("bazdmeg2")
	}
}
