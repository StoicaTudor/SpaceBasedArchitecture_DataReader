package reader

import (
	dtos2 "DataReader/data_supplier_reader/reader/dtos"
	"DataReader/environment"
	"DataReader/util"
	"fmt"
	"os"
	"sync"
)

type DataReader interface {
	Read(util.Consumer[dtos2.ReadCommand], *sync.WaitGroup)
}

func GetDataReader() (DataReader, error) {
	switch os.Getenv(string(environment.DataSupplySource)) {
	case "kafka":
		return &KafkaReader{}, nil
	case "mock":
		return &MockReader{}, nil
	default:
		return nil, fmt.Errorf("unknown environment: %s", environment.DataSupplySource)
	}
}
