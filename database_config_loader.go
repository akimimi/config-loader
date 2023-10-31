package config_loader

import (
	"fmt"
)

const DataTypeMongo = "mongo"
const DataTypeMysql = "mysql"

type DatasourceConfig struct {
	Server         string `json:"server" yaml:"server"`
	Port           int    `json:"port" yaml:"port"`
	Username       string `json:"username" yaml:"username"`
	Password       string `json:"password" yaml:"password"`
	Query          string `json:"query" yaml:"query"`
	Database       string `json:"database" yaml:"database"`
	Dbtype         string `json:"dbytpe" yaml:"dbtype"`
	ReplicaSet     string `json:"replica_set" yaml:"replica_set,omitempty"`
	ReplicaServers string `json:"replica_servers" yaml:"replica_servers,omitempty"`
	ReadPreference string `json:"replica_read_preference" yaml:"replica_read_preference,omitempty"`
	Prefix         string `json:"prefix" yaml:"prefix"`
	ConfigLoader
}

type DatasourceConfigGroup map[string]DatasourceConfig

func (dsg *DatasourceConfigGroup) LoadByFile(filename string) {
	LoadByFile(filename, dsg)
	dsg.SetDefault()
}

func (dsg *DatasourceConfigGroup) LoadByBytes(content []byte) {
	LoadByBytes(content, dsg)
	dsg.SetDefault()
}

func (dsg *DatasourceConfigGroup) SetDefault() {
	for name, c := range *dsg {
		if c.Port == 0 {
			c.Port = dsg.defaultDatasourcePort(c.Dbtype)
		}
		(*dsg)[name] = c
	}
}

func (dsg *DatasourceConfigGroup) ConnectString(configName string) string {
	if c, ok := (*dsg)[configName]; ok {
		return c.String()
	}
	return ""
}

func (dsg *DatasourceConfigGroup) DatabaseName(configName string) string {
	if c, ok := (*dsg)[configName]; ok {
		return c.Database
	}
	return ""
}

func (dsg *DatasourceConfigGroup) defaultDatasourcePort(dbtype string) int {
	switch dbtype {
	case DataTypeMongo:
		return 27017
	default:
		return 0
	}
}

func (dsc *DatasourceConfig) String() string {
	s := ""
	if dsc.ReplicaSet != "" {
		if dsc.Username != "" && dsc.Password != "" {
			s = fmt.Sprintf("%s:%s@%s", dsc.Username, dsc.Password, dsc.ReplicaServers)
		} else {
			s = fmt.Sprintf("%s", dsc.ReplicaServers)
		}
		if dsc.ReadPreference == "" {
			dsc.ReadPreference = "primary"
		}
		s += fmt.Sprintf("/?replicaSet=%s&readPreference=%s",
			dsc.ReplicaSet, dsc.ReadPreference)
		if dsc.Query != "" {
			s += "&" + dsc.Query
		}
	} else {
		if dsc.Username != "" && dsc.Password != "" {
			s = fmt.Sprintf("%s:%s@%s:%d", dsc.Username, dsc.Password, dsc.Server, dsc.Port)
		} else {
			s = fmt.Sprintf("%s:%d", dsc.Server, dsc.Port)
		}
		s += "/"
		if dsc.Dbtype != DataTypeMongo && dsc.Database != "" {
			s += dsc.Database
		}
		if dsc.Query != "" {
			s += "?" + dsc.Query
		}
	}
	if dsc.Protocol() != "" {
		return dsc.Protocol() + "://" + s
	} else {
		return s
	}
}

func (dsc *DatasourceConfig) Protocol() string {
	switch dsc.Dbtype {
	case DataTypeMongo:
		return "mongodb"
	default:
		return ""
	}
}
