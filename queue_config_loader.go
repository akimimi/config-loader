package config_loader

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type QueueConfig struct {
	Url             string `json:"url" yaml:"url"`
	AccessKeyId     string `json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret"`
	QueueName       string `json:"queue_name" yaml:"queue_name"`
	MaxDequeueCount int    `json:"max_dequeue_count" yaml:"max_dequeue_count"`
}

func (c *QueueConfig) LoadByFile(filename string) {
	if file, e := ioutil.ReadFile(filename); e != nil {
		panic(e)
	} else {
		c.LoadByBytes(file)
	}
}

func (c *QueueConfig) LoadByBytes(content []byte) {
	if e := yaml.Unmarshal(content, c); e != nil {
		panic(e)
	}
}
