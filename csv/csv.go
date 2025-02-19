package csv

import (
	"context"
	"fmt"
	"log"

	"github.com/omniful/go_commons/csv"
)

func ParseCSVFile(filepath string) []map[string]string {
	fmt.Println("parsing....")
	//configure csvReader
	csvReader, err := csv.NewCommonCSV(
		csv.WithBatchSize(100),    //read in batches(100 rows at a time)
		csv.WithSource(csv.Local), //will take file from local
		csv.WithLocalFileInfo("/Users/namita/Downloads/project.csv"),
	)
	if err != nil {
		fmt.Println("Error while initializing:", err)
	}
	// initialize csv reader
	err = csvReader.InitializeReader(context.TODO())
	if err != nil {
		fmt.Println("Error while initializing csv reader:", err)
	}
	var data []map[string]string
	for !csvReader.IsEOF() {
		var records csv.Records //2D slice of strings for csv records
		records, err = csvReader.ReadNextBatch()
		if err != nil {
			log.Fatal(err)
		}
		// Process the records
		var batch []map[string]string
		var headers csv.Headers = []string{"id", "order_id", "seller_id", "sku_id", "item_count", "status"}
		records.Unmarshal(headers, &batch)
		data = append(data, batch...)
	}
	fmt.Println("CSV parsed successfully!")
	fmt.Println(data)
	return data
}
