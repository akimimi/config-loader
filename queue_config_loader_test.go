package config_loader

import "testing"

func TestAliqueueConfig(t *testing.T) {
	config := QueueConfig{}
	yamlstr := `
    url: "http://xxxx.mns.cn-beijing.aliyuncs.com/"
    access_key_id: ""
    access_key_secret: ""
    queue_name: "gotest"
    `

	config.LoadByBytes([]byte(yamlstr))
	if config.QueueName != "gotest" || config.Verbose {
		t.Log("Aliqueue config YAML unmarshal failed!")
		t.Fail()
	}
}
