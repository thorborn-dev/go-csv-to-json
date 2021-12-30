# Go: csv-to-json

Convert CSV to (nested) JSON.

## Table of Contents

- [Usage](#usage)
- [Examples](#examples)
  - [Basic Data Types](#basic-data-types)
  - [Objects](#objects)
  - [Arrays](#arrays)
    - [Arrays of Basic Data Types](#arrays-of-basic-data-types)
    - [Arrays of Objects](#arrays-of-objects)
- [Currenty not supported CSV input](#currently-not-supported-csv-input)
  - [Arrays of Arrays](#arrays-of-arrays)

## Usage

```go
package main

import (
    "fmt"

    "github.com/thorborn-dev/go-csv-to-json"
)

func main() {
    csv, err := csv.ReadCSV("path/to/csv")
    if err != nil {
        panic(err)
    }
    json, err := csv.ToJSON()
    if err != nil {
        panic(err)
    }
    fmt.Println(json)
}
```

## Examples

### Basic Data Types

CSV Input:

```csv
name,age,registered,verified
John,25,TRUE,false
```

JSON Output:

```json
[
  {
    "name": "John",
    "age": 25,
    "registered": true,
    "verified": false
  }
]
```

### Objects

CSV Input:

```csv
user.firstname,user.lastname
John,Doe
```

JSON Output:

```json
[
  {
    "user": {
      "firstname": "John",
      "lastname": "Doe"
    }
  }
]
```

### Arrays

### Arrays of Basic Data Types

CSV Input:

```csv
names[0],names[1]
John,Jane
```

JSON Output:

```json
[
  {
    "names": ["John", "Jane"]
  }
]
```

### Arrays of Objects

CSV Input:

```csv
users[0].firstname,users[0].lastname,users[1].firstname,users[1].lastname
John,Doe,Jane,Doe
```

JSON Output:

```json
[
  {
    "users": [
      {
        "firstname": "John",
        "lastname": "Doe"
      },
      {
        "firstname": "Jane",
        "lastname": "Doe"
      }
    ]
  }
]
```

## Currently not supported CSV input

### Arrays of Arrays

CSV Input:

```csv
users[0][0],users[0][1]
John,Jane
```

Expected JSON Output:

```json
[
  {
    "users": [["John", "Jane"]]
  }
]
```

## License

[MIT](LICENSE)
