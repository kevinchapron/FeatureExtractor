package main

import "github.com/kevinchapron/FeatureExtractor/extractor"

func GetSensorsData(float_data [][]float64) (extractor.ListSensor,extractor.ListSensor,extractor.ListSensor){
	return extractor.GetSensorFromData(float_data,[3]int8{0,1,2}),
	extractor.GetSensorFromData(float_data,[3]int8{3,4,5}),
	extractor.GetSensorFromData(float_data,[3]int8{6,7,8});
}