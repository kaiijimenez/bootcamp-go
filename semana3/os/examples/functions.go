package examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

const (
	O_RDWR   int = syscall.O_RDWR  // open the file read-write.
	O_CREATE int = syscall.O_CREAT // create a new file if none exists.
)

type Data struct {
	M map[int]string
}

var Unmarshal = func(b []byte, d Data) (Data, error) {
	e := json.Unmarshal(b, &d)
	if e != nil {
		return Data{}, e
	}
	return d, nil
}

var Marshal = func(d Data) ([]byte, error) {
	m, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d Data) Save(f string) error {
	file, e := os.OpenFile(f, O_RDWR|O_CREATE, 0777)
	defer file.Close()
	if e != nil {
		return e
	}
	b, err := Marshal(d)
	by, e := file.Read(b)
	if by == 0 {
		_, err = file.WriteAt(b, 0)
	} else {
		_, err = file.WriteAt(b, int64(by)+1)
	}
	if err != nil {
		return err
	}
	//_, err = file.WriteAt(b, int64(by))
	return err
}

func (d Data) Add(id int, v string, f string) {
	_, ok := d.M[id]
	if !ok {
		d.M[id] = v
		err := d.Save(f)
		fmt.Println(err)
	}
	//fmt.Println("Duplicate key in the file!")
}

func Load(f string) error {
	var m Data
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	u, err := Unmarshal(data, m)
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
