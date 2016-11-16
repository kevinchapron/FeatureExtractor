package csv

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
	"strconv"
)
/**
 Global method to extract Data from CSV FIle
 	filename:= path to file to extract data from
 	headers:=  Does the file contains headers name in first line ?

 	return: CSVFileStruct object
 **/

func GetDataFromCSV(filename string, headers bool) CSVFileStruct {
	f, err := os.Open(filename)
	if(err != nil){
		panic(fmt.Sprintf("Error : %v",err))
	}
	r := csv.NewReader(bufio.NewReader(f))
	data, _ := r.ReadAll()
	return CSVFileStruct{headers:data[0],data:data[1:]}
}


type CSVFileStruct struct{
	headers []string
	data [][]string

	float_data [][]float64
}
func(csvFileStruct *CSVFileStruct) GetFloatData() [][]float64{
	if(len(csvFileStruct.float_data) != len(csvFileStruct.data)){
		for _, row_raw_data := range(csvFileStruct.data){
			var row_float_data []float64;
			for _, row_str_data_item := range(row_raw_data){
				value, err := strconv.ParseFloat(row_str_data_item,64)
				if(err != nil){
					panic(fmt.Sprintf("Error: %v\n",err))
				}
				row_float_data = append(row_float_data, value)
			}
			csvFileStruct.float_data = append(csvFileStruct.float_data,row_float_data)
		}
	}
	return csvFileStruct.float_data
}