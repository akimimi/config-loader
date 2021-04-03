package config_loader

import "testing"

func TestGetuiConfigLoader(t *testing.T) {
	config := GetuiConfigGroup{}
	yamlstr := `
    shopend:
        appid: yNLX9pFEWY9hA
        appsecret: Xf
        appkey: eLP
        mastersecret: OU9YH3omlZ617
    staffend:
        appid: yNLX9pFEWY9hA
        appsecret: Xf
        appkey: eLP
        mastersecret: OU9YH3omlZ617
    `
	config.LoadByBytes([]byte(yamlstr))
	assert1 := (config["shopend"].AppId == "yNLX9pFEWY9hA")
	assert2 := (config["staffend"].MasterSecret == "OU9YH3omlZ617")
	if !assert1 || !assert2 {
		t.Log("Getui config YAML unmarshal failed!")
		t.Fail()
	}
}
