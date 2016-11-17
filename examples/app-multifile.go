package examples

import (
	"fmt"
	"github.com/kevinchapron/FeatureExtractor/csv"
	"io/ioutil"
	"github.com/kevinchapron/FeatureExtractor/extractor"
)

func Multifile(){
	CSV_FOLDER, FEATURES_FILE := csv.ParseArgsMultifile()

	// Printings args
	fmt.Printf("CSV_FOLDER: %s\nFEATURES_FILE: %s\n",CSV_FOLDER,FEATURES_FILE)

	files, err := ioutil.ReadDir(CSV_FOLDER)
	if err != nil {
		fmt.Printf("Error while reading %s : %v",CSV_FOLDER,err)
		return
	}

	// Declare slice containing all featureExtractor
	var featuresExtractors [][]extractor.FeatureExtractor

	for _, file := range(files){
		var filename = CSV_FOLDER+file.Name()

		// Getting Data from CSV
		var dataFromCsv = csv.GetDataFromCSV(filename,true)
		// Convert Data to Float64 slices
		var float_data = dataFromCsv.GetFloatData()
		// Extract Accelerometer Data
		accelerometer_data, gyroscope_data, magnetometer_data := extractor.GetSensorsData(float_data)

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