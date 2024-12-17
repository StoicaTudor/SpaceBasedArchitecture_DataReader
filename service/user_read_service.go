package service

import (
	"DataReader/data"
	"DataReader/data_supplier_reader/reader/dtos/request_dtos"
	"DataReader/repository"
)

// TODO: handle errors
func GetUser(dto request_dtos.UserFetchRequestDTO) (data.User, error) {
	return repository.GetUser(dto.ID)
}

func GetAllUsers(dto request_dtos.AllUsersFetchRequestDTO) ([]data.User, error) {
	return repository.GetAllUsers()
}
