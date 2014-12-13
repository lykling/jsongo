package jsongo

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

// JSONGo a type for using json
type JSONGo struct {
	raw interface{}
}

// Get get json attribute
func (cj *JSONGo) Get(key string) (ret *JSONGo) {
	switch cj.GetType() {
	case "String":
		ret = &JSONGo{cj.raw}
	case "Array":
		tmp := cj.raw.([]interface{})
		idx, _ := strconv.Atoi(key)
		if idx < len(tmp) && idx >= 0 {
			ret = &JSONGo{tmp[idx]}
		} else {
			ret = nil
		}
	case "Object":
		tmp := cj.raw.(map[string]interface{})[key]
		if tmp != nil {
			ret = &JSONGo{tmp}
		} else {
			ret = nil
		}
	}
	return
}

// GetType get type of JSONGo object
func (cj *JSONGo) GetType() (ret string) {
	switch cj.raw.(type) {
	case string:
		ret = "String"
	case []interface{}:
		ret = "Array"
	case interface{}:
		ret = "Object"
	}
	return
}

// ToString convert JSONGo to string
func (cj *JSONGo) ToString() (ret string) {
	tmp, _ := json.Marshal(cj.raw)
	ret = string(tmp)
	return
}

// Pretty convert JSONGo to pretty string
func (cj *JSONGo) Pretty() (ret string) {
	tmp, _ := json.MarshalIndent(cj.raw, "", "    ")
	ret = string(tmp)
	return
}

// LoadString load json from string
func (cj *JSONGo) LoadString(src string) {
	json.Unmarshal([]byte(src), &cj.raw)
}

// Load load json from Reader
func (cj *JSONGo) Load(r io.Reader) {
	json.NewDecoder(r).Decode(&cj.raw)
}

// LoadFile load json from file
func (cj *JSONGo) LoadFile(name string) {
	file, _ := os.OpenFile(name, os.O_RDONLY, 0644)
	cj.Load(file)
	defer file.Close()
}
