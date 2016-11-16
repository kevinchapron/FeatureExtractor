package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/kevinchapron/FeatureExtractor/csv"
)

func main(){
	if(len(os.Args) < 2 || len(os.Args) > 3){
		fmt.Errorf("Error: Program should be called this way : \n")
		fmt.Errorf("        %s   CSV_FILE    [FOLDER_FOR_FEATURES_FILE]",os.Args[1])
		return
	}
	var CSV_FILE string
	var FEATURES_FOLDER string

	// Loading CSV_FILE from args,
	// And Reading features_folder from it, except if args[2] exists
	CSV_FILE = os.Args[1]
	if(len(os.Args) == 3){
		FEATURES_FOLDER=os.Args[2]
	}else{
		var last_index = strings.LastIndex(CSV_FILE,"/")
		if(last_index!=-1){
			FEATURES_FOLDER=CSV_FILE[:last_index+1]
			fmt.Println(last_index)
		}else{
			FEATURES_FOLDER="./"
		}
	}

	// Printings args
	fmt.Printf("CSV_FILE: %s\nFEATURES_FOLDER: %s\n",CSV_FILE,FEATURES_FOLDER)

	// Getting Data from CSV
	var data = csv.GetDataFromCSV(CSV_FILE)
	print(data)
}