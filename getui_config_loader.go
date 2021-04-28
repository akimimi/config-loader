package config_loader

import (
	"gitee.com/akimimi/getuigo"
)

type GetuiConfigGroup map[string]getuigo.GetuiConfig

func (g *GetuiConfigGroup) LoadByFile(filename string) {
	LoadByFile(filename, g)
}

func (g *GetuiConfigGroup) LoadByBytes(content []byte) {
	LoadByBytes(content, g)
}
