package csv

import (
	"fmt"
	"io/ioutil"
)

func GetFilesInFolder(folder string) []string {
	files_and_folder, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(fmt.Sprintf("Error while reading %s : %v",folder,err))
	}
	var files_only []string
	for _, file_or_folder := range(files_and_folder){
		if(file_or_folder.IsDir()){
			continue
		}
		files_only = append(files_only,file_or_folder.Name())
	}
	return files_only
}
