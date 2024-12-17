package reader

import (
	dtos2 "DataReader/data_supplier_reader/reader/dtos"
	"DataReader/data_supplier_reader/reader/dtos/request_dtos"
	"DataReader/util"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

type MockReader struct{}

func (mockSupplier MockReader) Read(consumer util.Consumer[dtos2.ReadCommand], wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			data := getRandomCommand()
			consumer.Consume(data)
			time.Sleep(2 * time.Second)
		}
	}()
}

func getRandomCommand() dtos2.ReadCommand {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch random.Intn(2) {
	case 0:
		return &request_dtos.UserFetchRequestDTO{
			ID: uuid.New().String(),
		}
	case 1:
		return &request_dtos.AllUsersFetchRequestDTO{}
	default:
		return nil
	}
}
