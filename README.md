[![GoDoc](http://godoc.org/github.com/omotto/basicCron?status.png)](http://godoc.org/github.com/omotto/jsonIterator)
[![Build Status](https://travis-ci.com/omotto/basicCron.svg?branch=main)](https://travis-ci.com/omotto/jsonIterator)
[![Coverage Status](https://coveralls.io/repos/github/omotto/basicCron/badge.svg)](https://coveralls.io/github/omotto/jsonIterator)

# jsonIterator

## GoLang JSON Iterator

In order to avoid reading big size files with JSON content format from multiple Goroutines; It's developed a JSON iterator to read directly from file one object element only from JSON array each time. 

Avoiding memory filling, in contrast, each reading takes more time than reading preloaded memory file  content.

```
type MyStruct struct {
    Id			int64 						`json:"id"`
    Entity		string						`json:"entity"`
    Update		[]map[string]interface{}	`json:"update"`
    Operation	string						`json:"operation"`
}

if file, err := os.Open("myjsonfile.txt"); err == nil {
    iter := jsoniterator.NewJSONIterator(file, MyStruct{})
    for value, ok := iter.Next(); ok; value, ok = iter.Next() {
        if err := iter.Error(); err == nil {
            jsonValue := MyStruct{}
            jsonString, _ := json.Marshal(value)
            if err := json.Unmarshal(jsonString, &jsonValue); err == nil {
                // Read data from jsonValue
            } else {
                fmt.Errorf(err)
            }
        } else if err.Error() != "EOF" { // ! End Of File
            fmt.Errorf(err)
        }
    }
    e := file.Close()
    if err == nil {
        err = e
    }
} else {
    fmt.Errorf(err)
}
```
   
