package csvplate

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
)


func ParseLocation(file string) ([]map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvr := csv.NewReader(f)
	data := make([]map[string]string, 0)
	header, err := csvr.Read()
	if err != nil {
		return nil, fmt.Errorf("[Failed to read header] %w", err)
	}

	for {
		rowraw, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
		row := map[string]string{}
		for i, r := range rowraw {
			row[header[i]] = r
		}
		data = append(data, row)
	}
}
