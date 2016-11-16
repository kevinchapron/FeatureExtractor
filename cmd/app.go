package main

import (
	"fmt"
	"github.com/kevinchapron/FeatureExtractor/csv"
	"github.com/kevinchapron/FeatureExtractor/extractor"
)

func main(){
	CSV_FILE, FEATURES_FOLDER := csv.ParseArgs()

	// Printings args
	fmt.Printf("CSV_FILE: %s\nFEATURES_FOLDER: %s\n",CSV_FILE,FEATURES_FOLDER)


	// Getting Data from CSV
	var dataFromCsv = csv.GetDataFromCSV(CSV_FILE,true)
	var float_data = dataFromCsv.GetFloatData()
	var accelerometer_data 	= extractor.GetSensorFromData(float_data,[3]int8{0,1,2})
	var gyroscope_data 		= extractor.GetSensorFromData(float_data,[3]int8{3,4,5})
	var magnetometer_data	= extractor.GetSensorFromData(float_data,[3]int8{6,7,8})
}