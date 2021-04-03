package config_loader

import (
	"gitee.com/akimimi/getuigo"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type GetuiConfigGroup map[string]getuigo.GetuiConfig

func (g *GetuiConfigGroup) LoadByFile(filename string) {
	if file, e := ioutil.ReadFile(filename); e != nil {
		panic(e)
	} else {
		g.LoadByBytes(file)
	}
}

func (g *GetuiConfigGroup) LoadByBytes(content []byte) {
	if e := yaml.Unmarshal(content, g); e != nil {
		panic(e)
	}
}
