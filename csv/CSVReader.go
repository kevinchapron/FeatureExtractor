package csv

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
)

func GetDataFromCSV(filename string) [][]string {
	f, err := os.Open(filename)
	if(err != nil){
		panic(fmt.Sprintf("Error : %v",err))
	}
	r := csv.NewReader(bufio.NewReader(f))
	data, _ := r.ReadAll()
	return data
}