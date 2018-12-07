package examples

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var l sync.Mutex

type Colors struct {
	ColorStc []ColorsStruct `json:"colors"`
}

type ColorsStruct struct {
	Color    string     `json:"color"`
	Category string     `json:"category"`
	Type     string     `json:"primary"`
	Code     TypeOfCode `json:"code"`
}

type TypeOfCode struct {
	Rgba []int  `json:"rgba"`
	Hex  string `json:"hex"`
}

var Unmarshal = func(b []byte, c Colors) (Colors, error) {
	e := json.Unmarshal(b, &c)
	if e != nil {
		return Colors{}, e
	}
	return c, nil
}

var result Colors

func GetData(file string) {
	l.Lock()
	defer l.Unlock()
	path := filepath.Join("files/", file)
	b, e := ioutil.ReadFile(path)
	Check("Error trying to read the file", e)
	j, e := Unmarshal(b, result)
	Check("Error trying to unmarshing file", e)

	//combining colors
	WriteFromRead(j, "rtow.txt")
	j.CombiningColors()

}

func (c Colors) CombiningColors() {
	color := c.ColorStc
	black := color[0].Color
	white := color[1].Color
	red := color[2].Color
	blue := color[3].Color
	yellow := color[4].Color
	green := color[5].Color
	fmt.Println("Combining colors, to get:")
	fmt.Printf("Gray: %v and %v\n", white, black)
	fmt.Printf("Pink: %v and %v\n", blue, red)
	fmt.Printf("Orange: %v and %v\n", yellow, red)
	fmt.Printf("Cyan: %v and %v\n", blue, green)
	fmt.Printf("Yellow:  %v and %v\n", green, red)
	//would like to create another file where these combinations can apply considering also their rgba and hx
}
