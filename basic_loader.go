package config_loader

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

func LoadByFile(filename string, out interface{}) {
	if file, e := ioutil.ReadFile(filename); e != nil {
		panic(e)
	} else {
		LoadByBytes(file, out)
	}
}

func LoadByBytes(content []byte, out interface{}) {
	if e := yaml.Unmarshal(content, out); e != nil {
		panic(e)
	}
}
