package main

import (
	"fmt"
	"github.com/kevinchapron/FeatureExtractor/csv"
	"github.com/kevinchapron/FeatureExtractor/extractor"
)

func getSensorsData(float_data [][]float64) (extractor.ListSensor,extractor.ListSensor,extractor.ListSensor){
	return extractor.GetSensorFromData(float_data,[3]int8{0,1,2}),
	extractor.GetSensorFromData(float_data,[3]int8{3,4,5}),
	extractor.GetSensorFromData(float_data,[3]int8{6,7,8});
}

func main(){
	CSV_FILE, FEATURES_FILE := csv.ParseArgs()

	// Printings args
	fmt.Printf("CSV_FILE: %s\nFEATURES_FILE: %s\n",CSV_FILE,FEATURES_FILE)

	// Getting Data from CSV
	var dataFromCsv = csv.GetDataFromCSV(CSV_FILE,true)
	// Convert Data to Float64 slices
	var float_data = dataFromCsv.GetFloatData()
	// Extract Accelerometer Data
	accelerometer_data, gyroscope_data, magnetometer_data := getSensorsData(float_data)




	// Initialize FeatureExtractor
	var features_accelerometers = extractor.FeatureExtractor{Name:"Accelerometer",Data:accelerometer_data}
	// Calculation of Temporal & Frequential Features
	features_accelerometers.CalcTemporalFeatures()
	features_accelerometers.CalcFrequentialFeatures()
	// Print Results !
	//features_accelerometers.Print()

	// Initialize FeatureExtractor
	var features_gyroscope = extractor.FeatureExtractor{Name:"Gyroscope",Data:gyroscope_data}
	// Calculation of Temporal & Frequential Features
	features_gyroscope.CalcTemporalFeatures()
	features_gyroscope.CalcFrequentialFeatures()
	// Print Results !
	//features_gyroscope.Print()

	// Initialize FeatureExtractor
	var features_magnetometer = extractor.FeatureExtractor{Name:"Magnetometer",Data:magnetometer_data}
	// Calculation of Temporal & Frequential Features
	features_magnetometer.CalcTemporalFeatures()
	features_magnetometer.CalcFrequentialFeatures()
	// Print Results !
	//features_magnetometer.Print()




	var listExtractors []extractor.FeatureExtractor
	listExtractors = append(listExtractors,features_accelerometers)
	listExtractors = append(listExtractors,features_gyroscope)
	listExtractors = append(listExtractors,features_magnetometer)

	// Write it in a file
	csv.WriteExtractorsInFile(listExtractors,FEATURES_FILE)
}