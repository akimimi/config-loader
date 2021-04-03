package config_loader

import (
	"gitee.com/akimimi/getuigo"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type GetuiConfigGroup map[string]getuigo.GetuiConfig

func (this *GetuiConfigGroup) LoadByFile(filename string) {
	if file, e := ioutil.ReadFile(filename); e != nil {
		panic(e)
	} else {
		this.LoadByBytes(file)
	}
}

func (this *GetuiConfigGroup) LoadByBytes(content []byte) {
	if e := yaml.Unmarshal(content, this); e != nil {
		panic(e)
	}
}
