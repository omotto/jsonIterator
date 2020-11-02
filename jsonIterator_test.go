package jsonIterator

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

type MyStruct struct {
	Id			int64 						`json:"id"`
	Entity		string						`json:"entity"`
	Update		[]map[string]interface{}	`json:"update"`
	Operation	string						`json:"operation"`
}

func TestJSONIterator(t *testing.T) {
	numIters := 0
	if file, err := os.Open("test.txt"); err == nil {
		iter := NewJSONIterator(file, MyStruct{})
		for value, ok := iter.Next(); ok; value, ok = iter.Next() {
			if err := iter.Error(); err == nil {
				jsonValue := MyStruct{}
				jsonString, _ := json.Marshal(value)
				if err := json.Unmarshal(jsonString, &jsonValue); err == nil {
					log.Println(jsonValue.Id)
					log.Println(jsonValue.Operation)
					log.Println(jsonValue.Entity)
					log.Println(jsonValue.Update)
					numIters++
				} else {
					t.Error(err)
				}
			} else if err.Error() != "EOF" {
				t.Error(err)
			}
		}
		_ = file.Close()
	} else {
		t.Error(err)
	}
	if numIters != 12 {
		t.Error("missing entries")
	}

}