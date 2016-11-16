package csv

import (
	"os"
	"fmt"
	"strings"
)

func ParseArgs() (string,string){
	if(len(os.Args) < 2 || len(os.Args) > 3){
		panic(fmt.Sprintf("Error: Program should be called this way : \n    %s   CSV_FILE    [FOLDER_FOR_FEATURES_FILE]",os.Args[0]))
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

	// Get filename for Feature_file
	var filename_only string
	var last_index = strings.LastIndex(CSV_FILE,"/")
	if(last_index!=-1){
		filename_only=CSV_FILE[last_index+1:]
	}else{
		filename_only=CSV_FILE
	}
	// Add _feature_ between filename & extension
	last_index = strings.LastIndex(filename_only,".")
	if(last_index!=-1){
		filename_only=filename_only[:last_index]+"__features__"+filename_only[last_index:]
	}else{
		panic(fmt.Sprintf("Error: First parameter have to be a file with extension"))
	}

	return CSV_FILE,FEATURES_FOLDER+filename_only
}
