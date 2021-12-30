package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	testcases := []struct {
		name         string
		filename     string
		expectedJSON string
	}{
		{
			name:         "Strings",
			filename:     "mocks/01-strings.csv",
			expectedJSON: `[{"name":"John Doe"}]`,
		},
		{
			name:         "Integers",
			filename:     "mocks/02-integers.csv",
			expectedJSON: `[{"age":25}]`,
		},
		{
			name:         "Objects",
			filename:     "mocks/03-objects.csv",
			expectedJSON: `[{"user":{"firstname":"John","lastname":"Doe"}}]`,
		},
		{
			name:         "Multi-level objects",
			filename:     "mocks/04-multi-level-objects.csv",
			expectedJSON: `[{"user":{"birthdate":{"day":1,"month":1,"year":1990}}}]`,
		},
		{
			name:         "Arrays of strings",
			filename:     "mocks/05-arrays-of-strings.csv",
			expectedJSON: `[{"names":["John","Jane"]}]`,
		},
		{
			name:         "Arrays of integers",
			filename:     "mocks/06-arrays-of-integers.csv",
			expectedJSON: `[{"age":[25,26]}]`,
		},
		{
			name:         "Arrays of objects",
			filename:     "mocks/07-arrays-of-objects.csv",
			expectedJSON: `[{"users":[{"firstname":"John","lastname":"Doe"},{"firstname":"Jane","lastname":"Doe"}]}]`,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// Assert that the csv is valid
			csv, _ := ReadCSV(testcase.filename)
			json, _ := csv.ToJSON()
			assert.Equal(t, testcase.expectedJSON, json)
		})
	}
}
