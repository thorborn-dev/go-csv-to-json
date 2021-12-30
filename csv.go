package csv

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/thorborn-dev/go-csv-to-json/util"
)

type CSV struct {
	rows [][]string
}

func ReadCSV(filename string) (CSV, error) {
	f, err := os.Open(filename)
	if err != nil {
		return CSV{}, err
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return CSV{}, err
	}
	return CSV{rows: rows}, nil
}

func (csv *CSV) ToMap() []map[string]interface{} {
	var entries []map[string]interface{}
	keys := csv.rows[0]

	for _, row := range csv.rows[1:] {
		entry := map[string]interface{}{}

		for i, value := range row {
			key := keys[i]
			// Split key for nested objects or arrays
			keyParts := strings.Split(key, ".")
			internalEntry := entry

			for j, keyPart := range keyParts {
				arrayKey, arrayIndex := getArrayKeyAndIndex(keyPart)
				if arrayIndex != -1 {
					if internalEntry[arrayKey] == nil {
						internalEntry[arrayKey] = []interface{}{}
					}
					internalArray := internalEntry[arrayKey].([]interface{})
					if j == len(keyParts)-1 {
						isInt := util.IsStringInt(value)
						if isInt {
							v, _ := strconv.Atoi(value)
							internalArray = append(internalArray, v)
						} else {
							internalArray = append(internalArray, value)
						}
						internalEntry[arrayKey] = internalArray
						// No further processing needed since we reached the end...
						break
					}
					if arrayIndex >= len(internalArray) {
						internalArray = append(internalArray, map[string]interface{}{})
					}
					internalEntry[arrayKey] = internalArray
					internalEntry = internalArray[arrayIndex].(map[string]interface{})
				} else {
					// Is either the last part of a nested object or isn't nested at all
					if j == len(keyParts)-1 {
						isInt := util.IsStringInt(value)
						if isInt {
							v, _ := strconv.Atoi(value)
							internalEntry[keyPart] = v
						} else {
							internalEntry[keyPart] = value
						}
						// No further processing needed since we reached the end...
						break
					}
					if internalEntry[keyPart] == nil {
						internalEntry[keyPart] = map[string]interface{}{}
					}
					internalEntry = internalEntry[keyPart].(map[string]interface{})
				}
			}
		}
		entries = append(entries, entry)
	}

	return entries
}

func (csv *CSV) ToJSON() (string, error) {
	entries := csv.ToMap()
	bytes, err := json.Marshal(entries)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getArrayKeyAndIndex(str string) (string, int) {
	i := strings.Index(str, "[")
	if i >= 0 {
		j := strings.Index(str, "]")
		if j >= 0 {
			index, _ := strconv.Atoi(str[i+1 : j])
			return str[0:i], index
		}
	}
	return str, -1
}
