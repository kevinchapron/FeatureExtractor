package examples

import (
	"fmt"
	"github.com/kevinchapron/FeatureExtractor/csv"
	"github.com/kevinchapron/FeatureExtractor"
)


func Singlefile(){
	CSV_FILE, FEATURES_FILE := csv.ParseArgs()

	// Printings args
	fmt.Printf("CSV_FILE: %s\nFEATURES_FILE: %s\n",CSV_FILE,FEATURES_FILE)

	// Getting Data from CSV
	var dataFromCsv = csv.GetDataFromCSV(CSV_FILE,true)
	// Convert Data to Float64 slices
	var float_data = dataFromCsv.GetFloatData()
	// Extract Accelerometer Data
	var devices_data = extractor.GetSensorsData(float_data,3)
	var accelerometer_data = devices_data[0]
	var gyroscope_data = devices_data[1]
	var magnetometer_data = devices_data[2]

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