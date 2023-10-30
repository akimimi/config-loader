package config_loader

import "testing"

func TestFcServiceConfigLoader(t *testing.T) {
	config := FcServiceConfig{}
	yamlstr := `
    endpoint: http://url
    access_key_id: some_key_id
    access_key_secret: some_secret
    `
	config.LoadByBytes([]byte(yamlstr))

	if config.AccessKeyId != "some_key_id" {
		t.Log("Aliyun Function Service config YAML unmarshal failed!")
		t.Fail()
	}
}

func TestFunctionConfigLoader(t *testing.T) {
	config := FunctionConfig{}
	yamlstr := `
    service_name: service
    function_name: function
    `
	config.LoadByBytes([]byte(yamlstr))

	if config.FunctionName != "function" {
		t.Log("Aliyun Function Config YAML unmarshal failed!")
		t.Fail()
	}
}
