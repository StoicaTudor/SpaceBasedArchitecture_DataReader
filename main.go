package main

import (
	"DataReader/data_supplier_reader/reader"
	"DataReader/data_supplier_reader/reader/consumers"
	"DataReader/environment"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	environment.Load()

	dataSupplier, _ := reader.GetDataReader()
	dataSupplier.Read(&consumers.DBReaderConsumer{}, nil)
}
