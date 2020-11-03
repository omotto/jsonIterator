package jsoniterator

import (
	"encoding/json"
	"os"
)

type iterator struct {
	valueChan 	<-chan interface{}
	okChan 		<-chan bool
	errChan 	<-chan error
	err			error
}

func (i *iterator) Next() (interface{}, bool) {
	var (
		value 	interface{}
		ok 		bool
	)
	value, ok, i.err = <-i.valueChan, <-i.okChan, <-i.errChan
	return value, ok
}

func (i *iterator) Error() error {
	return i.err
}

// NewJSONIterator Generator function that produces data
func NewJSONIterator(f *os.File, jsonLine interface{}) iterator {
	var (
		chain    []byte
		b        []byte
		e		 error
	)
	// --
	b = make([]byte, 1)
	_, e = f.Read(b) // remove '['
	// --
	out := make(chan interface{})
	ok := make(chan bool)
	err := make(chan error)
	// Go Routine
	go func() {
		defer close(out) // closes channel upon fn return
		for e == nil {
			for e == nil {
				_, e = f.Read(b)
				chain = append(chain, b[0])
				if json.Unmarshal(chain, &jsonLine) == nil {
					break
				}
			}
			out <- jsonLine // Send word to channel and waits for its reading
			ok <- true
			err <- e // if there was any error, change its value
			// --
			_, _ = f.Read(b) // remove ','
			chain = nil // clear current bytes
			jsonLine = nil // clear current jsonLine
		}
		out <- nil
		ok <- false
		err <- nil
	}()
	return iterator{ out, ok, err, nil }
}

