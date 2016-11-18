package examples

import (
	"fmt"
	"github.com/kevinchapron/FeatureExtractor/csv"
	"github.com/kevinchapron/FeatureExtractor"
)

func Multifile(){
	CSV_FOLDER, FEATURES_FILE := csv.ParseArgsMultifile()

	// Printings args
	fmt.Printf("CSV_FOLDER: %s\nFEATURES_FILE: %s\n",CSV_FOLDER,FEATURES_FILE)

	files := csv.GetFilesInFolder(CSV_FOLDER)

	// Declare slice containing all featureExtractor
	var featuresExtractors [][]extractor.FeatureExtractor

	for _, file := range(files){

		var filename = CSV_FOLDER+file

		// Getting Data from CSV
		var dataFromCsv = csv.GetDataFromCSV(filename,true)
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


		// Append featureExtractors to slices already declared
		featuresExtractors = append(featuresExtractors,[]extractor.FeatureExtractor{features_accelerometers,features_gyroscope,features_magnetometer})
	}

	// Write featuresExtractors in a CSV File
	csv.WriteMultiExtractorsInFile(featuresExtractors,FEATURES_FILE)
}