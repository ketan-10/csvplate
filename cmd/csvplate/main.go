package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/ketan-10/csvplate"
)

func main() {

	if len(os.Args) < 4 {
		panic("ERROR invalid call missing datafile templatefile or outputfilename")
	}
	var split int
	if len(os.Args) > 4 {
		var err error
		split, err = strconv.Atoi(os.Args[4])
		if err != nil {
			panic("ERROR last argument split must be integer")
		}
	}
	dataFile := os.Args[1]
	templateFile := os.Args[2]
	outputFileName := os.Args[3]

	csvdata, err := csvplate.ParseLocation(dataFile)
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	
	if split == 0 {
		file, err := os.Create(outputFileName)
		if err != nil {
			panic(err)
		}
		err = t.Execute(file, csvdata)
		if err != nil {
			panic(err)
		}
	} else {
		
		lastIdx := strings.LastIndex(outputFileName, ".")
		if lastIdx == -1 {
			lastIdx = len(outputFileName)
		}
		prefix := outputFileName[:lastIdx]
		sufix := outputFileName[lastIdx:]
		fileIdx := 1
		for i := 0; i < len(csvdata); i += split {

			file, err := os.Create(fmt.Sprintf("%s-%d%s",prefix, fileIdx, sufix))
			if err != nil {
				panic(err)
			}
			if i+split >= len(csvdata) {
				t.Execute(file, csvdata[i:])
			}else {
				t.Execute(file, csvdata[i : i+split])
			}
			fileIdx++
		}
	}

}
