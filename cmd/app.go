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
	// Convert Data to Float64 slices
	var float_data = dataFromCsv.GetFloatData()
	// Extract Accelerometer Data
	var accelerometer_data 	= extractor.GetSensorFromData(float_data,[3]int8{0,1,2})

	// Initialize FeatureExtractor
	var features_accelerometers = extractor.FeatureExtractor{Name:"Features from Accelerometer",Data:accelerometer_data}
	// Calculation of Temporal & Frequential Features
	features_accelerometers.CalcTemporalFeatures()
	features_accelerometers.CalcFrequentialFeatures()
	// Results !
	features_accelerometers.Print()

}