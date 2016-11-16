package csv

import (
	"github.com/kevinchapron/FeatureExtractor/extractor"
	"fmt"
	"strconv"
	"os"
	"encoding/csv"
)

// Combine all sensors in a single array
// 	Will panic if the FeatureExtractors sent haven't the same length (amount of sensors)
func combineExtractors(extractors []extractor.FeatureExtractor) [][]string{
	var data_return [][]string
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
	// Get Headers and put it first
	var headers []string
	for _, fext := range(extractors){
		for _, header := range(extractor.LIST_OF_FEATURES_NAME){
			headers = append(headers,fext.Name+"_"+header)
		}
	}
	// Get Content and put it next !

	data_return = append(data_return,headers)
	var str_append []string
	for _, fext := range(extractors){
		for index_feature, _ := range(extractor.LIST_OF_FEATURES_NAME){
			str_append = append(str_append,strconv.FormatFloat(fext.GetFeature(index_feature),'f',-1,64))
		}
	}
	data_return = append(data_return,str_append)

	return data_return
}

func WriteExtractorsInFile(listExtractors []extractor.FeatureExtractor, filename string){
	fmt.Printf("Registering %v list of sensors in %s\n",len(listExtractors),filename)
	data := combineExtractors(listExtractors)

	f, err := os.Open(filename)
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