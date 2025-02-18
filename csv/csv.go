package csv

import (
	"context"
	"fmt"
	"log"

	"github.com/omniful/go_commons/csv"
)

func ParseCSVFile(filepath string) [][]string {
	//configure csvReader
	csvReader, err := csv.NewCommonCSV(
		csv.WithBatchSize(100),    //read in batches(100 rows at a time)
		csv.WithSource(csv.Local), //will take file from local
		csv.WithLocalFileInfo("data.csv"),
	)
	if err != nil {
		fmt.Println("Error while initializing:", err)
	}
	// initialize csv reader
	err = csvReader.InitializeReader(context.TODO())
	if err != nil {
		fmt.Println("Error while initializing csv reader:", err)
	}
	var data [][]string
	for !csvReader.IsEOF() {
		var records csv.Records //2D slice of strings for csv records
		records, err = csvReader.ReadNextBatch()
		if err != nil {
			log.Fatal(err)
		}
		// Process the records
		var batch []string
		var headers csv.Headers = []string{"id", "seller_id", "item_count", "mode_of_payment", "status", "amount"}
		records.Unmarshal(headers, &batch)
		fmt.Println(batch)
		data = append(data, batch)
	}
	return data
}
