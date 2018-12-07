package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type JsonStruct struct {
	JStruct map[int]string
}

var (
	lock sync.Mutex
)

func Check(s string, e error) {
	if e != nil {
		log.Fatalf("%v, %v", s, e)
	}
}

var Marshal = func(data JsonStruct) ([]byte, error) {
	m, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (data JsonStruct) Add(key int, value string, file string) {
	lock.Lock()
	defer lock.Unlock()
	path := filepath.Join("files/", file)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	Check("Error trying get the file", err)
	_, ok := data.JStruct[key]
	if !ok {
		data.JStruct[key] = value
		b, e := Marshal(data)
		Check("Error trying to marshal the data", e)
		_, e = f.Write(b)
		Check("Error trying to write in to the file", e)
	} else {
		fmt.Printf("Duplicated key: %v\n\n", key)
	}
	defer f.Close()
}

func WriteFromRead(v interface{}, file string) {
	lock.Lock()
	defer lock.Unlock()
	path := filepath.Join("files/", file)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	Check("Error trying to get the file: ", err)
	b, e := json.Marshal(v)
	Check("Error trying to marsh the data", e)
	_, e = f.Write(b)
	Check("Error trying to write the data", e)
	defer f.Close()
}
