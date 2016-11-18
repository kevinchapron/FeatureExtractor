package csv

import (
	"github.com/kevinchapron/FeatureExtractor"
	"fmt"
	"strconv"
	"os"
	"encoding/csv"
)

// Combine all sensors in a single array
// 	Will panic if the FeatureExtractors sent haven't the same length (amount of sensors)
func combineExtractors(extractors []extractor.FeatureExtractor) []string{
	var data_return []string
	// Test for panic
	var f_len = 0
	for _, ext := range(extractors){
		if(f_len==0){
			f_len = ext.Length()
			continue
		}
		if(ext.Length()!=f_len){
			panic(fmt.Sprintf("Error: FeatureExtractors sent haven't the same length !\n"))
		}
	}
	// Panic test passed, let's move it !
	for _, fext := range(extractors){
		for index_feature, _ := range(extractor.LIST_OF_FEATURES_NAME){
			data_return = append(data_return,strconv.FormatFloat(fext.GetFeature(index_feature),'f',-1,64))
		}
	}

	return data_return
}
func getHeaders(extractors []extractor.FeatureExtractor) []string{
	var headers []string
	for _, fext := range(extractors){
		for _, header := range(extractor.LIST_OF_FEATURES_NAME){
			headers = append(headers,fext.Name+"_"+header)
		}
	}
	return headers
}

// Write single row of multiple devices in a file
func WriteExtractorsInFile(listExtractors []extractor.FeatureExtractor, filename string){
	fmt.Printf("Registering %v list of sensors in %s\n",len(listExtractors),filename)
	var fullData [][]string
	headers := getHeaders(listExtractors)
	data := combineExtractors(listExtractors)

	fullData = append(fullData,headers)
	fullData = append(fullData,data)

	createCSVFile(filename,fullData)

}

func WriteMultiExtractorsInFile(listExtractors [][]extractor.FeatureExtractor, filename string){
	fmt.Printf("Registering %v list of sensors in %s\n",len(listExtractors),filename)
	var headers []string = getHeaders(listExtractors[0])
	var fullData [][]string
	fullData = append(fullData,headers)
	for _, row := range listExtractors{
		fullData = append(fullData,combineExtractors(row))
	}
	createCSVFile(filename,fullData)
}

func createCSVFile(filename string, data [][]string){
	f, err := os.OpenFile(filename,os.O_WRONLY,0660)
	if(err != nil){
		if(os.IsNotExist(err)){
			f, err = os.Create(filename)
			if(err != nil){
				panic(fmt.Sprintf("Error : %v",err))
			}
		}else{
			panic(fmt.Sprintf("Error : %v",err))
		}
	}
	w := csv.NewWriter(f)
	w.WriteAll(data)

	err = w.Error()
	if(err!=nil){
		fmt.Printf("Error while Writing file : %v",err)
	}else{
		fmt.Printf("Features have been stored in file %s !\n",filename)
	}
}